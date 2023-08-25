package bw

import (
	"sort"
	"strconv"
	"strings"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xutil"
)

// FindFloatBinWidth determines the bin width for a series of float values.
//
// Returns a string representation of the bin width for the given series of
// values.  An NA bin width is represented as an empty string.
//
// @param values Series of values for which the bin width should be calculated.
//
// @param rmNa Whether NA values should be removed from the input series.  If
// this value is false and the input series contains one or more NA values, an
// empty string will be returned.
//
// @returns A string representation of the bin width.  NA return values will be
// represented as an empty string.
func FindFloatBinWidth(values []float64, rmNa bool) string {
	if len(values) == 0 {
		return "0"
	}

	if rmNa {
		values = floatsRemoveNAs(values)
	} else if floatsContainNA(values) {
		return ""
	}

	if xmath.UniqueN(values) == 1 {
		return "0"
	}

	sort.Float64s(values)

	numBins := findNumBins(values)
	if numBins == 0 {
		return "0"
	}

	res := floatNumBinsToBinWidth(values, numBins)

	return strconv.FormatFloat(xmath.NonZeroRound(res.binWidth, res.avgDigits), 'f', -1, 64)
}

type fnb2bwResult struct {
	avgDigits int
	binWidth  float64
}

func floatNumBinsToBinWidth(values []float64, numBins int) fnb2bwResult {
	info := floatInfo(values)

	numDigits := xutil.IfElse(info.avgDigits > 6, 4, info.avgDigits-1)
	binWidth := xmath.NonZeroRound(info.max-info.min, numDigits) / float64(numBins)

	return fnb2bwResult{info.avgDigits, binWidth}
}

type floatInfoResult struct {
	min       float64
	max       float64
	avgDigits int
}

func floatInfo(values []float64) floatInfoResult {
	sum := 0
	i := values[0]
	a := values[0]

	for _, v := range values {
		tmp := strconv.FormatFloat(v, 'f', -1, 64)
		sum += len(tmp)
		if strings.ContainsRune(tmp, '.') {
			sum--
		}

		if v < i {
			i = v
		} else if v > a {
			a = v
		}
	}

	return floatInfoResult{i, a, sum / len(values)}
}
