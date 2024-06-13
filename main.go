package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	wapi "github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed testdata/foo.wasm
var fooWasm []byte

func exportFunction(
	env wazero.HostModuleBuilder,
	name string,
	fn func(ctx context.Context, m wapi.Module, stack []uint64),
	params, results []wapi.ValueType,
) {
	env.NewFunctionBuilder().
		WithGoModuleFunction(wapi.GoModuleFunc(fn), params, results).
		Export(name)
}

func hello(ctx context.Context, m wapi.Module, stack []uint64) {
	fmt.Printf("host: Hello: stack[0]: %d\n",
		stack[0])

	stack[0] = 42
	stack[1] = 43
}

func main() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	env := r.NewHostModuleBuilder("env")

	exportFunction(env, "hello", hello,
		[]wapi.ValueType{
			wapi.ValueTypeI64,
		},
		[]wapi.ValueType{
			wapi.ValueTypeI64,
			wapi.ValueTypeI64,
		},
	)

	if _, err := env.Instantiate(ctx); err != nil {
		log.Panicf("instantiating host module: %s", err)
	}

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	modConfig := wazero.NewModuleConfig().
		WithStdout(os.Stdout)

	mod, err := r.InstantiateWithConfig(ctx, fooWasm, modConfig)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	foo := mod.ExportedFunction("foo")
	if _, err := foo.Call(ctx); err != nil {
		log.Panicf("failed to call foo: %v", err)
	}
}
