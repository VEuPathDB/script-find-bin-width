package stats

import (
	"sort"
	"strconv"
	"strings"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

// findFloatBinWidth determines the bin width for a series of float values.
//
// Returns a string representation of the bin width for the given series of
// values.  An NA bin width is represented as an empty string.
//
// @param values Series of values for which the bin width should be calculated.
//
// @returns A string representation of the bin width.  NA return values will be
// represented as an empty string.
func findFloatBinWidth(values []float64) Stats {
	if len(values) == 0 {
		return stats[float64]{stringifier: floatStringifier}
	}

	if xmath.UniqueN(values) == 1 {
		mnx := values[0]
		return stats[float64]{
			min:           mnx,
			max:           mnx,
			binWidth:      0,
			mean:          mnx,
			median:        mnx,
			lowerQuartile: mnx,
			upperQuartile: mnx,
			stringifier:   floatStringifier,
			dataType:      xtype.DataTypeFloat,
		}
	}

	sort.Float64s(values)

	numBins := findNumBins(values)
	if numBins == 0 {
		return stats[float64]{stringifier: floatStringifier}
	}

	res := floatNumBinsToBinWidth(values, numBins)

	return stats[float64]{
		min:           res.min,
		max:           res.max,
		binWidth:      xmath.NonZeroRound(res.binWidth, res.avgDigits),
		mean:          xmath.Mean(values),
		median:        xmath.Median(values),
		lowerQuartile: xmath.LowerQuartile(values),
		upperQuartile: xmath.UpperQuartile(values),
		stringifier:   floatStringifier,
		dataType:      xtype.DataTypeFloat,
	}
}

type fnb2bwResult struct {
	min       float64
	max       float64
	avgDigits int
	binWidth  float64
}

func floatNumBinsToBinWidth(values []float64, numBins int) fnb2bwResult {
	info := floatInfo(values)

	numDigits := xutil.IfElse(info.avgDigits > 6, 4, info.avgDigits-1)
	binWidth := xmath.NonZeroRound(info.max-info.min, numDigits) / float64(numBins)

	return fnb2bwResult{min: info.min, max: info.max, avgDigits: info.avgDigits, binWidth: binWidth}
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
