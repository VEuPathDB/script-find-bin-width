package xmath

import "math"

func Skewness[V float64 | int64](x []V) (skew, mean float64) {
	n := float64(len(x))
	sum3 := 0.0
	sum2 := 0.0

	mean = Mean(x)

	for _, f := range x {
		tmp := float64(f) - mean
		sum3 += tmp * tmp * tmp
		sum2 += tmp * tmp
	}

	return (sum3 / n) / math.Pow(sum2/n, 3.0/2.0), mean
}
