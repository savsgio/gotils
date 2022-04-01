package math

import (
	"math"
	"testing"
)

func TestFloat64Calculator_Add(t *testing.T) {
	fc := NewFloat64Calculator()

	x := 1.0
	y := 0.1

	for i := 0; i < 10; i++ {
		x = fc.Add(x, y)
	}

	want := 2.0

	if want != x {
		t.Errorf("Add() == %v, want %v", want, x)
	}
}

func TestFloat64Calculator_Sub(t *testing.T) {
	fc := NewFloat64Calculator()

	x := 1.0
	y := 0.1

	for i := 0; i < 10; i++ {
		x = fc.Sub(x, y)
	}

	want := 0.0

	if want != x {
		t.Errorf("Sub() == %v, want %v", want, x)
	}
}

func TestFloat64Calculator_Mul(t *testing.T) {
	fc := NewFloat64Calculator()

	x := 1.0
	y := 0.1

	for i := 0; i < 10; i++ {
		x = fc.Mul(x, y)
	}

	want := math.Pow10(-10)

	if want != x {
		t.Errorf("Mul() == %v, want %v", x, want)
	}
}

func TestFloat64Calculator_Div(t *testing.T) {
	fc := NewFloat64Calculator()

	x := 1.0
	y := 0.1

	for i := 0; i < 10; i++ {
		x = fc.Div(x, y)
	}

	want := math.Pow10(10)

	if want != x {
		t.Errorf("Div() == %v, want %v", want, x)
	}
}

func BenchmarkFloat64Calculator_Add(b *testing.B) {
	fc := NewFloat64Calculator()

	x := 1.2345
	y := 6.789123

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fc.Add(x, y)
	}
}

func BenchmarkFloat64Calculator_Sub(b *testing.B) {
	fc := NewFloat64Calculator()

	x := 1.2345
	y := 6.789123

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fc.Sub(x, y)
	}
}

func BenchmarkFloat64Calculator_Mul(b *testing.B) {
	fc := NewFloat64Calculator()

	x := 1.2345
	y := 6.789123

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fc.Mul(x, y)
	}
}

func BenchmarkFloat64Calculator_Div(b *testing.B) {
	fc := NewFloat64Calculator()

	x := 1.2345
	y := 6.789123

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fc.Div(x, y)
	}
}

func BenchmarkFloat64Calculator_Mod(b *testing.B) {
	fc := NewFloat64Calculator()

	x := 1.2345
	y := 6.789123

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fc.Mod(x, y)
	}
}
