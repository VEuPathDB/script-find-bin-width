package xmath

import (
	"math"

	"find-bin-width/pkg/xutil"
)

// Round rounds the given value to the given number of decimal places.
func Round(value float64, digits int) float64 {
	if digits < 1 {
		return HardRound(value)
	}

	m := math.Pow(10, float64(digits))
	y := value * m

	return HardRound(y) / m
}

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

	r := Round(value, digits)

	if r == 0 {
		return NonZeroRound(value, digits+1)
	} else {
		return r
	}
}
