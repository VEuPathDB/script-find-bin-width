package xmath

import (
	"math"

	"find-bin-width/pkg/xutil"
)

func Round(value float64, digits int) float64 {
	if digits < 1 {
		return HardRound(value)
	}

	m := math.Pow(10, float64(digits))
	y := value * m

	return HardRound(y) / m
}

func HardRound(value float64) float64 {
	return float64(HardRoundToInt(value))
}

func HardRoundToInt(value float64) int64 {
	if value < 0 {
		return int64(value - 0.5)
	}

	if value > 0 {
		return int64(value + 0.5)
	}

	return 0
}

func CeilToInt(x float64) int64 {
	return int64(Ceil(x))
}

func Ceil(x float64) float64 {
	i := float64(int(x))
	return xutil.IfElse(i == x, x, x+1)
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
