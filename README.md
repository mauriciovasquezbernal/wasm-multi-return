# Wasm multireturn error

This example tries to show how to return multiple values from a function exposed
on the host to a wasm module. However it doesn't work:

```bash
$ go run .
2024/06/13 16:14:22 failed to instantiate module: import func[env.hello]: signature mismatch: i32i64_v != i64_i64i64
panic: failed to instantiate module: import func[env.hello]: signature mismatch: i32i64_v != i64_i64i64

goroutine 1 [running]:
log.Panicf({0x652e5f?, 0x69d8f0?}, {0xc0009bbf10?, 0x803600?, 0xc00008c300?})
	/usr/local/go/src/log/log.go:439 +0x65
main.main()
	/home/mauriciov/kinvolk/ebpf/experiments/wasm-multi-value/main.go:66 +0x4b2
exit status 2
```
