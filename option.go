package optres

import (
	"github.com/wlrgo/option"
	"github.com/wlrgo/result"
)

func OkOr[T, E any](o option.Option[T], e E) result.Result[T, E] {
	return OkOrElse(o, func() E { return e })
}

func OkOrElse[T, E any](o option.Option[T], e func() E) result.Result[T, E] {
	return option.MapOrElse(
		o,
		func() result.Result[T, E] { return result.Err[T](e()) },
		func(t T) result.Result[T, E] { return result.Ok[T, E](t) },
	)
}

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
