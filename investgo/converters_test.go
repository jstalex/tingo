package investgo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/jstalex/tingo/proto"
)

func TestFloatToQuotation(t *testing.T) {
	type args struct {
		number float64
		step   *pb.Quotation
	}
	tests := []struct {
		name string
		args args
		want *pb.Quotation
	}{
		{
			name: "ok",
			args: args{
				number: 30.403,
				step: &pb.Quotation{
					Units: 0,
					Nano:  200000000,
				},
			},
			want: &pb.Quotation{
				Units: 30,
				Nano:  400000000,
			},
		},
		{
			name: "zero step",
			args: args{
				number: 20.123,
				step: &pb.Quotation{
					Units: 0,
					Nano:  0,
				},
			},
			want: &pb.Quotation{
				Units: 0,
				Nano:  0,
			},
		},
		{
			name: "negative number",
			args: args{
				number: -15.678,
				step: &pb.Quotation{
					Units: 0,
					Nano:  500000000,
				},
			},
			want: &pb.Quotation{
				Units: -15,
				Nano:  -500000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FloatToQuotation(tt.args.number, tt.args.step), "FloatToQuotation(%v, %v)", tt.args.number, tt.args.step)
		})
	}
}
