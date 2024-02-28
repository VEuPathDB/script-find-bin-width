package stats

import (
	"math"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

type dateBinWidth = int64

const (
	dateBinWidthDay dateBinWidth = iota
	dateBinWidthWeek
	dateBinWidthMonth
	dateBinWidthYear
)

func calculateDateSummary(values []int64) summary[int64] {
	if xmath.UniqueN(values) == 1 {
		mnx := xutil.IfElse(len(values) == 0, 0, values[0])
		return summary[int64]{
			min:           mnx,
			max:           mnx,
			binWidth:      dateBinWidthDay,
			mean:          mnx,
			median:        mnx,
			lowerQuartile: NullablePrimitive[int64]{},
			upperQuartile: NullablePrimitive[int64]{},
			stringifier:   dateStringifier,
			dataType:      xtype.DataTypeDate,
		}
	}

	res := sharedCalcIntegerStats(values)
	binWidth := res.binWidth / 86400

	mea := int64(math.Ceil(xmath.Mean(values)))
	med := int64(math.Ceil(xmath.Median(values)))

	var low NullablePrimitive[int64]
	var upp NullablePrimitive[int64]

	if len(values) > 3 {
		low = NewNullableInt(int64(math.Ceil(xmath.LowerQuartile(values))))
		upp = NewNullableInt(int64(math.Ceil(xmath.UpperQuartile(values))))
	}

	bin := dateBinWidthDay
	switch true {
	case binWidth > 365:
		bin = dateBinWidthYear
	case binWidth > 31:
		bin = dateBinWidthMonth
	case binWidth > 7:
		bin = dateBinWidthWeek
	}

	return summary[int64]{
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
