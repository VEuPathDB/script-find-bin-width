package input

import (
	"fmt"

	"find-bin-width/pkg/xstr"
	"find-bin-width/pkg/xtype"
)

type lineEntries struct {
	dataType xtype.DataType
	entries  any
	na       bool
}

func (e *lineEntries) append(key *lineKey, value string, rmNA bool) error {
	if e.na || e.dataType == xtype.DataTypeUnknown || !e.dataType.IsValid() {
		return nil
	}

	if len(value) == 0 {
		if rmNA {
			return nil
		}

		e.entries = nil
		e.na = true
		return nil
	}

	isInt := e.dataType == xtype.DataTypeInteger
	e.dataType = xtype.FindDataType(value, e.dataType)

	if e.dataType == xtype.DataTypeUnknown || !e.dataType.IsValid() {
		return fmt.Errorf("illegal value or mixed data types found for %s", key)
	}

	// Because it is possible for the data type int to transition to the data type
	// float, check for that here and convert the entries slice if needed.
	if isInt && e.dataType == xtype.DataTypeFloat {
		e.convertIntSliceToFloatSlice()
	}

	e.appendTyped(value)

	return nil
}

func (e *lineEntries) convertIntSliceToFloatSlice() {
	oldTyped := e.entries.([]int64)
	newTyped := make([]float64, len(oldTyped))

	for i, val := range oldTyped {
		newTyped[i] = float64(val)
	}

	e.entries = newTyped
}

func (e *lineEntries) appendTyped(value string) {
	switch e.dataType {
	case xtype.DataTypeInteger:
		e.entries = append(e.entries.([]int64), xstr.MustParseInt64(value))

	case xtype.DataTypeFloat:
		e.entries = append(e.entries.([]float64), xstr.MustParseFloat64(value))

	case xtype.DataTypeDate:
		e.entries = append(e.entries.([]int64), xstr.MustParseDate(value).Unix())

	case xtype.DataTypeBoolean:
		e.entries = append(e.entries.([]bool), xstr.MustParseBool(value))

	default:
		panic("illegal state: attempted to append a typed string with type " + e.dataType.String())
	}
}
