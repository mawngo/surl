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
Benchmark/sign/query/decimal/no_opts-24                  1726131               695.9 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/no_opts-24                1861173               650.2 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix-24                   1523070               779.5 ns/op           768 B/op         14 allocs/op
Benchmark/verify/query/decimal/prefix-24                 1862731               645.9 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/skip_query-24               1664428               749.2 ns/op           712 B/op         11 allocs/op
Benchmark/verify/query/decimal/skip_query-24             1000000              1081 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/decimal/skip_scheme-24              1683492               704.0 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/skip_scheme-24            1880794               654.7 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query-24            1433503               806.8 ns/op           776 B/op         15 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query-24          1000000              1042 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/base32/no_opts-24                           1686919               719.1 ns/op           784 B/op         11 allocs/op
Benchmark/verify/query/base32/no_opts-24                         1750507               702.5 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix-24                            1445919               825.3 ns/op           864 B/op         15 allocs/op
Benchmark/verify/query/base32/prefix-24                          1614448               706.8 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/skip_query-24                        1545638               744.4 ns/op           792 B/op         12 allocs/op
Benchmark/verify/query/base32/skip_query-24                      1000000              1056 ns/op             856 B/op         14 allocs/op
Benchmark/sign/query/base32/skip_scheme-24                       1706456               725.9 ns/op           768 B/op         11 allocs/op
Benchmark/verify/query/base32/skip_scheme-24                     1741126               715.2 ns/op           352 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_query-24             1387015               875.9 ns/op           872 B/op         16 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_query-24           1000000              1091 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/decimal/no_opts-24                         1670108               757.3 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/no_opts-24                       1831152               675.9 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix-24                          1431799               854.6 ns/op           768 B/op         14 allocs/op
Benchmark/verify/compat/decimal/prefix-24                        1784296               655.4 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/skip_query-24                      1654407               785.4 ns/op           712 B/op         11 allocs/op
Benchmark/verify/compat/decimal/skip_query-24                    1000000              1023 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/decimal/skip_scheme-24                     1733173               695.3 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/skip_scheme-24                   1900695               636.0 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_query-24           1451875               822.4 ns/op           776 B/op         15 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_query-24         1000000              1041 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/base32/no_opts-24                          1667378               722.6 ns/op           784 B/op         11 allocs/op
Benchmark/verify/compat/base32/no_opts-24                        1742229               685.7 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix-24                           1440866               831.2 ns/op           864 B/op         15 allocs/op
Benchmark/verify/compat/base32/prefix-24                         1710800               703.3 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/skip_query-24                       1606017               740.3 ns/op           792 B/op         12 allocs/op
Benchmark/verify/compat/base32/skip_query-24                     1000000              1065 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/base32/skip_scheme-24                      1674478               718.3 ns/op           768 B/op         11 allocs/op
Benchmark/verify/compat/base32/skip_scheme-24                    1761490               681.8 ns/op           352 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_query-24            1401469               861.4 ns/op           872 B/op         16 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_query-24          1000000              1077 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat++/decimal/no_opts-24                        697929              1732 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/no_opts-24                      759848              1511 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix-24                         651716              1830 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/decimal/prefix-24                       805158              1522 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/skip_query-24                     842181              1462 ns/op            1512 B/op         27 allocs/op
Benchmark/verify/compat++/decimal/skip_query-24                  1202240              1030 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/decimal/skip_scheme-24                    665287              1750 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/skip_scheme-24                  795554              1492 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_query-24          780578              1554 ns/op            1576 B/op         31 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_query-24       1000000              1034 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/base32/no_opts-24                         693008              1720 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/base32/no_opts-24                       762519              1587 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix-24                          640447              2039 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/base32/prefix-24                        828042              1614 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/skip_query-24                      839318              1516 ns/op            1512 B/op         27 allocs/op
Benchmark/verify/compat++/base32/skip_query-24                    823310              1571 ns/op            1432 B/op         25 allocs/op
Benchmark/sign/compat++/base32/skip_scheme-24                     725416              1892 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/base32/skip_scheme-24                   754630              1520 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_query-24           788736              1621 ns/op            1576 B/op         31 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_query-24         666514              1685 ns/op            1432 B/op         25 allocs/op
PASS
ok      github.com/mawngo/surl  97.801s
```
