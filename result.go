package optres

import (
	"github.com/wlrgo/option"
	"github.com/wlrgo/result"
)

func Err[T, E any](r result.Result[T, E]) option.Option[E] {
	return result.MapOrElse(
		r,
		func(e E) option.Option[E] { return option.Some(e) },
		func(T) option.Option[E] { return option.None[E]() },
	)
}

func Ok[T, E any](r result.Result[T, E]) option.Option[T] {
	return result.MapOrElse(
		r,
		func(E) option.Option[T] { return option.None[T]() },
		func(t T) option.Option[T] { return option.Some(t) },
	)
}

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
