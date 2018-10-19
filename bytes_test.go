package gotils

import (
	"testing"
)

func TestExtendByteSlice(t *testing.T) {
	type args struct {
		b       []byte
		needLen int
	}
	type want struct {
		sliceLen int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Initial 0",
			args: args{
				b:       make([]byte, 0),
				needLen: 10,
			},
			want: want{
				sliceLen: 10,
			},
		},
		{
			name: "Initial 10",
			args: args{
				b:       make([]byte, 10),
				needLen: 5,
			},
			want: want{
				sliceLen: 5,
			},
		},
		{
			name: "Initial 0",
			args: args{
				b:       make([]byte, 45),
				needLen: 69,
			},
			want: want{
				sliceLen: 69,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtendByteSlice(tt.args.b, tt.args.needLen)

			gotLen := len(got)
			if gotLen != tt.want.sliceLen {
				t.Errorf("ExtendByteSlice() length = %v, want = %v", gotLen, tt.want.sliceLen)
			}
		})
	}
}
