package stats

import (
	"fmt"
	"unsafe"
)

type primTypeIndicator = uint8

const (
	primTypeFloat primTypeIndicator = (iota + 1) << 1
	primTypeInt
	primTypeBool
)

type NullablePrimitive[T float64 | int64 | bool] [9]byte

func (n *NullablePrimitive[T]) IsFloat() bool {
	return n[8]&primTypeFloat == primTypeFloat
}

func (n *NullablePrimitive[T]) IsInt() bool {
	return n[8]&primTypeInt == primTypeInt
}

func (n *NullablePrimitive[T]) IsBool() bool {
	return n[8]&primTypeBool == primTypeBool
}

func (n *NullablePrimitive[T]) IsNull() bool {
	return n[8]&1 == 0
}

func (n *NullablePrimitive[T]) errString(exp primTypeIndicator) string {
	if n.IsNull() {
		if n[8]&exp == exp {
			return "a null value"
		}

		switch true {
		case n.IsFloat():
			return "a null float"
		case n.IsInt():
			return "a null integer"
		case n.IsBool():
			return "a null boolean"
		default:
			panic("illegal state")
		}
	}

	switch true {
	case n.IsFloat():
		return "a float"
	case n.IsInt():
		return "an integer"
	case n.IsBool():
		return "a boolean"
	default:
		panic("illegal state")
	}
}

func (n *NullablePrimitive[T]) AsFloat() (float64, error) {
	if n.IsNull() || !n.IsFloat() {
		return 0, fmt.Errorf("cannot unwrap %s as a float value", n.errString(primTypeFloat))
	}

	tmp := [8]byte{}
	copy(tmp[:], n[:])

	return *(*float64)(unsafe.Pointer(&tmp)), nil
}

func (n *NullablePrimitive[T]) AsInt() (int64, error) {
	if n.IsNull() || !n.IsInt() {
		return 0, fmt.Errorf("cannot unwrap a %s as an integer value", n.errString(primTypeInt))
	}

	tmp := [8]byte{}
	copy(tmp[:], n[:])

	return *(*int64)(unsafe.Pointer(&tmp)), nil
}

func (n *NullablePrimitive[T]) AsBool() (bool, error) {
	if n.IsNull() || !n.IsBool() {
		return false, fmt.Errorf("cannot unwrap a %s as a boolean value", n.errString(primTypeBool))
	}

	return n[0] == 1, nil
}

func NewNullableFloat(value float64) (out NullablePrimitive[float64]) {
	tmp := *(*[8]byte)(unsafe.Pointer(&value))
	copy(out[:], tmp[:])
	out[8] = 1 | primTypeFloat
	return
}

func NewNullableInt(value int64) (out NullablePrimitive[int64]) {
	tmp := *(*[8]byte)(unsafe.Pointer(&value))
	copy(out[:], tmp[:])
	out[8] = 1 | primTypeInt
	return
}

func NewNullableBool(value bool) (out NullablePrimitive[bool]) {
	if value {
		out[0] = 1
	}
	out[8] = 1 | primTypeBool
	return
}
