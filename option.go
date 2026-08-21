package optres

import (
	"github.com/wlrgo/option"
	"github.com/wlrgo/result"
)

// OkOr converts o to a [result.Result], mapping [option.Some] to [result.Ok]
// and [option.None] to [result.Err] of e.
//
// e is evaluated before OkOr is called. Use [OkOrElse] to compute an error
// only when it is needed.
func OkOr[T, E any](o option.Option[T], e E) result.Result[T, E] {
	return OkOrElse(o, func() E { return e })
}

// OkOrElse converts o to a [result.Result], mapping [option.Some] to
// [result.Ok] and [option.None] to [result.Err] of e().
//
// The function e is not called when o contains a value.
func OkOrElse[T, E any](o option.Option[T], e func() E) result.Result[T, E] {
	return option.MapOrElse(
		o,
		func() result.Result[T, E] { return result.Err[T](e()) },
		func(t T) result.Result[T, E] { return result.Ok[T, E](t) },
	)
}

// TransposeOption converts an [option.Option] of a [result.Result] into a
// [result.Result] of an [option.Option].
//
// [option.Some] of [result.Ok] maps to [result.Ok] of [option.Some],
// [option.Some] of [result.Err] maps to [result.Err], and [option.None]
// maps to [result.Ok] of [option.None].
func TransposeOption[T, E any](
	o option.Option[result.Result[T, E]],
) result.Result[option.Option[T], E] {
	type out = result.Result[option.Option[T], E]

	return option.MapOrElse(
		o,
		func() out {
			return result.Ok[option.Option[T], E](option.None[T]())
		},
		func(r result.Result[T, E]) out {
			return result.Map(r, func(t T) option.Option[T] {
				return option.Some(t)
			})
		},
	)
}
