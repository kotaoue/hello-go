# mapmem
Study the relationship of map with memory.

## Result
```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/kotaoue/mapmem
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
BenchmarkMakeRuneMap-8           4193691               290.7 ns/op          2072 B/op          2 allocs/op
BenchmarkMake2RuneMap-8          3215701               370.0 ns/op          2712 B/op          2 allocs/op
BenchmarkMakeStringMap-8         1607607               741.6 ns/op          4120 B/op          2 allocs/op