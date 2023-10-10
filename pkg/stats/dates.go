package stats

import (
	"time"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
	"find-bin-width/pkg/xutil"
)

const dateFormat = "2006-01-02T15:04:05"

// FindDateBinWidth determines the bin width for a series of date values.
//
// Returns one of the string values "year", "month", "week", or "day";
// additionally, in the case of an NA value, returns an empty string.
//
// @param values Series of values for which the bin width should be calculated.
//
// @param rmNa Whether NA values should be removed from the input series.  If
// this value is false and the input series contains one or more NA values, an
// empty string will be returned.
//
// @returns A string representation of the bin width.  NA return values will be
// represented as an empty string.
func FindDateBinWidth(values []time.Time, rmNa bool) Stats {
	if rmNa {
		values = xtype.DatesRemoveNAs(values)
	} else if xtype.DatesContainNA(values) {
		return Stats{}
	}

	intValues := datesToInts(values)
	values = nil

	if xmath.UniqueN(intValues) == 1 {
		mnx := time.Unix(xutil.IfElse(len(values) == 0, 0, intValues[0]), 0).Format(dateFormat)
		return Stats{
			Min:           mnx,
			Max:           mnx,
			BinWidth:      "day",
			Mean:          mnx,
			Median:        mnx,
			LowerQuartile: mnx,
			UpperQuartile: mnx,
		}
	}

	res := intSafeFindBinWidth(intValues)
	binWidth := res.binWidth / 86400

	min := time.Unix(res.min, 0)
	max := time.Unix(res.max, 0)
	mea := time.Unix(int64(xmath.Ceil(xmath.Mean(intValues))), 0).Format(dateFormat)
	med := time.Unix(int64(xmath.Ceil(xmath.Median(intValues))), 0).Format(dateFormat)
	low := time.Unix(int64(xmath.Ceil(xmath.LowerQuartile(intValues))), 0).Format(dateFormat)
	upp := time.Unix(int64(xmath.Ceil(xmath.UpperQuartile(intValues))), 0).Format(dateFormat)

	if binWidth > 365 {
		return Stats{
			Min:           min.Format(dateFormat),
			Max:           max.Format(dateFormat),
			BinWidth:      "year",
			Mean:          mea,
			Median:        med,
			LowerQuartile: low,
			UpperQuartile: upp,
		}
	} else if binWidth > 31 {
		return Stats{
			Min:           min.Format(dateFormat),
			Max:           max.Format(dateFormat),
			BinWidth:      "month",
			Mean:          mea,
			Median:        med,
			LowerQuartile: low,
			UpperQuartile: upp,
		}
	} else if binWidth > 7 {
		return Stats{
			Min:           min.Format(dateFormat),
			Max:           max.Format(dateFormat),
			BinWidth:      "week",
			Mean:          mea,
			Median:        med,
			LowerQuartile: low,
			UpperQuartile: upp,
		}
	} else {
		return Stats{
			Min:           min.Format(dateFormat),
			Max:           max.Format(dateFormat),
			BinWidth:      "day",
			Mean:          mea,
			Median:        med,
			LowerQuartile: low,
			UpperQuartile: upp,
		}
	}
}

func datesToInts(values []time.Time) []int64 {
	out := make([]int64, len(values))

	for i := range values {
		out[i] = values[i].Unix()
	}

	return out
}
