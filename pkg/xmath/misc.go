package xmath

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
