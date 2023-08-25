package bw

func ifElse[V interface{}](cond bool, a V, b V) V {
	if cond {
		return a
	} else {
		return b
	}
}
