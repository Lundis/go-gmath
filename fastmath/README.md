# Benchmarking howto

To benchmark WASM, install wasmtime, then run

    cmd /c "set GOOS=wasip1&& set GOARCH=wasm&& go test -bench=. -exec wasmtime"

or

    GOOS=wasip1 GOARCH=wasm go test -bench=. -exec wasmtime

# Benchmark results

    function       AMD64             WASM

    math.Abs       2.039 ns/op       1.000 ns/op

    Atan2          2.242 ns/op       4.131 ns/op
    math.Atan2     5.242 ns/op      18.270 ns/op

    CopySign       2.135 ns/op       3.265 ns/op
    math.CopySign  2.151 ns/op       1.000 ns/op

    math.Floor     1.000 ns/op       1.000 ns/op

    Log2           2.317 ns/op       3.314 ns/op
    math.Log2      7.263 ns/op      28.370 ns/op

    Mod            2.212 ns/op       4.563 ns/op
    ModAbs         2.194 ns/op       4.468 ns/op
    math.Mod      70.290 ns/op     170.700 ns/op

    Round          2.232 ns/op       3.659 ns/op
    RoundPos       2.195 ns/op       2.318 ns/op
    math.Round     2.179 ns/op       3.504 ns/op

    Modf           2.178 ns/op       3.633 ns/op
    Math.Modf      2.227 ns/op       5.860 ns/op

    Cos            2.565 ns/op       9.908 ns/op
    math.Cos       4.873 ns/op      13.380 ns/op

    Sin            2.753 ns/op       9.340 ns/op
    math.Sin       4.260 ns/op      14.100 ns/op

    CosSinFast     2.343 ns/op       4.821 ns/op
    CosSin         4.017 ns/op      11.590 ns/op
    math.SinCos    4.986 ns/op      14.430 ns/op

    math.Sqrt      2.203 ns/op       2.231 ns/op
