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
Benchmark/sign/query/decimal/no_opts-24                  1739136               701.8 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/no_opts-24                1848099               634.2 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix-24                   1522428               796.6 ns/op           752 B/op         13 allocs/op
Benchmark/verify/query/decimal/prefix-24                 1817942               657.6 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/skip_query-24               1612502               747.6 ns/op           712 B/op         11 allocs/op
Benchmark/verify/query/decimal/skip_query-24             1000000              1011 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/decimal/skip_scheme-24              1687784               706.2 ns/op           704 B/op         10 allocs/op
Benchmark/verify/query/decimal/skip_scheme-24            1883076               638.4 ns/op           336 B/op          5 allocs/op
Benchmark/sign/query/decimal/prefix_and_skip_query-24            1435072               852.0 ns/op           760 B/op         14 allocs/op
Benchmark/verify/query/decimal/prefix_and_skip_query-24          1000000              1045 ns/op             824 B/op         13 allocs/op
Benchmark/sign/query/base32/no_opts-24                           1627647               718.9 ns/op           768 B/op         10 allocs/op
Benchmark/verify/query/base32/no_opts-24                         1689589               698.6 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix-24                            1440754               882.3 ns/op           832 B/op         13 allocs/op
Benchmark/verify/query/base32/prefix-24                          1518026               792.8 ns/op           368 B/op          6 allocs/op
Benchmark/sign/query/base32/skip_query-24                        1516468               789.9 ns/op           776 B/op         11 allocs/op
Benchmark/verify/query/base32/skip_query-24                      1000000              1126 ns/op             856 B/op         14 allocs/op
Benchmark/sign/query/base32/skip_scheme-24                       1599981               747.2 ns/op           752 B/op         10 allocs/op
Benchmark/verify/query/base32/skip_scheme-24                     1609664               761.9 ns/op           352 B/op          6 allocs/op
Benchmark/sign/query/base32/prefix_and_skip_query-24             1330557               898.1 ns/op           840 B/op         14 allocs/op
Benchmark/verify/query/base32/prefix_and_skip_query-24           1000000              1172 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/decimal/no_opts-24                         1633335               728.8 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/no_opts-24                       1717754               693.4 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix-24                          1454764               820.6 ns/op           752 B/op         13 allocs/op
Benchmark/verify/compat/decimal/prefix-24                        1666099               714.7 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/skip_query-24                      1534538               773.2 ns/op           712 B/op         11 allocs/op
Benchmark/verify/compat/decimal/skip_query-24                    1000000              1051 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/decimal/skip_scheme-24                     1661626               727.9 ns/op           704 B/op         10 allocs/op
Benchmark/verify/compat/decimal/skip_scheme-24                   1710469               696.2 ns/op           336 B/op          5 allocs/op
Benchmark/sign/compat/decimal/prefix_and_skip_query-24           1389909               868.1 ns/op           760 B/op         14 allocs/op
Benchmark/verify/compat/decimal/prefix_and_skip_query-24         1000000              1076 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat/base32/no_opts-24                          1576250               761.4 ns/op           768 B/op         10 allocs/op
Benchmark/verify/compat/base32/no_opts-24                        1596122               758.6 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix-24                           1392487               859.1 ns/op           832 B/op         13 allocs/op
Benchmark/verify/compat/base32/prefix-24                         1554723               771.8 ns/op           368 B/op          6 allocs/op
Benchmark/sign/compat/base32/skip_query-24                       1503386               804.0 ns/op           776 B/op         11 allocs/op
Benchmark/verify/compat/base32/skip_query-24                     1000000              1106 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat/base32/skip_scheme-24                      1587984               746.9 ns/op           752 B/op         10 allocs/op
Benchmark/verify/compat/base32/skip_scheme-24                    1598216               750.1 ns/op           352 B/op          6 allocs/op
Benchmark/sign/compat/base32/prefix_and_skip_query-24            1326300               886.8 ns/op           840 B/op         14 allocs/op
Benchmark/verify/compat/base32/prefix_and_skip_query-24          1000000              1118 ns/op             856 B/op         14 allocs/op
Benchmark/sign/compat++/decimal/no_opts-24                        621435              1891 ns/op            1761 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/no_opts-24                      715648              1623 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix-24                         567412              2015 ns/op            1810 B/op         38 allocs/op
Benchmark/verify/compat++/decimal/prefix-24                       774292              1646 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/skip_query-24                     736412              1772 ns/op            1513 B/op         27 allocs/op
Benchmark/verify/compat++/decimal/skip_query-24                  1000000              1053 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/decimal/skip_scheme-24                    639912              1848 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/decimal/skip_scheme-24                  718605              1607 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/decimal/prefix_and_skip_query-24          636094              1699 ns/op            1561 B/op         30 allocs/op
Benchmark/verify/compat++/decimal/prefix_and_skip_query-24        951142              1080 ns/op             824 B/op         13 allocs/op
Benchmark/sign/compat++/base32/no_opts-24                         598425              1876 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/base32/no_opts-24                       743203              1607 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix-24                          576135              1983 ns/op            1810 B/op         38 allocs/op
Benchmark/verify/compat++/base32/prefix-24                        738406              1606 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/skip_query-24                      817572              1594 ns/op            1513 B/op         27 allocs/op
Benchmark/verify/compat++/base32/skip_query-24                    732273              1622 ns/op            1433 B/op         25 allocs/op
Benchmark/sign/compat++/base32/skip_scheme-24                     620462              1858 ns/op            1762 B/op         35 allocs/op
Benchmark/verify/compat++/base32/skip_scheme-24                   753816              1606 ns/op            1177 B/op         24 allocs/op
Benchmark/sign/compat++/base32/prefix_and_skip_query-24           723240              1684 ns/op            1561 B/op         30 allocs/op
Benchmark/verify/compat++/base32/prefix_and_skip_query-24         716280              1661 ns/op            1433 B/op         25 allocs/op
BenchmarkConcurrent/sign-24                                         4197            289339 ns/op          365738 B/op       5504 allocs/op
PASS
ok      github.com/mawngo/surl  98.120s
```
