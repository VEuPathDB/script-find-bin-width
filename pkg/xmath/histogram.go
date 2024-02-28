package xmath

import "math"

// FD computes the number of bins for a histogram using the Freedman-Diaconis
// method applied to the given set of observations.
func FD[V float64 | int64](values []V) float64 {
	iqr := IQR(values)

	if iqr == 0 {
		iqr = MAD(values).MAD
	}

	if iqr > 0 {
		return math.Ceil(float64(Diff(Range(values))) / (2 * iqr * math.Pow(float64(len(values)), -1.0/3.0)))
	}

	return 1
}

// Sturges computes the number of bins for a histogram using the Sturges
// method applied to the given set of observations.
func Sturges[V float64 | int64](values []V) float64 {
	res := math.Ceil(1 + math.Log2(float64(len(values))))
	return max(res, 1)
}

// Doane implements Doane's formula for calculating the number of bins for a
// histogram.
func Doane[V float64 | int64](skewness float64, values []V) float64 {
	absSkew := math.Abs(skewness)

	n := float64(len(values))
	sd := math.Sqrt(6 * (n - 2) / ((n + 1) * (n + 3)))
	ke := math.Log2(1 + absSkew/sd)

	return math.Ceil(Sturges(values) + ke)
}
