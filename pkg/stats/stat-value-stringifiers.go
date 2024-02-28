package stats

import (
	"strconv"
	"time"

	"find-bin-width/pkg/output"
	"find-bin-width/pkg/xutil"
)

// valueStringifier defines a function type that converts the given raw value
// for a specified summary field to a string and writes it to the given buffer.
type valueStringifier = func(buf *output.FieldValueBuffer, raw any, field summaryField) int

// intStringifier converts the given value into a string and writes it to the
// given buffer.
//
// As some of the fields for int data sets are expected to be floats, all values
// stored in the summary struct for ints are stored as floats.  This means that
// for fields that expect int output, the value must be converted to int, and
// for fields that expect a float output, the value can be stringified directly.
//
// @param buf = Buffer that the stringified field value will be written to.
//
// @param raw = Raw field value that will be stringified.
//
// @param field = Identifier for the specific summary field the raw value is
// being stringified for.
//
// @return The number of bytes written to the given buffer.
func intStringifier(buf *output.FieldValueBuffer, raw any, field summaryField) int {
	switch field {

	case summaryFieldMin, summaryFieldMax, summaryFieldBinWidth:
		return copy(buf[:], strconv.FormatInt(int64(raw.(float64)), 10))

	case summaryFieldLowerQuartile, summaryFieldUpperQuartile:
		field := raw.(NullablePrimitive[float64])
		if field.IsNull() {
			return 0
		}

		return copy(buf[:], strconv.FormatFloat(xutil.MustReturn(field.AsFloat()), 'f', -1, 64))

	default:
		return copy(buf[:], strconv.FormatFloat(raw.(float64), 'f', -1, 64))
	}
}

// floatStringifier converts the given value into a string and writes it to the
// given buffer.
//
// @param buf = Buffer that the stringified field value will be written to.
//
// @param raw = Raw field value that will be stringified.
//
// @param field = Identifier for the specific summary field the raw value is
// being stringified for.
//
// @return The number of bytes written to the given buffer.
func floatStringifier(buf *output.FieldValueBuffer, raw any, field summaryField) int {
	if field == summaryFieldLowerQuartile || field == summaryFieldUpperQuartile {
		field := raw.(NullablePrimitive[float64])
		if field.IsNull() {
			return 0
		}

		return copy(buf[:], strconv.FormatFloat(xutil.MustReturn(field.AsFloat()), 'f', -1, 64))
	}

	return copy(buf[:], strconv.FormatFloat(raw.(float64), 'f', -1, 64))
}

const (
	// dateFormat defines the output format for date values.
	dateFormat = "2006-01-02"

	// sDay defines the string representation of the day bin width.
	sDay = "day"

	// sWeek defines the string representation of the week bin width.
	sWeek = "week"

	// sMonth defines the string representation of the month bin width.
	sMonth = "month"

	// sYear defines the string representation of the year bin width.
	sYear = "year"
)

// dateStringifier converts the given value into a string and writes it to the
// given buffer.
//
// @param buf = Buffer that the stringified field will be written to.
//
// @param raw = Raw field value that will be stringified.
//
// @param field = Identifier for the specific summary field the raw value is
// being stringified for.
//
// @return The number of bytes written to the given buffer.
func dateStringifier(buf *output.FieldValueBuffer, raw any, f summaryField) int {
	if f == summaryFieldBinWidth {
		switch raw.(int64) {
		case dateBinWidthDay:
			return copy(buf[:], sDay)
		case dateBinWidthWeek:
			return copy(buf[:], sWeek)
		case dateBinWidthMonth:
			return copy(buf[:], sMonth)
		case dateBinWidthYear:
			return copy(buf[:], sYear)
		default:
			panic("invalid date bin width type")
		}
	}

	if f == summaryFieldLowerQuartile || f == summaryFieldUpperQuartile {
		field := raw.(NullablePrimitive[int64])
		if field.IsNull() {
			return 0
		}

		return copy(buf[:], time.Unix(xutil.MustReturn(field.AsInt()), 0).Format(dateFormat))
	}

	return copy(buf[:], time.Unix(raw.(int64), 0).Format(dateFormat))
}
