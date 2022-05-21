# flagx

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Go flag utils.

## Features

* Simple API.
* Dependency-free.
* Clean and tested code.
* Fully compatible with `flag` package.

See [GUIDE.md](https://github.com/cristalhq/flagx/blob/main/GUIDE.md) for more details.

## Install

Go version 1.17+

```
go get github.com/cristalhq/flagx
```

## Example

```go
args := []string{"-t", "20s"} // or os.Args[1:]

var d time.Duration
fset := flagx.NewFlagSet("testing", os.Stderr)
fset.Duration(&d, "timeout", "t", 10*time.Second, "just a timeout")

err := fset.Parse(args)
if err != nil {
	panic(err)
}

fmt.Println(d)

// Output: 20s
```

Also see examples: [examples_test.go](https://github.com/cristalhq/flagx/blob/main/example_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/flagx/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/flagx/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/flagx
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/flagx
[reportcard-img]: https://goreportcard.com/badge/cristalhq/flagx
[reportcard-url]: https://goreportcard.com/report/cristalhq/flagx
[coverage-img]: https://codecov.io/gh/cristalhq/flagx/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/flagx
[version-img]: https://img.shields.io/github/v/release/cristalhq/flagx
[version-url]: https://github.com/cristalhq/flagx/releases
