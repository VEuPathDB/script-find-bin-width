package xtype

import (
	"math"
	"time"
)

var (
	NaInt  int64 = math.MinInt64
	NaTime       = time.Date(1, 01, 01, 0, 0, 0, 0, time.UTC)
)

// IntsContainNA tests whether the given slice of int64 values contain an NA
// value.
//
// NA values are represented as the minimum int64 value -9223372036854775808.
//
// @param values Slice of integers to test.
//
// @returns true if the slice of int64 values contains one or more NA values,
// otherwise false.
func IntsContainNA(values []int64) bool {
	for _, v := range values {
		if v == NaInt {
			return true
		}
	}

	return false
}

// IntsRemoveNAs returns a new slice of int64 values based on the given input
// slice, omitting any NA values present in the input slice.
//
// NA values are represented as the minimum int64 value -9223372036854775808.
//
// @param values Source integer slice.
//
// @returns A new slice of integers based on the source slice with NA values
// omitted.
func IntsRemoveNAs(values []int64) []int64 {
	out := make([]int64, 0, len(values))

	for _, v := range values {
		if v != NaInt {
			out = append(out, v)
		}
	}

	return out
}

// FloatsContainNA tests whether the given slice of float64 values contains an
// NA value.
//
// NA values are represented as the float64 value NaN.
//
// @param values Slice of floats to test.
//
// @returns true if the slice of float64 values contains one or more NA values,
// otherwise false.
func FloatsContainNA(values []float64) bool {
	for _, v := range values {
		if math.IsNaN(v) {
			return true
		}
	}

	return false
}

// FloatsRemoveNAs returns a new slice of float64 values based on the given
// input slice, omitting any NA values present in the input slice.
//
// NA values are represented as the float64 value NaN.
//
// @param values Source float slice.
//
// @returns A new slice of floats based on the source slice with NA values
// omitted.
func FloatsRemoveNAs(values []float64) []float64 {
	out := make([]float64, 0, len(values))

	for _, v := range values {
		if !math.IsNaN(v) {
			out = append(out, v)
		}
	}

	return out
}

// DatesContainNA tests whether the given slice of date values contains an NA
// value.
//
// NA values are represented as the date 0001-01-01T00:00:00.0Z.
//
// @param values Slice of dates to test.
//
// @returns true if the given slice contains one or more NA values, otherwise
// false.
func DatesContainNA(values []time.Time) bool {
	for i := range values {
		if values[i] == NaTime {
			return true
		}
	}

	return false
}

// DatesRemoveNAs returns a new slice of date values based on the given input
// slice, omitting any NA values present in the input slice.
//
// NA values are represented as the date 0001-01-01T00:00:00.0Z.
//
// @param values Source date slice.
//
// @returns A new slice of dates based on the source slice with NA values
// omitted.
func DatesRemoveNAs(values []time.Time) []time.Time {
	out := make([]time.Time, 0, len(values))

	for i := range values {
		if values[i] != NaTime {
			out = append(out, values[i])
		}
	}

	return out
}
