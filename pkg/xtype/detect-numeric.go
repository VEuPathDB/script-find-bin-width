package xtype

import "find-bin-width/pkg/xutil"

type detectNumberResult uint8

const (
	dnrNo detectNumberResult = iota
	dnrYes
	dnrMaybe
)

func attemptNumeric(value string) DataType {
	var r detectNumberResult
	var i int

	if r, i = _valueIsInt(value); r != dnrMaybe {
		if r == dnrNo {
			return DataTypeUnknown
		} else {
			return DataTypeInteger
		}
	}

	if r, i = _valueIsPlainFloat(value, i); r != dnrMaybe {
		if r == dnrNo {
			return DataTypeUnknown
		} else {
			return DataTypeFloat
		}
	}

	if _valueIsSciNoFloat(value, i) == dnrYes {
		return DataTypeFloat
	} else {
		return DataTypeUnknown
	}
}

func _valueIsInt(val string) (res detectNumberResult, idx int) {
	if val[idx] == '-' || val[idx] == '+' {
		idx++

		if idx == len(val) {
			return dnrNo, idx
		}
	}

	for ; idx < len(val); idx++ {
		if !xutil.CharIsNumeric(val[idx]) {
			return dnrMaybe, idx
		}
	}

	return dnrYes, idx
}

func _valueIsPlainFloat(val string, i int) (detectNumberResult, int) {
	d := false

	for ; i < len(val); i++ {
		if xutil.CharIsNumeric(val[i]) {
			continue
		}

		if val[i] == '.' {
			if d {
				return dnrNo, i
			}

			d = true
			continue
		}

		return dnrMaybe, i
	}

	if val[i-1] == '.' {
		return dnrNo, i
	}

	return dnrYes, i
}

func _valueIsSciNoFloat(val string, i int) detectNumberResult {
	l := len(val)

	if l < i+2 || (val[i] != 'e' && val[i] != 'E') {
		return dnrNo
	}

	i++

	if val[i] == '-' || val[i] == '+' {
		i++

		if l < i+2 {
			return dnrNo
		}
	}

	for ; i < l; i++ {
		if !xutil.CharIsNumeric(val[i]) {
			return dnrNo
		}
	}

	return dnrYes
}
