# optres

[![Go Reference](https://pkg.go.dev/badge/github.com/wlrgo/optres.svg)](https://pkg.go.dev/github.com/wlrgo/optres)
[![CI](https://github.com/wlrgo/optres/actions/workflows/ci.yml/badge.svg)](https://github.com/wlrgo/optres/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/wlrgo/optres)](https://github.com/wlrgo/optres/blob/main/go.mod)
[![Release](https://img.shields.io/github/v/release/wlrgo/optres)](https://github.com/wlrgo/optres/releases)
[![License](https://img.shields.io/github/license/wlrgo/optres)](LICENSE)

Conversions between `Option[T]` and `Result[T, E]`.

## Rust API in Go

This is a [wlrgo](https://github.com/wlrgo) package. wlrgo ports Rust
standard-library types to Go and **intentionally copies the Rust API** instead
of reshaping it into an idiomatic Go design. Later packages in the org follow
the same rule.

Names, combinators, and eager vs lazy evaluation follow
[`Option::ok_or`](https://doc.rust-lang.org/std/option/enum.Option.html#method.ok_or)
and [`Result::ok`](https://doc.rust-lang.org/std/result/enum.Result.html#method.ok)
as closely as Go generics allow. [option](https://github.com/wlrgo/option) and
[result](https://github.com/wlrgo/result) stay independent; this package is
the glue.

## Install

Requires Go 1.26.5 or later.

```bash
go get github.com/wlrgo/optres
```

## Example

```go
package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/wlrgo/option"
	"github.com/wlrgo/optres"
)

func main() {
	opt := option.FromOK(strconv.Atoi("2"))
	res := optres.OkOr(opt, errors.New("missing"))
	if v, ok := res.Get(); ok {
		fmt.Println(v)
	}

	fmt.Println(optres.Ok(res).UnwrapOr(-1))
}
```

`Ok` and `Err` convert a Result to an Option. They are not the `result.Ok`
and `result.Err` constructors. `OkOr` is the Rust conversion to Result;
`option.Option.OkOr` is the Go `(T, error)` helper.

## API

| Group | Highlights |
| --- | --- |
| Option to Result | `OkOr`, `OkOrElse` |
| Result to Option | `Ok`, `Err` |
| Transpose | `TransposeOption`, `TransposeResult` |

See [pkg.go.dev/github.com/wlrgo/optres](https://pkg.go.dev/github.com/wlrgo/optres)
for the full API and package contract.

## License

[MIT](LICENSE)
