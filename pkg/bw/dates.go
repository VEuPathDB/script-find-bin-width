package bw

import (
	"time"

	"find-bin-width/pkg/xmath"
	"find-bin-width/pkg/xtype"
)

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
func FindDateBinWidth(values []time.Time, rmNa bool) string {
	if rmNa {
		values = xtype.DatesRemoveNAs(values)
	} else if xtype.DatesContainNA(values) {
		return ""
	}

	intValues := datesToInts(values)
	values = nil

	if xmath.UniqueN(intValues) == 1 {
		return "day"
	}

	binWidth := intSafeFindBinWidth(intValues) / 86400

	if binWidth > 365 {
		return "year"
	} else if binWidth > 31 {
		return "month"
	} else if binWidth > 7 {
		return "week"
	} else {
		return "day"
	}
}

func datesToInts(values []time.Time) []int64 {
	out := make([]int64, len(values))

	for i := range values {
		out[i] = values[i].Unix()
	}

	return out
}
