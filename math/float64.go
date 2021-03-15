package math

import (
	"bytes"
	"math"
	"strconv"

	"github.com/valyala/bytebufferpool"
)

// NewFloat64Calculator returns a float64 calculator
// for basic operations without losing precision in decimals.
func NewFloat64Calculator() *Float64Calculator {
	return &Float64Calculator{}
}

// places returns the number of decimal.
func (fc *Float64Calculator) places(x float64) int {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	buf.B = strconv.AppendFloat(buf.B, x, 'f', -1, 64)
	i := bytes.IndexByte(buf.B, '.')

	if i > -1 {
		return len(buf.B) - i - 1
	}

	return 0
}

func (fc *Float64Calculator) max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// Add returns a + b.
func (fc *Float64Calculator) Add(a, b float64) float64 {
	exp := math.Pow10(fc.max(fc.places(a), fc.places(b)))

	intA := math.Trunc(a * exp)
	intB := math.Trunc(b * exp)

	return (intA + intB) / exp
}

// Sub returns a - b.
func (fc *Float64Calculator) Sub(a, b float64) float64 {
	exp := math.Pow10(fc.max(fc.places(a), fc.places(b)))

	intA := math.Trunc(a * exp)
	intB := math.Trunc(b * exp)

	return (intA - intB) / exp
}

// Mul returns a * b.
func (fc *Float64Calculator) Mul(a, b float64) float64 {
	placesA := fc.places(a)
	placesB := fc.places(b)

	expA := math.Pow10(placesA)
	expB := math.Pow10(placesB)

	intA := math.Trunc(a * expA)
	intB := math.Trunc(b * expB)

	exp := math.Pow10(placesA + placesB)

	return (intA * intB) / exp
}

// Div returns a / b.
func (fc *Float64Calculator) Div(a, b float64) float64 {
	placesA := fc.places(a)
	placesB := fc.places(b)

	expA := math.Pow10(placesA)
	expB := math.Pow10(placesB)

	intA := math.Trunc(a * expA)
	intB := math.Trunc(b * expB)

	exp := math.Pow10(placesA - placesB)

	return (intA / intB) / exp
}

// Mod returns a % b.
func (fc *Float64Calculator) Mod(a, b float64) float64 {
	quo := math.Trunc(fc.Div(a, b))

	return fc.Sub(a, fc.Mul(b, quo))
}
