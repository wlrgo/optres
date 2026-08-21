package optres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
	"github.com/wlrgo/optres"
	"github.com/wlrgo/result"
)

func TestErr(t *testing.T) {
	tests := []struct {
		name string
		give result.Result[int, error]
		want option.Option[error]
	}{
		{"err", result.Err[int](errTest), option.Some(errTest)},
		{"ok", result.Ok[int, error](67), option.None[error]()},
		{"ok zero", result.Ok[int, error](0), option.None[error]()},
		{"zero value", result.Result[int, error]{}, option.Some(error(nil))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optres.Err(tt.give)
			assert.True(t, option.Equal(got, tt.want))
		})
	}
}

func TestOk(t *testing.T) {
	tests := []struct {
		name string
		give result.Result[int, error]
		want option.Option[int]
	}{
		{"err", result.Err[int](errTest), option.None[int]()},
		{"ok", result.Ok[int, error](67), option.Some(67)},
		{"ok zero", result.Ok[int, error](0), option.Some(0)},
		{"zero value", result.Result[int, error]{}, option.None[int]()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optres.Ok(tt.give)
			assert.True(t, option.Equal(got, tt.want))
		})
	}
}

func TestTransposeResult(t *testing.T) {
	tests := []struct {
		name string
		give result.Result[option.Option[int], error]
		want option.Option[result.Result[int, error]]
	}{
		{
			"err",
			result.Err[option.Option[int]](errTest),
			option.Some(result.Err[int](errTest)),
		},
		{
			"ok none",
			result.Ok[option.Option[int], error](option.None[int]()),
			option.None[result.Result[int, error]](),
		},
		{
			"ok some",
			result.Ok[option.Option[int], error](option.Some(67)),
			option.Some(result.Ok[int, error](67)),
		},
		{
			"ok some zero",
			result.Ok[option.Option[int], error](option.Some(0)),
			option.Some(result.Ok[int, error](0)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optres.TransposeResult(tt.give)
			assert.True(t, option.Equal(got, tt.want))
		})
	}
}

func TestTransposeResultRoundtrip(t *testing.T) {
	tests := []result.Result[option.Option[int], error]{
		result.Err[option.Option[int]](errTest),
		result.Ok[option.Option[int], error](option.None[int]()),
		result.Ok[option.Option[int], error](option.Some(67)),
		result.Ok[option.Option[int], error](option.Some(0)),
	}

	for _, give := range tests {
		got := optres.TransposeOption(optres.TransposeResult(give))
		assert.True(t, result.Equal(got, give))
	}
}
