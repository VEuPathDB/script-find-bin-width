package xtype

import (
	"math"
	"time"

	"find-bin-width/pkg/xstr"
)

func ToBools(values []string) []PseudoBool {
	out := make([]PseudoBool, len(values))

	for i := range values {
		if values[i] == "" {
			out[i] = BoolNA
		} else {
			switch values[i][0] {

			case 't', 'T', 'y', 'Y', '1':
				out[i] = BoolTrue
				break

			case 'f', 'F', 'n', 'N', '0':
				out[i] = BoolFalse

			default:
				panic("unexpected non-boolean value: " + values[i])
			}
		}
	}

	return out
}

func ToDates(values []string) []time.Time {
	out := make([]time.Time, len(values))

	for i := range values {
		if values[i] == "" {
			out[i] = NaTime
		} else {
			out[i] = xstr.MustParseDate(values[i])
		}
	}

	return out
}

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
			out[i] = NaInt
		} else {
			out[i] = xstr.MustParseInt64(values[i])
		}
	}

	return out
}
