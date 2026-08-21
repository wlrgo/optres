// Package optres converts between [option.Option] and [result.Result].
//
// It follows the Rust Option and Result conversion API as closely as Go
// allows, rather than reshaping it into an idiomatic Go design. option and
// result stay independent; this package is the glue.
//
// [Ok] and [Err] convert a [result.Result] to an [option.Option]. They are
// not the [result.Ok] and [result.Err] constructors. [option.Option.OkOr]
// is the Go (T, error) helper; [OkOr] is the Rust conversion to
// [result.Result].
//
// # API
//
// Option to Result: [OkOr], [OkOrElse].
//
// Result to Option: [Ok], [Err].
//
// Transpose: [TransposeOption], [TransposeResult]. They are inverses of each
// other.
//
// # Evaluation
//
// [OkOr] evaluates the error before the call. Use [OkOrElse] to compute an
// error only when the option contains no value.
package optres
