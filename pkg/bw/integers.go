package bw

import (
	"math"
	"sort"
	"strconv"

	"find-bin-width/pkg/xmath"
)

// FindIntegerBinWidth determines the bin width for a series of integral values.
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
func FindIntegerBinWidth(values []int64, rmNa bool) string {
	out := intFindBinWidth(values, rmNa)

	if out == NaInt {
		return ""
	}

	return strconv.FormatInt(out, 10)
}

func intFindBinWidth(values []int64, rmNa bool) int64 {
	if len(values) == 0 {
		return 0
	}

	if rmNa {
		values = intsRemoveNAs(values)
	} else if intsContainNA(values) {
		return NaInt
	}

	return intSafeFindBinWidth(values)
}

func intSafeFindBinWidth(values []int64) int64 {
	if xmath.UniqueN(values) == 1 {
		return 1
	}

	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

	numBins := findNumBins(values)
	if numBins == 0 {
		return 0
	}

	res := intNumBinsToBinWidth(values, numBins)

	return res.bw
}

type inb2bwResult struct {
	avg int
	bw  int64
}

func intNumBinsToBinWidth(values []int64, numBins int) inb2bwResult {
	info := intInfo(values)

	binWidth := float64(info.max-info.min) / float64(numBins)

	return inb2bwResult{info.avg, xmath.CeilToInt(binWidth)}
}

type intInfoResult struct {
	min int64
	max int64
	avg int
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
