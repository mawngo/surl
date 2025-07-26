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
Benchmark/sign/query/decimal/no_opts-24                  1739308               680.6 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/no_opts-24                1829563               649.6 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix-24                   1539720               778.9 ns/op           768 B/op         14 allocs/op
Benchmark/verify/query/decimal/prefix-24                 1784552               670.1 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/skip_query-24               1692614               712.5 ns/op           712 B/op         11 allocs/op
Benchmark/verify/query/decimal/skip_query-24             1000000              1004 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/decimal/skip_scheme-24              1749734               679.6 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/skip_scheme-24            1886732               642.8 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query-24            1475628               813.6 ns/op           776 B/op         15 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query-24          1000000              1033 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_scheme-24           1513455               777.8 ns/op           768 B/op         14 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_scheme-24         1798352               677.5 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query_and_skip_scheme-24            1461530               804.0 ns/op           760 B/op         15 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query_and_skip_scheme-24          1000000              1032 ns/op             808 B/op         13 allocs/op
Benchmark/sign/query/base32/no_opts-24                                           1577683               740.3 ns/op           784 B/op         11 allocs/op
Benchmark/verify/query/base32/no_opts-24                                         1579974               743.5 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix-24                                            1413436               846.0 ns/op           864 B/op         15 allocs/op
Benchmark/verify/query/base32/prefix-24                                          1537825               778.0 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/skip_query-24                                        1558791               767.7 ns/op           792 B/op         12 allocs/op
Benchmark/verify/query/base32/skip_query-24                                      1000000              1124 ns/op             856 B/op         14 allocs/op
Benchmark/sign/query/base32/skip_scheme-24                                       1634894               731.6 ns/op           768 B/op         11 allocs/op
Benchmark/verify/query/base32/skip_scheme-24                                     1607076               738.1 ns/op           352 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_query-24                             1340395               883.2 ns/op           872 B/op         16 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_query-24                           1041976              1151 ns/op             856 B/op         14 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_scheme-24                            1432869               858.8 ns/op           848 B/op         15 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_scheme-24                          1532074               770.8 ns/op           352 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_query_and_skip_scheme-24             1362766               883.3 ns/op           856 B/op         16 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_query_and_skip_scheme-24            980760              1103 ns/op             840 B/op         14 allocs/op
Benchmark/sign/compat/decimal/no_opts-24                                         1753341               691.8 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/no_opts-24                                       1863262               639.9 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix-24                                          1551464               785.9 ns/op           768 B/op         14 allocs/op
Benchmark/verify/compat/decimal/prefix-24                                        1779942               681.1 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/skip_query-24                                      1670769               713.3 ns/op           712 B/op         11 allocs/op
Benchmark/verify/compat/decimal/skip_query-24                                    1000000              1003 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/decimal/skip_scheme-24                                     1741468               690.3 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/skip_scheme-24                                   1855576               644.8 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_query-24                           1479768               819.2 ns/op           776 B/op         15 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_query-24                         1000000              1032 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_scheme-24                          1498940               801.6 ns/op           768 B/op         14 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_scheme-24                        1782526               665.6 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_query_and_skip_scheme-24           1481260               813.0 ns/op           760 B/op         15 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_query_and_skip_scheme-24         1000000              1019 ns/op             808 B/op         13 allocs/op
Benchmark/sign/compat/base32/no_opts-24                                          1689433               710.4 ns/op           784 B/op         11 allocs/op
Benchmark/verify/compat/base32/no_opts-24                                        1726156               697.9 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix-24                                           1476652               811.7 ns/op           864 B/op         15 allocs/op
Benchmark/verify/compat/base32/prefix-24                                         1632342               728.2 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/skip_query-24                                       1661058               731.1 ns/op           792 B/op         12 allocs/op
Benchmark/verify/compat/base32/skip_query-24                                     1000000              1055 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/base32/skip_scheme-24                                      1694796               711.7 ns/op           768 B/op         11 allocs/op
Benchmark/verify/compat/base32/skip_scheme-24                                    1742863               683.2 ns/op           352 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_query-24                            1406416               848.2 ns/op           872 B/op         16 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_query-24                          1000000              1085 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_scheme-24                           1459045               819.1 ns/op           848 B/op         15 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_scheme-24                         1681944               722.4 ns/op           352 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_query_and_skip_scheme-24            1422495               842.9 ns/op           856 B/op         16 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_query_and_skip_scheme-24          1089763              1091 ns/op             840 B/op         14 allocs/op
Benchmark/sign/compat++/decimal/no_opts-24                                        704377              1755 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/no_opts-24                                      784983              1551 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix-24                                         591325              1899 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/decimal/prefix-24                                       778603              1554 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/skip_query-24                                     809388              1487 ns/op            1512 B/op         27 allocs/op
Benchmark/verify/compat++/decimal/skip_query-24                                  1000000              1004 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/decimal/skip_scheme-24                                    689298              1807 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/skip_scheme-24                                  821410              1546 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_query-24                          739192              1618 ns/op            1576 B/op         31 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_query-24                       1000000              1039 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_scheme-24                         613204              1890 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_scheme-24                       755762              1547 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_query_and_skip_scheme-24          769288              1601 ns/op            1560 B/op         31 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_query_and_skip_scheme-24       1000000              1016 ns/op             808 B/op         13 allocs/op
Benchmark/sign/compat++/base32/no_opts-24                                         672054              1792 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/base32/no_opts-24                                       778902              1523 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix-24                                          631080              1890 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/base32/prefix-24                                        725058              1554 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/skip_query-24                                      785128              1515 ns/op            1512 B/op         27 allocs/op
Benchmark/verify/compat++/base32/skip_query-24                                    720949              1564 ns/op            1432 B/op         25 allocs/op
Benchmark/sign/compat++/base32/skip_scheme-24                                     661575              1776 ns/op            1760 B/op         35 allocs/op
Benchmark/verify/compat++/base32/skip_scheme-24                                   788114              1531 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_query-24                           731122              1610 ns/op            1576 B/op         31 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_query-24                         747901              1590 ns/op            1432 B/op         25 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_scheme-24                          619080              1888 ns/op            1824 B/op         39 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_scheme-24                        782395              1554 ns/op            1176 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_query_and_skip_scheme-24           714735              1613 ns/op            1560 B/op         31 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_query_and_skip_scheme-24         768324              1576 ns/op            1416 B/op         25 allocs/op
PASS
ok      github.com/mawngo/surl       135.516s
```
