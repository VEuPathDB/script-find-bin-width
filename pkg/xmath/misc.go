package xmath

func Max[V float64 | int64](a, b V) V {
	if a > b {
		return a
	}

	return b
}

func Min[V float64 | int64](a, b V) V {
	if a > b {
		return b
	}

	return a
}

func UniqueN[V float64 | int64](values []V) int {
	hold := make(map[V]bool)

	for _, v := range values {
		hold[v] = true
	}

	return len(hold)
}
