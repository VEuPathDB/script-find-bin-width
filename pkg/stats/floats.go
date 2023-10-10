package stats

import (
	"sort"
	"strconv"
	"strings"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
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
func FindFloatBinWidth(values []float64, rmNa bool) Stats {
	if len(values) == 0 {
		return Stats{
			Min:           "0",
			Max:           "0",
			BinWidth:      "0",
			Mean:          "0",
			Median:        "0",
			LowerQuartile: "0",
			UpperQuartile: "0",
		}
	}

	if rmNa {
		values = xtype.FloatsRemoveNAs(values)
	} else if xtype.FloatsContainNA(values) {
		return Stats{}
	}

	if xmath.UniqueN(values) == 1 {
		mnx := strconv.FormatFloat(xutil.IfElse(len(values) == 0, 0, values[0]), 'f', -1, 64)
		return Stats{
			Min:           mnx,
			Max:           mnx,
			BinWidth:      "0",
			Mean:          mnx,
			Median:        mnx,
			LowerQuartile: mnx,
			UpperQuartile: mnx,
		}
	}

	sort.Float64s(values)

	numBins := findNumBins(values)
	if numBins == 0 {
		return Stats{
			Min:           "0",
			Max:           "0",
			BinWidth:      "0",
			Mean:          "0",
			Median:        "0",
			LowerQuartile: "0",
			UpperQuartile: "0",
		}
	}

	res := floatNumBinsToBinWidth(values, numBins)

	return Stats{
		Min:           strconv.FormatFloat(res.min, 'f', -1, 64),
		Max:           strconv.FormatFloat(res.max, 'f', -1, 64),
		BinWidth:      strconv.FormatFloat(xmath.NonZeroRound(res.binWidth, res.avgDigits), 'f', -1, 64),
		Mean:          strconv.FormatFloat(xmath.Mean(values), 'f', -1, 64),
		Median:        strconv.FormatFloat(xmath.Median(values), 'f', -1, 64),
		LowerQuartile: strconv.FormatFloat(xmath.LowerQuartile(values), 'f', -1, 64),
		UpperQuartile: strconv.FormatFloat(xmath.UpperQuartile(values), 'f', -1, 64),
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
