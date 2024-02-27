package xtype

import (
	"fmt"
)

func FindDataType(value string, assumption DataType) (DataType, error) {
	if len(value) == 0 {
		return assumption, nil
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
		return DataTypeUnknown, fmt.Errorf("attempted to find data type for value while in an unknown state")
	}
}

func stateUndecided(value string) (DataType, error) {
	if dt := attemptNumeric(value); dt.IsNumeric() {
		return dt, nil
	}

	if valueIsDate(value) {
		return DataTypeDate, nil
	}

	if valueIsBool(value) {
		return DataTypeBoolean, nil
	}

	return DataTypeUnknown, fmt.Errorf(`could not parse value "%s" as any of the supported types`, value)
}

func stateInteger(value string) (DataType, error) {
	dt := attemptNumeric(value)

	if !dt.IsNumeric() {
		return DataTypeUnknown, fmt.Errorf(`could not parse value "%s" as an int or float value`, value)
	}

	return dt, nil
}

func stateFloat(value string) (DataType, error) {
	if attemptNumeric(value).IsNumeric() {
		return DataTypeFloat, nil
	}

	return DataTypeUnknown, fmt.Errorf(`could not parse value "%s" as a float value`, value)
}

func stateDate(value string) (DataType, error) {
	if valueIsDate(value) {
		return DataTypeDate, nil
	}

	return DataTypeUnknown, fmt.Errorf(`could not parse value "%s" as a date value`, value)
}

func stateBool(value string) (DataType, error) {
	if valueIsBool(value) {
		return DataTypeBoolean, nil
	}

	return DataTypeUnknown, fmt.Errorf(`could not parse value "%s" as a boolean value`, value)
}
