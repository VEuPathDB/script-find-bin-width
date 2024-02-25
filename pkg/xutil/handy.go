package xutil

import (
	"fmt"
	"math"
)

// IfElse returns either the value of "ifTrue" or "ifFalse" depending on whether
// the given "condition" value is true or false.
//
// @param condition A boolean condition.
//
// @param ifTrue The value to be returned if the given condition is true.
//
// @param ifFalse The value to be returned if the given condition is false.
//
// @returns Either the value of "ifTrue" or the value of "ifFalse" depending on
// the boolean value of "condition".
func IfElse[V interface{}](condition bool, ifTrue V, ifFalse V) V {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustReturn[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func StringToUint16(value string) (out uint16) {
	l := len(value)

	for i := 0; i < min(l, 4); i++ {
		out *= 10
		out += uint16(value[i] - '0')
	}

	if l == 5 {
		if math.MaxUint16/10 < out {
			panic(fmt.Sprintf("value %s would overflow the type uint16", value))
		}

		out *= 10
		tmp := uint16(value[4] - '0')

		if math.MaxUint16-tmp < out {
			panic(fmt.Sprintf("value %s would overflow the type uint16", value))
		}

		out += tmp
	} else if l > 5 {
		panic(fmt.Sprintf("value %s would overflow the type uint16", value))
	}

	return
}

func CharIsNumeric(b byte) bool {
	return !(b < '0' || b > '9')
}

func CharIsNumberNoMoreThan(b, max byte) bool {
	return !(b < '0' || b > max)
}
