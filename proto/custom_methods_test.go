package investapi

import (
	"testing"

	"google.golang.org/protobuf/runtime/protoimpl"
)

func TestQuotation_ToFloat(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Units         int64
		Nano          int32
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "",
			fields: fields{
				Units: 1,
				Nano:  346700000,
			},
			want: 1.3467,
		},
		{
			name: "",
			fields: fields{
				Units: 1345,
				Nano:  900800000,
			},
			want: 1345.9008,
		},
		{
			name: "",
			fields: fields{
				Units: 0,
				Nano:  6700000,
			},
			want: 0.0067,
		},
		{
			name: "",
			fields: fields{
				Units: 0,
				Nano:  99,
			},
			want: 0.000000099,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quotation{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Units:         tt.fields.Units,
				Nano:          tt.fields.Nano,
			}
			if got := q.ToFloat(); got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
