# Benchmark

## Run

```shell
go test -bench=. -benchtime=1s -benchmem
```

## Result

```
❯ go test -bench="."  -benchtime=1s -benchmem                
goos: windows
goarch: amd64
pkg: github.com/mawngo/surl
cpu: AMD Ryzen 9 7900 12-Core Processor
Benchmark/sign/query/decimal/no_opts-24                  1748126               687.8 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/no_opts-24                1902770               630.8 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix-24                   1556365               782.5 ns/op           768 B/op         14 allocs/op
Benchmark/verify/query/decimal/prefix-24                 1818980               654.8 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/skip_query-24               1709386               713.6 ns/op           712 B/op         11 allocs/op
Benchmark/verify/query/decimal/skip_query-24             1207921               996.8 ns/op           824 B/op         13 allocs/op
Benchmark/sign/query/decimal/skip_scheme-24              1741792               690.7 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/skip_scheme-24            1930707               624.4 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query-24            1493556               813.0 ns/op           776 B/op         15 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query-24          1000000              1023 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/base32/no_opts-24                           1617576               709.8 ns/op           784 B/op         11 allocs/op
Benchmark/verify/query/base32/no_opts-24                         1777153               684.4 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix-24                            1443153               818.1 ns/op           864 B/op         15 allocs/op
Benchmark/verify/query/base32/prefix-24                          1696420               702.5 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/skip_query-24                        1635277               732.0 ns/op           792 B/op         12 allocs/op
Benchmark/verify/query/base32/skip_query-24                      1000000              1040 ns/op             856 B/op         14 allocs/op
Benchmark/sign/query/base32/skip_scheme-24                       1685454               713.4 ns/op           768 B/op         11 allocs/op
Benchmark/verify/query/base32/skip_scheme-24                     1776368               666.8 ns/op           352 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_query-24             1417084               854.3 ns/op           872 B/op         16 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_query-24           1000000              1079 ns/op             856 B/op         14 allocs/op
Benchmark/sign/unpool/decimal/no_opts-24                         1789202               673.1 ns/op           704 B/op         10 allocs/op
Benchmark/verify/unpool/decimal/no_opts-24                       1922337               613.3 ns/op           336 B/op          5 allocs/op
Benchmark/sign/unpool/decimal/prefix-24                          1566883               764.3 ns/op           768 B/op         14 allocs/op
Benchmark/verify/unpool/decimal/prefix-24                        1901044               633.4 ns/op           336 B/op          5 allocs/op
Benchmark/sign/unpool/decimal/skip_query-24                      1708538               704.7 ns/op           712 B/op         11 allocs/op
Benchmark/verify/unpool/decimal/skip_query-24                    1217576               985.9 ns/op           824 B/op         13 allocs/op
Benchmark/sign/unpool/decimal/skip_scheme-24                     1733740               677.6 ns/op           704 B/op         10 allocs/op
Benchmark/verify/unpool/decimal/skip_scheme-24                   1959720               617.3 ns/op           336 B/op          5 allocs/op
Benchmark/sign/unpool/decimal/prefix_and_skip_query-24           1508331               795.9 ns/op           776 B/op         15 allocs/op
Benchmark/verify/unpool/decimal/prefix_and_skip_query-24         1000000              1002 ns/op             824 B/op         13 allocs/op
Benchmark/sign/unpool/base32/no_opts-24                          1726084               697.8 ns/op           784 B/op         11 allocs/op
Benchmark/verify/unpool/base32/no_opts-24                        1803924               664.2 ns/op           368 B/op          6 allocs/op
Benchmark/sign/unpool/base32/prefix-24                           1497648               801.3 ns/op           864 B/op         15 allocs/op
Benchmark/verify/unpool/base32/prefix-24                         1736946               693.9 ns/op           368 B/op          6 allocs/op
Benchmark/sign/unpool/base32/skip_query-24                       1661658               722.9 ns/op           792 B/op         12 allocs/op
Benchmark/verify/unpool/base32/skip_query-24                     1000000              1023 ns/op             856 B/op         14 allocs/op
Benchmark/sign/unpool/base32/skip_scheme-24                      1702486               700.0 ns/op           768 B/op         11 allocs/op
Benchmark/verify/unpool/base32/skip_scheme-24                    1796462               668.7 ns/op           352 B/op          6 allocs/op
Benchmark/sign/unpool/base32/prefix_and_skip_query-24            1444988               836.4 ns/op           872 B/op         16 allocs/op
Benchmark/verify/unpool/base32/prefix_and_skip_query-24          1000000              1057 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/decimal/no_opts-24                         1772416               689.3 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/no_opts-24                       1905871               632.5 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix-24                          1541660               789.8 ns/op           768 B/op         14 allocs/op
Benchmark/verify/compat/decimal/prefix-24                        1867768               646.9 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/skip_query-24                      1676830               707.1 ns/op           712 B/op         11 allocs/op
Benchmark/verify/compat/decimal/skip_query-24                    1000000              1007 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/decimal/skip_scheme-24                     1731417               694.6 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/skip_scheme-24                   1905055               634.0 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_query-24           1484649               809.5 ns/op           776 B/op         15 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_query-24         1000000              1014 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/base32/no_opts-24                          1673432               713.6 ns/op           784 B/op         11 allocs/op
Benchmark/verify/compat/base32/no_opts-24                        1769037               675.8 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix-24                           1454277               825.6 ns/op           864 B/op         15 allocs/op
Benchmark/verify/compat/base32/prefix-24                         1700491               704.4 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/skip_query-24                       1605958               739.0 ns/op           792 B/op         12 allocs/op
Benchmark/verify/compat/base32/skip_query-24                     1000000              1046 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/base32/skip_scheme-24                      1706346               716.6 ns/op           768 B/op         11 allocs/op
Benchmark/verify/compat/base32/skip_scheme-24                    1785651               672.7 ns/op           352 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_query-24            1424018               843.6 ns/op           872 B/op         16 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_query-24          1000000              1183 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat++/decimal/no_opts-24                        688684              1872 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/no_opts-24                      770474              1604 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix-24                         612778              2018 ns/op            1826 B/op         39 allocs/op
Benchmark/verify/compat++/decimal/prefix-24                       768408              1614 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/skip_query-24                     773604              1576 ns/op            1513 B/op         27 allocs/op
Benchmark/verify/compat++/decimal/skip_query-24                  1000000              1042 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/decimal/skip_scheme-24                    657318              1879 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/skip_scheme-24                  772941              1595 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_query-24          736624              1692 ns/op            1577 B/op         31 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_query-24       1000000              1071 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/base32/no_opts-24                         663279              1887 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/base32/no_opts-24                       769867              1592 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix-24                          548052              2030 ns/op            1826 B/op         39 allocs/op
Benchmark/verify/compat++/base32/prefix-24                        661546              1635 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/skip_query-24                      746680              1579 ns/op            1513 B/op         27 allocs/op
Benchmark/verify/compat++/base32/skip_query-24                    749614              1622 ns/op            1433 B/op         25 allocs/op
Benchmark/sign/compat++/base32/skip_scheme-24                     639490              1899 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/base32/skip_scheme-24                   668433              1601 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_query-24           710336              1702 ns/op            1577 B/op         31 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_query-24         639668              1646 ns/op            1433 B/op         25 allocs/op
BenchmarkConcurrent/sign-24                                         2016            574187 ns/op          731671 B/op      11011 allocs/op
BenchmarkConcurrent/sign_unpool-24                                   596           2055263 ns/op          734285 B/op      11066 allocs/op
PASS
ok      github.com/mawngo/surl  133.904s
```
