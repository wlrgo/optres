package optres_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
	"github.com/wlrgo/optres"
	"github.com/wlrgo/result"
)

var errTest = errors.New("test")

func TestOkOr(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want result.Result[int, error]
	}{
		{"none", option.None[int](), result.Err[int](errTest)},
		{"some", option.Some(67), result.Ok[int, error](67)},
		{"some zero", option.Some(0), result.Ok[int, error](0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optres.OkOr(tt.give, errTest)
			assert.True(t, result.Equal(got, tt.want))
		})
	}
}

func TestOkOrElse(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		want      result.Result[int, error]
		wantCalls int
	}{
		{"none", option.None[int](), result.Err[int](errTest), 1},
		{"some", option.Some(67), result.Ok[int, error](67), 0},
		{"some zero", option.Some(0), result.Ok[int, error](0), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := optres.OkOrElse(tt.give, func() error {
				calls++
				return errTest
			})
			assert.Equal(t, tt.wantCalls, calls)
			assert.True(t, result.Equal(got, tt.want))
		})
	}
}

func TestTransposeOption(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[result.Result[int, error]]
		want result.Result[option.Option[int], error]
	}{
		{
			"none",
			option.None[result.Result[int, error]](),
			result.Ok[option.Option[int], error](option.None[int]()),
		},
		{
			"some err",
			option.Some(result.Err[int](errTest)),
			result.Err[option.Option[int]](errTest),
		},
		{
			"some ok",
			option.Some(result.Ok[int, error](67)),
			result.Ok[option.Option[int], error](option.Some(67)),
		},
		{
			"some ok zero",
			option.Some(result.Ok[int, error](0)),
			result.Ok[option.Option[int], error](option.Some(0)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optres.TransposeOption(tt.give)
			assert.True(t, result.Equal(got, tt.want))
		})
	}
}

func TestTransposeOptionRoundtrip(t *testing.T) {
	tests := []option.Option[result.Result[int, error]]{
		option.None[result.Result[int, error]](),
		option.Some(result.Err[int](errTest)),
		option.Some(result.Ok[int, error](67)),
		option.Some(result.Ok[int, error](0)),
	}

	for _, give := range tests {
		got := optres.TransposeResult(optres.TransposeOption(give))
		assert.True(t, option.Equal(got, give))
	}
}
