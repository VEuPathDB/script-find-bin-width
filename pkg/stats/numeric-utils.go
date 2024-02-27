package stats

import (
	"math"

	"find-bin-width/pkg/xmath"
)

func findNumBins[V float64 | int64](values []V) int {
	var numBins float64
	setBins := false

	if len(values) > 200 {
		numBins = xmath.FD(values)
		setBins = true
	}

	skewness, _ := xmath.Moments_Skewness(values)
	absSkew := math.Abs(skewness)

	if absSkew > 0.5 {
		n := float64(len(values))
		se := math.Sqrt(6 * (n - 2) / ((n + 1) * (n + 3)))
		ke := math.Log2(1 + absSkew/se)
		numBins = math.Ceil(xmath.Sturges(values) + ke)
		setBins = true
	}

	if !setBins {
		numBins = xmath.Sturges(values)
	}

	return int(numBins)
}