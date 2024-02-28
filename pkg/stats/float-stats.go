package stats

import (
	"sort"
	"strconv"
	"strings"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

func calculateFloatSummary(values []float64) Summary {
	if len(values) == 0 {
		return summary[float64]{stringifier: floatStringifier}
	}

	if xmath.UniqueN(values) == 1 {
		mnx := values[0]
		return summary[float64]{
			min:           mnx,
			max:           mnx,
			binWidth:      1,
			mean:          mnx,
			median:        mnx,
			lowerQuartile: NullablePrimitive[float64]{},
			upperQuartile: NullablePrimitive[float64]{},
			stringifier:   floatStringifier,
			dataType:      xtype.DataTypeFloat,
		}
	}

	sort.Float64s(values)

	numBins := findNumBins(values)
	if numBins == 0 {
		return summary[float64]{stringifier: floatStringifier}
	}

	res := floatNumBinsToBinWidth(values, numBins)

	var lq NullablePrimitive[float64]
	var uq NullablePrimitive[float64]

	if len(values) > 3 {
		lq = NewNullableFloat(xmath.LowerQuartile(values))
		uq = NewNullableFloat(xmath.UpperQuartile(values))
	}

	return summary[float64]{
		min:           res.min,
		max:           res.max,
		binWidth:      xmath.NonZeroRound(res.binWidth, res.avgDigits),
		mean:          xmath.Mean(values),
		median:        xmath.Median(values),
		lowerQuartile: lq,
		upperQuartile: uq,
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

	return fnb2bwResult{min: info.min, max: info.max, avgDigits: numDigits, binWidth: binWidth}
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
