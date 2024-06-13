package main

import "fmt"

//export hello
func hello(a uint64) (uint64, uint64)

//export foo
func foo() {
	fmt.Printf("wasm: foo\n")

	ret0, ret1 := hello(55)

	fmt.Printf("was: ret0: %d, ret1: %d\n", ret0, ret1)
}

// main is required for the `wasi` target, even if it isn't used.
// See https://wazero.io/languages/tinygo/#why-do-i-have-to-define-main
func main() {}
