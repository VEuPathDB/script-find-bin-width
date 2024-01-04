package xmath

import (
	"math"
)

type StdDeviationResult struct {
	Mean         float64
	StdDeviation float64
}

type SkewnessResult struct {
	Mean         float64
	StdDeviation float64
	Skewness     float64
}

type MADResult struct {
	Mean float64
	MAD  float64
}

// //
//
//  GENERIC FUNCTIONS
//
// //

// FD computes the number of bins for a histogram using the Freedman-Diaconis
// method applied to the given set of observations.
func FD[V float64 | int64](values []V) float64 {
	iqr := IQR(values)

	if iqr == 0 {
		iqr = MAD(values).MAD
	}

	if iqr > 0 {
		return Ceil(float64(Diff(Range(values))) / (2 * iqr * math.Pow(float64(len(values)), -1.0/3.0)))
	}

	return 1
}

func UpperQuartile[V float64 | int64](values []V) float64 {
	ln := len(values)

	if ln%2 == 1 {
		values = values[ln/2+1:]
	} else {
		values = values[ln/2:]
	}

	ln = len(values)

	if ln%2 == 1 {
		return float64(values[ln/2])
	} else {
		return (float64(values[ln/2-1]) + float64(values[ln/2])) / 2
	}
}

func LowerQuartile[V float64 | int64](values []V) float64 {
	values = values[:len(values)/2]

	ln := len(values)

	if ln%2 == 1 {
		return float64(values[ln/2])
	} else {
		return (float64(values[ln/2-1]) + float64(values[ln/2])) / 2
	}
}

// IQR calculates the interquantile range of the given set of observations.
func IQR[V float64 | int64](values []V) float64 {
	midLeft := 0
	midRight := 0
	size := len(values)

	if size%2 == 0 {
		midLeft = size / 2
		midRight = midLeft
	} else {
		midLeft = size / 2
		midRight = midLeft + 1
	}

	q1 := Median(values[0:midLeft])
	q3 := Median(values[midRight:])

	return q3 - q1
}

// MAD calculates the mean absolute deviation of the given set of observations.
func MAD[V float64 | int64](values []V) MADResult {
	sum := 0.0
	mean := Mean(values)

	for _, v := range values {
		sum += math.Abs(float64(v) - mean)
	}

	return MADResult{mean, sum / float64(len(values))}
}

// Mean calculates the mean of the given set of observation.
func Mean[V float64 | int64](values []V) float64 {
	var sum V

	for _, v := range values {
		sum += v
	}

	return float64(sum) / float64(len(values))
}

// Median calculates the median value for the given set of observation.
func Median[V float64 | int64](values []V) float64 {
	size := len(values)

	if size%2 == 0 {
		h := size / 2
		a := float64(values[h])
		b := float64(values[h-1])
		return (a + b) / 2
	} else {
		return float64(values[size/2])
	}
}

// Skewness returns the skewness of the given set of observations.
func Skewness[V float64 | int64](values []V) SkewnessResult {
	sum := float64(0)
	res := StdDeviation(values)
	sd3 := res.StdDeviation * res.StdDeviation * res.StdDeviation

	for _, v := range values {
		tmp := float64(v) - res.Mean
		sum += tmp * tmp * tmp
	}

	skew := sum / ((float64(len(values)) - 1) * sd3)

	return SkewnessResult{res.Mean, res.StdDeviation, skew}
}

// StdDeviation returns the standard deviation of the given set of observations.
func StdDeviation[V float64 | int64](values []V) StdDeviationResult {
	sum := float64(0)
	mean := Mean(values)

	for _, v := range values {
		tmp := float64(v) - mean
		sum += tmp * tmp
	}

	std := math.Sqrt(sum / float64(len(values)-1))

	return StdDeviationResult{mean, std}
}

// Sturges computes the number of bins for a histogram using the Sturges
// method applied to the given set of observations.
func Sturges[V float64 | int64](values []V) float64 {
	res := Ceil(1 + math.Log2(float64(len(values))))
	return Max(res, 1)
}
