package xmath

import (
	"math"

	"find-bin-width/pkg/xutil"
)

// HardRound rounds the given value to the nearest whole number.
func HardRound(value float64) float64 {
	return float64(HardRoundToInt(value))
}

// HardRoundToInt rounds the given value to the nearest whole number and returns
// the value as an int.
func HardRoundToInt(value float64) int64 {
	if value < 0 {
		return int64(value - 0.5)
	}

	if value > 0 {
		return int64(value + 0.5)
	}

	return 0
}

func Ceil(x float64) float64 {
	i := float64(int(x))
	return xutil.IfElse(i == x, x, i+1)
}

func Floor(x float64) float64 {
	return float64(int64(x))
}

func NonZeroRound(value float64, digits int) float64 {
	if value == 0 {
		return 0
	}

	r := r_Round_r3(value, digits)

	if r == 0 {
		return NonZeroRound(value, digits+1)
	} else {
		return r
	}
}

// direct port of R's round_r3 function.
func r_Round_r3(x float64, digits int) float64 {
	p10 := math.Pow10(digits)

	if math.IsInf(p10, 1) {
		return x
	} else if p10 == 0 {
		return 0
	}

	x10 := p10 * x
	i10 := math.Floor(x10)
	xd := i10 / p10
	xu := math.Ceil(x10) / p10
	D := (xu - x) - (x - xd)
	e := math.Mod(i10, 2)
	r := x

	if D < 0 || (e != 0 && D == 0) {
		r = xu
	} else {
		r = xd
	}

	return r
}
