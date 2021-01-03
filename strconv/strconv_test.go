package strconv

import (
	"bytes"
	"testing"
)

func Test_B2S(t *testing.T) {
	type args struct {
		b []byte
	}

	tests := struct {
		args args
		want string
	}{
		args: args{
			b: []byte("Test"),
		},
		want: "Test",
	}

	if got := B2S(tests.args.b); got != tests.want {
		t.Errorf("B2S(): '%v', want: '%v'", got, tests.want)
	}
}

func Test_S2B(t *testing.T) {
	type args struct {
		b string
	}

	tests := struct {
		args args
		want []byte
	}{
		args: args{
			b: "Test",
		},
		want: []byte("Test"),
	}

	if got := S2B(tests.args.b); !bytes.Equal(got, tests.want) {
		t.Errorf("S2B(): '%v', want: '%v'", got, tests.want)
	}
}
