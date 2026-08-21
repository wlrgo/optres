package optres

import (
	"github.com/wlrgo/option"
	"github.com/wlrgo/result"
)

// Err converts r to an [option.Option], mapping [result.Err] to [option.Some]
// and [result.Ok] to [option.None]. The success value is discarded.
func Err[T, E any](r result.Result[T, E]) option.Option[E] {
	return result.MapOrElse(
		r,
		func(e E) option.Option[E] { return option.Some(e) },
		func(T) option.Option[E] { return option.None[E]() },
	)
}

// Ok converts r to an [option.Option], mapping [result.Ok] to [option.Some]
// and [result.Err] to [option.None]. The error is discarded.
func Ok[T, E any](r result.Result[T, E]) option.Option[T] {
	return result.MapOrElse(
		r,
		func(E) option.Option[T] { return option.None[T]() },
		func(t T) option.Option[T] { return option.Some(t) },
	)
}

// TransposeResult converts a [result.Result] of an [option.Option] into an
// [option.Option] of a [result.Result].
//
// [result.Ok] of [option.None] maps to [option.None]. [result.Ok] of
// [option.Some] and [result.Err] map to [option.Some] of [result.Ok] and
// [option.Some] of [result.Err].
func TransposeResult[T, E any](
	r result.Result[option.Option[T], E],
) option.Option[result.Result[T, E]] {
	type out = option.Option[result.Result[T, E]]

	return result.MapOrElse(
		r,
		func(e E) out {
			return option.Some(result.Err[T](e))
		},
		func(o option.Option[T]) out {
			return option.Map(o, func(t T) result.Result[T, E] {
				return result.Ok[T, E](t)
			})
		},
	)
}
