package xtype

import (
	"regexp"
)

type DataType byte

const (
	DataTypeUndecided DataType = 0
	DataTypeUnknown   DataType = 1
	DataTypeInteger   DataType = 2
	DataTypeFloat     DataType = 3
	DataTypeDate      DataType = 4
	DataTypeBoolean   DataType = 5
)

func (t DataType) String() string {
	switch t {
	case DataTypeUndecided:
		return "undecided"
	case DataTypeUnknown:
		return "unknown"
	case DataTypeInteger:
		return "integer"
	case DataTypeFloat:
		return "float"
	case DataTypeDate:
		return "date"
	default:
		return "invalid data type"
	}
}

var (
	boolRgx     = regexp.MustCompile(`^(?:[tTfFyYnN]|true|false|TRUE|FALSE|yes|no|YES|NO)$`)
	dateRgx     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	dateTimeRgx = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(?:T\d{2}:\d{2}:\d{2}(?:\.\d+(?:z|Z|[+-]\d{2}:\d{2}))?)?$`)
	floatRgx    = regexp.MustCompile(`^-?\d+(?:.\d+)?$`)
	intRgx      = regexp.MustCompile(`^-?\d+$`)
)

func FindDataType(values []string) DataType {
	if len(values) == 0 {
		return DataTypeUndecided
	}

	t := DataTypeUndecided

	for _, v := range values {

		if v == "" {
			continue
		}

		switch t {
		case DataTypeFloat:
			t = stateFloat(v)
		case DataTypeInteger:
			t = stateInteger(v)
		case DataTypeDate:
			t = stateDate(v)
		case DataTypeBoolean:
			t = stateBool(v)
		case DataTypeUndecided:
			t = stateUndecided(v)
		}

		if t == DataTypeUnknown {
			break
		}
	}

	return t
}

func stateUndecided(value string) DataType {
	if len(value) == 0 {
		return DataTypeUndecided
	} else if intRgx.MatchString(value) {
		return DataTypeInteger
	} else if floatRgx.MatchString(value) {
		return DataTypeFloat
	} else if dateRgx.MatchString(value) {
		return DataTypeDate
	} else if dateTimeRgx.MatchString(value) {
		return DataTypeDate
	} else if boolRgx.MatchString(value) {
		return DataTypeBoolean
	} else {
		return DataTypeUnknown
	}
}

func stateInteger(value string) DataType {
	if intRgx.MatchString(value) {
		return DataTypeInteger
	} else if floatRgx.MatchString(value) {
		return DataTypeFloat
	} else {
		return DataTypeUnknown
	}
}

func stateFloat(value string) DataType {
	if floatRgx.MatchString(value) {
		return DataTypeFloat
	} else if intRgx.MatchString(value) {
		return DataTypeFloat
	} else {
		return DataTypeUnknown
	}
}

func stateDate(value string) DataType {
	if dateRgx.MatchString(value) {
		return DataTypeDate
	} else if dateTimeRgx.MatchString(value) {
		return DataTypeDate
	} else {
		return DataTypeUnknown
	}
}

func stateBool(value string) DataType {
	if boolRgx.MatchString(value) {
		return DataTypeBoolean
	} else {
		return DataTypeUnknown
	}
}
