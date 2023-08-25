package xmath

import (
	"find-bin-width/pkg/xutil"
)

func Max[V float64 | int64](a, b V) V {
	return xutil.IfElse(a > b, a, b)
}

func Min[V float64 | int64](a, b V) V {
	return xutil.IfElse(a < b, a, b)
}

func UniqueN[V float64 | int64](values []V) int {
	hold := make(map[V]bool)

	for _, v := range values {
		hold[v] = true
	}

	return len(hold)
}

func Diff[V float64 | int64](x, y V) V {
	return y - x
}

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
