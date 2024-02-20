package stats

import (
	"math"
	"sort"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

// findIntegerBinWidth determines the bin width for a series of integral values.
//
// Returns a string representation of the bin width for the given series of
// values.  An NA bin width is represented as an empty string.
//
// @param values Series of values for which the bin width should be calculated.
//
// @returns A string representation of the bin width.  NA return values will be
// represented as an empty string.
func findIntegerBinWidth(values []int64) Stats {
	out := intFindBinWidth(values)

	return stats[float64]{
		min:           float64(out.min),
		max:           float64(out.max),
		binWidth:      float64(out.binWidth),
		mean:          xmath.Mean(values),
		median:        xmath.Median(values),
		lowerQuartile: xmath.LowerQuartile(values),
		upperQuartile: xmath.UpperQuartile(values),
		stringifier:   intStringifier,
		dataType:      xtype.DataTypeInteger,
	}
}

func intFindBinWidth(values []int64) inb2bwResult {
	if len(values) == 0 {
		return inb2bwResult{}
	}

	if xmath.UniqueN(values) == 1 {
		mnx := xutil.IfElse(len(values) == 0, 0, values[0])
		return inb2bwResult{
			min:      mnx,
			max:      mnx,
			binWidth: 1,
		}
	}

	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

	numBins := findNumBins(values)
	if numBins == 0 {
		return inb2bwResult{}
	}

	res := intNumBinsToBinWidth(values, numBins)

	return res
}

// Num-Bins 2 Bin-Width ////////////////////////////////////////////////////////

type inb2bwResult struct {
	min      int64
	max      int64
	binWidth int64
}

func intNumBinsToBinWidth(values []int64, numBins int) inb2bwResult {
	info := intInfo(values)

	binWidth := float64(info.max-info.min) / float64(numBins)

	if numBins == 1 {
		binWidth += 1
	}

	return inb2bwResult{
		min:      info.min,
		max:      info.max,
		binWidth: int64(math.Round(binWidth)),
	}
}

// Int Info ////////////////////////////////////////////////////////////////////

type intInfoResult struct {
	min       int64
	max       int64
	avgDigits int
}

func intInfo(values []int64) intInfoResult {
	sum := float64(0)
	low := values[0]
	high := values[0]

	for _, v := range values {
		sum += float64(intSize(v))
		if v < low {
			low = v
		} else if v > high {
			high = v
		}
	}

	return intInfoResult{low, high, int(math.Round(sum / float64(len(values))))}
}

// intSize determines the digit count of a given int64 value.
//
// @param value Value whose digit width should be returned.
//
// @returns Digit width of the given value.
func intSize(value int64) int {
	sign := 0

	if value < 0 {
		// In the off chance that the value is the one we can't flip to positive,
		// bail out here.
		if value == math.MinInt64 {
			return 20
		}

		sign = 1
		value = -value
	}

	switch true {
	case value < 10:
		return 1 + sign
	case value < 100:
		return 2 + sign
	case value < 1000:
		return 3 + sign
	case value < 10_000:
		return 4 + sign
	case value < 100_000:
		return 5 + sign
	case value < 1_000_000:
		return 6 + sign
	case value < 10_000_000:
		return 7 + sign
	case value < 100_000_000:
		return 8 + sign
	case value < 1_000_000_000:
		return 9 + sign
	case value < 10_000_000_000:
		return 10 + sign
	case value < 100_000_000_000:
		return 11 + sign
	case value < 1_000_000_000_000:
		return 12 + sign
	case value < 10_000_000_000_000:
		return 13 + sign
	case value < 100_000_000_000_000:
		return 14 + sign
	case value < 1_000_000_000_000_000:
		return 15 + sign
	case value < 10_000_000_000_000_000:
		return 16 + sign
	case value < 100_000_000_000_000_000:
		return 17 + sign
	case value < 1_000_000_000_000_000_000:
		return 18 + sign
	default:
		return 19 + sign
	}
}
