package investgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestMessageFromHeader(t *testing.T) {
	type args struct {
		md metadata.MD
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				md: metadata.New(map[string]string{"message": "txt"}),
			},
			want: "txt",
		},
		{
			name: "fail",
			args: args{
				md: metadata.New(map[string]string{"messagge": "txt"}),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MessageFromHeader(tt.args.md), "MessageFromHeader(%v)", tt.args.md)
		})
	}
}

func TestRemainingLimitFromHeader(t *testing.T) {
	type args struct {
		md metadata.MD
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				md: metadata.New(map[string]string{"x-ratelimit-remaining": "100"}),
			},
			want: 100,
		},
		{
			name: "fail",
			args: args{
				md: metadata.New(map[string]string{"x-ratelimit-remaining": "hh"}),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemainingLimitFromHeader(tt.args.md), "RemainingLimitFromHeader(%v)", tt.args.md)
		})
	}
}
