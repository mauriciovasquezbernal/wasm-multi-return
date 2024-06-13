all: foo.wasm

foo.wasm: foo.go
	tinygo build -o foo.wasm -target=wasi foo.go
