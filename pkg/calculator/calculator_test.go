package calculator

import (
	"errors"
	"testing"

	internalerrors "github.com/ivanov-nikolay/calculator/internal/errors"
)

func TestCalculator(t *testing.T) {

	type args struct {
		expression string
		want       float64
		wantErr    error
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "positive",
			args: args{
				expression: "1 + 2",
				want:       3,
				wantErr:    nil,
			},
		},
		{
			name: "priority",
			args: args{
				expression: "(2+2)*2",
				want:       8,
				wantErr:    nil,
			},
		}, {
			name: "negative",
			args: args{
				expression: "",
				want:       0,
				wantErr:    internalerrors.ErrInvalidExpressionValues,
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculator(tc.args.expression)
			if got != tc.args.want {
				t.Errorf("Calculator() got = %v, want %v", got, tc.args.want)
			}
			if !errors.Is(err, tc.args.wantErr) {
				t.Errorf("Calculator() err = %v, want %v", err, tc.args.wantErr)
			}
		})
	}
}
