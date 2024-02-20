package stats

import (
	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

// findDateBinWidth determines the bin width for a series of date values.
//
// Returns one of the string values "year", "month", "week", or "day";
// additionally, in the case of an NA value, returns an empty string.
//
// @param values Series of values for which the bin width should be calculated.
//
// @returns A string representation of the bin width.  NA return values will be
// represented as an empty string.
func findDateBinWidth(values []int64) stats[int64] {
	if xmath.UniqueN(values) == 1 {
		mnx := xutil.IfElse(len(values) == 0, 0, values[0])
		return stats[int64]{
			min:           mnx,
			max:           mnx,
			binWidth:      dateBinWidthDay,
			mean:          mnx,
			median:        mnx,
			lowerQuartile: mnx,
			upperQuartile: mnx,
			stringifier:   dateStringifier,
			dataType:      xtype.DataTypeDate,
		}
	}

	res := intFindBinWidth(values)
	binWidth := res.binWidth / 86400

	mea := int64(xmath.Ceil(xmath.Mean(values)))
	med := int64(xmath.Ceil(xmath.Median(values)))
	low := int64(xmath.Ceil(xmath.LowerQuartile(values)))
	upp := int64(xmath.Ceil(xmath.UpperQuartile(values)))

	bin := dateBinWidthDay
	switch true {
	case binWidth > 365:
		bin = dateBinWidthYear
	case binWidth > 31:
		bin = dateBinWidthMonth
	case binWidth > 7:
		bin = dateBinWidthWeek
	}

	return stats[int64]{
		min:           res.min,
		max:           res.max,
		binWidth:      bin,
		mean:          mea,
		median:        med,
		lowerQuartile: low,
		upperQuartile: upp,
		stringifier:   dateStringifier,
		dataType:      xtype.DataTypeDate,
	}
}
