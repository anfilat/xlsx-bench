# Xlsx readers benchmark in Go

## Read 10K rows sheet from Excel file

| Benchmark                                       | version |     ns/op |  worse |      B/op | worse |  allocs/op | worse |
|-------------------------------------------------|--------:|----------:|-------:|----------:|------:|-----------:|------:|
| [xlsx-sax](https://github.com/anfilat/xlsx-sax) |         |  69830163 |        |    841921 |       |      30903 |       |
| [Tealeg xlsx](https://github.com/tealeg/xlsx)   | v3.3.13 | 366185497 |  x 5.2 | 133952376 | x 159 |    2586854 |  x 83 |
| [Excelize](https://github.com/qax-os/excelize)  |  v2.9.0 | 939339388 | x 13.5 | 170672588 | x 203 |    4556125 | x 147 |
