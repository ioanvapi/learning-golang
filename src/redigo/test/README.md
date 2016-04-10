## 运行时间

```
❯ go test -bench=BenchmarkWithPool
testing: warning: no tests to run
PASS
BenchmarkWithPool-8        20000             79033 ns/op
ok      _/Users/akagi201/Documents/learning-golang/src/redigo/test      2.383s

❯ go test -bench=BenchmarkNoPool
testing: warning: no tests to run
PASS
BenchmarkNoPool-8          10000            259747 ns/op
ok      _/Users/akagi201/Documents/learning-golang/src/redigo/test      2.629s
```

## 内存消耗

```
❯ go test -bench=BenchmarkNoPool -benchmem
testing: warning: no tests to run
PASS
BenchmarkNoPool-8           5000            232842 ns/op            9257 B/op         26 allocs/op
ok      _/Users/akagi201/Documents/learning-golang/src/redigo/test      1.199s

❯ go test -bench=BenchmarkWithPool -benchmem
testing: warning: no tests to run
PASS
BenchmarkWithPool-8        20000             78085 ns/op             277 B/op         11 allocs/op
ok      _/Users/akagi201/Documents/learning-golang/src/redigo/test      2.355s
```
