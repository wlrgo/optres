package optres

import (
	"github.com/wlrgo/option/v2"
	"github.com/wlrgo/result/v2"
)

// Err converts r to an [option.Option], mapping [result.Err] to [option.Some]
// and [result.Ok] to [option.None]. The success value is discarded.
func Err[T, E any](r result.Result[T, E]) option.Option[E] {
	return r.MapOrElse(
		func(e E) option.Option[E] { return option.Some(e) },
		func(T) option.Option[E] { return option.None[E]() },
	)
}

// Ok converts r to an [option.Option], mapping [result.Ok] to [option.Some]
// and [result.Err] to [option.None]. The error is discarded.
func Ok[T, E any](r result.Result[T, E]) option.Option[T] {
	return r.MapOrElse(
		func(E) option.Option[T] { return option.None[T]() },
		func(t T) option.Option[T] { return option.Some(t) },
	)
}

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
	return o.MapOrElse(
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

	return o.MapOrElse(
		func() out {
			return result.Ok[option.Option[T], E](option.None[T]())
		},
		func(r result.Result[T, E]) out {
			return r.Map(func(t T) option.Option[T] { return option.Some(t) })
		},
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

	return r.MapOrElse(
		func(e E) out {
			return option.Some(result.Err[T](e))
		},
		func(o option.Option[T]) out {
			return o.Map(func(t T) result.Result[T, E] { return result.Ok[T, E](t) })
		},
	)
}
