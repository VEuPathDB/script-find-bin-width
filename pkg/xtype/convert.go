package xtype

import (
	"math"
	"time"

	"find-bin-width/pkg/bw"
	"find-bin-width/pkg/xstr"
)

func ToFloats(values []string) []float64 {
	out := make([]float64, len(values))

	for i := range values {
		if values[i] == "" {
			out[i] = math.NaN()
		} else {
			out[i] = xstr.MustParseFloat64(values[i])
		}
	}

	return out
}

func ToIntegers(values []string) []int64 {
	out := make([]int64, len(values))

	for i := range values {
		if values[i] == "" {
			out[i] = bw.NaInt
		} else {
			out[i] = xstr.MustParseInt64(values[i])
		}
	}

	return out
}

func ToDates(values []string) []time.Time {
	out := make([]time.Time, len(values))

	for i := range values {
		if values[i] == "" {
			out[i] = bw.NaTime
		} else {
			out[i] = xstr.MustParseDate(values[i])
		}
	}

	return out
}
