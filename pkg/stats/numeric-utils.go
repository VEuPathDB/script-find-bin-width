package stats

import (
	"math"

	"find-bin-width/pkg/xmath"
)

func findNumBins[V float64 | int64](values []V) int {
	var numBins float64
	setBins := false

	// Freedman-Diaconis is our 'default' method of choice
	// it works well for reasonably large, normally distributed data
	// its also less sensitive to outliers than some alternatives
	if len(values) > 200 {
		numBins = xmath.FD(values)
		setBins = true
	}

	// this metric should give us a sense for normalcy
	skewness, _ := xmath.Skewness(values)
	absSkew := math.Abs(skewness)

	// these data cant be normal, and so should use Doane's Formula instead
	if absSkew > 0.5 {
		numBins = xmath.Doane(absSkew, values)
		setBins = true
	}

	// sturges is shown to work well for n < 200 w normal data 
	if !setBins {
		numBins = xmath.Sturges(values)
	}

	return int(numBins)
}
