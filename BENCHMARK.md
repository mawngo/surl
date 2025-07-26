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
pkg: github.com/leg100/surl/v2
cpu: AMD Ryzen 9 7900 12-Core Processor             
Benchmark/sign/path/decimal/no_opts-24           1452532               774.2 ns/op           592 B/op         10 allocs/op
Benchmark/verify/path/decimal/no_opts-24         1558614               768.3 ns/op           328 B/op          6 allocs/op
Benchmark/sign/path/decimal/prefix-24            1224864               983.8 ns/op           912 B/op         14 allocs/op
Benchmark/verify/path/decimal/prefix-24          1536451               787.2 ns/op           328 B/op          6 allocs/op
Benchmark/sign/path/decimal/skip_query-24        1545110               781.5 ns/op           576 B/op         10 allocs/op
Benchmark/verify/path/decimal/skip_query-24      1587921               777.7 ns/op           312 B/op          6 allocs/op
Benchmark/sign/path/decimal/skip_scheme-24       1492773               793.7 ns/op           592 B/op         10 allocs/op
Benchmark/verify/path/decimal/skip_scheme-24     1562737               757.0 ns/op           328 B/op          6 allocs/op
Benchmark/sign/path/decimal/prefix_and_skip_query-24             1252663               965.6 ns/op           896 B/op         14 allocs/op
Benchmark/verify/path/decimal/prefix_and_skip_query-24           1523506               792.0 ns/op           312 B/op          6 allocs/op
Benchmark/sign/path/decimal/prefix_and_skip_scheme-24            1204146               970.7 ns/op           912 B/op         14 allocs/op
Benchmark/verify/path/decimal/prefix_and_skip_scheme-24          1544934               779.6 ns/op           328 B/op          6 allocs/op
Benchmark/sign/path/decimal/prefix_and_skip_query_and_skip_scheme-24             1234324               971.8 ns/op           880 B/op         14 allocs/op
Benchmark/verify/path/decimal/prefix_and_skip_query_and_skip_scheme-24           1545310               765.1 ns/op           296 B/op          6 allocs/op
Benchmark/sign/path/base58/no_opts-24                                            1588351               750.6 ns/op           568 B/op         10 allocs/op
Benchmark/verify/path/base58/no_opts-24                                          1621383               730.9 ns/op           336 B/op          7 allocs/op
Benchmark/sign/path/base58/prefix-24                                             1256238               939.1 ns/op           888 B/op         14 allocs/op
Benchmark/verify/path/base58/prefix-24                                           1572859               753.4 ns/op           336 B/op          7 allocs/op
Benchmark/sign/path/base58/skip_query-24                                         1588005               755.0 ns/op           536 B/op         10 allocs/op
Benchmark/verify/path/base58/skip_query-24                                       1645081               716.2 ns/op           304 B/op          7 allocs/op
Benchmark/sign/path/base58/skip_scheme-24                                        1592076               754.0 ns/op           552 B/op         10 allocs/op
Benchmark/verify/path/base58/skip_scheme-24                                      1659190               727.4 ns/op           320 B/op          7 allocs/op
Benchmark/sign/path/base58/prefix_and_skip_query-24                              1287991               939.4 ns/op           856 B/op         14 allocs/op
Benchmark/verify/path/base58/prefix_and_skip_query-24                            1637271               743.5 ns/op           304 B/op          7 allocs/op
Benchmark/sign/path/base58/prefix_and_skip_scheme-24                             1279142               940.4 ns/op           872 B/op         14 allocs/op
Benchmark/verify/path/base58/prefix_and_skip_scheme-24                           1610176               748.4 ns/op           320 B/op          7 allocs/op
Benchmark/sign/path/base58/prefix_and_skip_query_and_skip_scheme-24              1283263               937.8 ns/op           856 B/op         14 allocs/op
Benchmark/verify/path/base58/prefix_and_skip_query_and_skip_scheme-24            1625966               738.5 ns/op           304 B/op          7 allocs/op
Benchmark/sign/query/decimal/no_opts-24                                           596506              2047 ns/op            2224 B/op         41 allocs/op
Benchmark/verify/query/decimal/no_opts-24                                         649807              1878 ns/op            1736 B/op         35 allocs/op
Benchmark/sign/query/decimal/prefix-24                                            561810              2172 ns/op            2288 B/op         45 allocs/op
Benchmark/verify/query/decimal/prefix-24                                          628659              1909 ns/op            1736 B/op         35 allocs/op
Benchmark/sign/query/decimal/skip_query-24                                        478165              2522 ns/op            2800 B/op         54 allocs/op
Benchmark/verify/query/decimal/skip_query-24                                      526518              2343 ns/op            2312 B/op         48 allocs/op
Benchmark/sign/query/decimal/skip_scheme-24                                       585028              2081 ns/op            2224 B/op         41 allocs/op
Benchmark/verify/query/decimal/skip_scheme-24                                     655225              1863 ns/op            1736 B/op         35 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query-24                             444846              2676 ns/op            2864 B/op         58 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query-24                           533792              2381 ns/op            2312 B/op         48 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_scheme-24                            541482              2158 ns/op            2288 B/op         45 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_scheme-24                          631735              1897 ns/op            1736 B/op         35 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query_and_skip_scheme-24             440062              3179 ns/op            2848 B/op         58 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query_and_skip_scheme-24           493749              2452 ns/op            2296 B/op         48 allocs/op
Benchmark/sign/query/base58/no_opts-24                                            594399              2265 ns/op            2152 B/op         41 allocs/op
Benchmark/verify/query/base58/no_opts-24                                          648003              1931 ns/op            1720 B/op         36 allocs/op
Benchmark/sign/query/base58/prefix-24                                             572259              2134 ns/op            2216 B/op         45 allocs/op
Benchmark/verify/query/base58/prefix-24                                           624164              1879 ns/op            1720 B/op         36 allocs/op
Benchmark/sign/query/base58/skip_query-24                                         481243              2447 ns/op            2704 B/op         54 allocs/op
Benchmark/verify/query/base58/skip_query-24                                       507008              2309 ns/op            2272 B/op         49 allocs/op
Benchmark/sign/query/base58/skip_scheme-24                                        528138              1991 ns/op            2136 B/op         41 allocs/op
Benchmark/verify/query/base58/skip_scheme-24                                      648511              1874 ns/op            1704 B/op         36 allocs/op
Benchmark/sign/query/base58/prefix_and_skip_query-24                              408465              2647 ns/op            2768 B/op         58 allocs/op
Benchmark/verify/query/base58/prefix_and_skip_query-24                            429054              2408 ns/op            2272 B/op         49 allocs/op
Benchmark/sign/query/base58/prefix_and_skip_scheme-24                             572409              2156 ns/op            2200 B/op         45 allocs/op
Benchmark/verify/query/base58/prefix_and_skip_scheme-24                           624967              1924 ns/op            1704 B/op         36 allocs/op
Benchmark/sign/query/base58/prefix_and_skip_query_and_skip_scheme-24              460671              2602 ns/op            2768 B/op         58 allocs/op
Benchmark/verify/query/base58/prefix_and_skip_query_and_skip_scheme-24            510258              2356 ns/op            2272 B/op         49 allocs/op
PASS
ok      github.com/leg100/surl/v2       92.300s
```
