# Learn Go With Tests

Reference to the online book: <https://quii.gitbook.io/learn-go-with-tests>

To run tests from any module, navigate inside that module root dir and run `go test ./...`

To run tests with coverage,

```bash
go test -cover

Output:
PASS
coverage: 100.0% of statements
```

To run benchmarks, run

```bash
go test -bench=.

Sample Output:-
goos: darwin
goarch: amd64
pkg: iteration
cpu: VirtualApple @ 2.50GHz
BenchmarkRepeat-8        8506176               139.1 ns/op
PASS
ok      iteration       2.017s
```'

To detect race condition, use `go test -race`.

To do vulnerability check, use this

```bash
go vet

Sample Output (if there is vulnerabilty):
./counter_test.go:14:20: call of assertCounter copies lock value: go-sync.Counter
./counter_test.go:34:20: call of assertCounter copies lock value: go-sync.Counter
./counter_test.go:38:42: assertCounter passes lock by value: go-sync.Counter
```
