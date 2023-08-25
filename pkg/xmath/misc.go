package xmath

import "find-bin-width/pkg/xutil"

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
