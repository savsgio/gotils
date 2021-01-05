package bytes

import (
	"testing"
)

func Test_Extend(t *testing.T) {
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
			name: "Initial 69",
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
		b := tt.args.b
		needLen := tt.args.needLen
		sliceLen := tt.want.sliceLen

		t.Run(tt.name, func(t *testing.T) {
			got := Extend(b, needLen)

			gotLen := len(got)
			if gotLen != sliceLen {
				t.Errorf("ExtendByteSlice() length = %v, want = %v", gotLen, sliceLen)
			}
		})
	}
}

func Test_Rand(t *testing.T) {
	n := 32
	dst := make([]byte, n)

	Rand(dst)

	for i := range dst {
		if string(rune(i)) == "" {
			t.Fatalf("RandBytes() invalid char '%v'", dst[i])
		}
	}

	if len(dst) != n {
		t.Fatalf("RandBytes() length '%d', want '%d'", len(dst), n)
	}
}

func Benchmark_Rand(b *testing.B) {
	n := 32
	dst := make([]byte, n)

	for i := 0; i < b.N; i++ {
		Rand(dst)
	}
}
