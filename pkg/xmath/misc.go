package xmath

import (
	"find-bin-width/pkg/xutil"
)

// Max returns the greater of two given values.
func Max[V float64 | int64](a, b V) V {
	return xutil.IfElse(a > b, a, b)
}

// Min returns the lesser of two given values.
func Min[V float64 | int64](a, b V) V {
	return xutil.IfElse(a < b, a, b)
}

// UniqueN returns the count of distinct values in a given set of values.
func UniqueN[V float64 | int64](values []V) int {
	hold := make(map[V]bool)

	for _, v := range values {
		hold[v] = true
	}

	return len(hold)
}

// Diff subtracts the first given value from the second given value.
func Diff[V float64 | int64](x, y V) V {
	return y - x
}

// Range returns the minimum and maximum values for a given set of numeric
// values.
func Range[V float64 | int64](values []V) (V, V) {
	mi := values[0]
	ma := values[0]

	for _, v := range values {
		if v < mi {
			mi = v
		} else if v > ma {
			ma = v
		}
	}

	return mi, ma
}
