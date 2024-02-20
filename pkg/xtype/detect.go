package xtype

import (
	"regexp"
)

type DataType uint8

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
	case DataTypeBoolean:
		return "boolean"
	default:
		return "invalid data type"
	}
}

func (t DataType) IsValid() bool {
	return t < 6
}

var (
	boolRgx     = regexp.MustCompile(`^(?:[tTfFyYnN]|true|false|TRUE|FALSE|yes|no|YES|NO)$`)
	dateRgx     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	dateTimeRgx = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(?:T\d{2}:\d{2}:\d{2}(?:\.\d+(?:z|Z|[+-]\d{2}:\d{2}))?)?$`)
	floatRgx    = regexp.MustCompile(`^-?\d+(?:.\d+)?$`)
	intRgx      = regexp.MustCompile(`^-?\d+$`)
)

func FindDataType(value string, assumption DataType) DataType {
	if len(value) == 0 {
		return assumption
	}

	switch assumption {
	case DataTypeFloat:
		return stateFloat(value)
	case DataTypeInteger:
		return stateInteger(value)
	case DataTypeDate:
		return stateDate(value)
	case DataTypeBoolean:
		return stateBool(value)
	case DataTypeUndecided:
		return stateUndecided(value)
	default:
		return DataTypeUnknown
	}
}

func stateUndecided(value string) DataType {
	if intRgx.MatchString(value) {
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
