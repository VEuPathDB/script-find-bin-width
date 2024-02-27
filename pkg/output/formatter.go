package output

import (
	"fmt"
	"strconv"
	"strings"
)

// Format defines a type indicator for the various output formats implemented
// in this tool.
type Format uint8

const (
	FormatTSV Format = iota
	FormatCSV
	FormatJSON
	FormatJSONL
)

// ParseFormat attempts to match the given input string to one of the valid
// Format types.
//
// @param val = String value to match.
//
// @return[0] The matched Format value if one was found.
// @return[1] An error value if no matching Format value could be found.
func ParseFormat(val string) (Format, error) {
	switch strings.ToLower(val) {
	case "tsv":
		return FormatTSV, nil
	case "csv":
		return FormatCSV, nil
	case "json":
		return FormatJSON, nil
	case "jsonl":
		return FormatJSONL, nil
	default:
		return 0, fmt.Errorf("unrecognized output format \"%s\"", val)
	}
}

// Formatter defines a type used to write Formattable values out to a target or
// targets in an implementation specific Format.
type Formatter interface {

	// Open performs any steps necessary to prepare for writing individual values.
	//
	// An example of this could be writing the starting '[' character before
	// writing output values as a JSON array.
	Open()

	// Write the given value out to the underlying stream or streams.
	Write(value Formattable)

	// Finalize performs any steps necessary to conclude the output writing.
	//
	// An example of this could be writing the closing ']' character after writing
	// the final output value in a JSON array.
	Finalize()
}

// FieldNameFormat defines the different formats in which a Formattable field
// name may be written out by a Formatter instance.
type FieldNameFormat uint8

const (
	// FieldNameFormatSnake indicates that a field name value should be written in
	// snake case.
	//
	// Example: "my_field_name"
	FieldNameFormatSnake FieldNameFormat = iota

	// FieldNameFormatCamel indicates that a field name value should be written in
	// camel case.
	//
	// Example: "myFieldName"
	FieldNameFormatCamel

	// FieldNameFormatPascal indicates that a field name value should be written
	// in pascal case.
	//
	// Example: "MyFieldName"
	FieldNameFormatPascal

	// FieldNameFormatTitle indicates that a field name value should be written as
	// a title.
	//
	// Example: "My Field Name"
	FieldNameFormatTitle

	// FieldNameFormatKebab indicates that a field name value should be written in
	// kebab case.
	//
	// Example: "my-field-name"
	FieldNameFormatKebab

	// FieldNameFormatSentence indicates that a field name value should be written
	// in sentence format.
	//
	// Example: "My field name"
	FieldNameFormatSentence
)

func (f FieldNameFormat) format(buf *FieldNameBuffer, originalLength int) int {
	// Field name values are supposed to be in snake case automatically coming out
	// of the Formatter instances.
	if f == FieldNameFormatSnake {
		return originalLength
	}

	capNext := f == FieldNameFormatTitle || f == FieldNameFormatPascal
	writePos := 0
	for readPos := 0; readPos < originalLength; readPos++ {
		c := buf[readPos]

		if c == '_' {
			switch f {
			case FieldNameFormatTitle, FieldNameFormatSentence:
				buf[writePos] = ' '
			case FieldNameFormatKebab:
				buf[writePos] = '-'
			case FieldNameFormatCamel, FieldNameFormatPascal:
				writePos--
			default:
				panic("invalid field name format " + strconv.Itoa(int(f)))
			}
		} else if capNext {
			buf[writePos] = capitalize(c)
			capNext = false
		} else {
			buf[writePos] = c
		}

		writePos++
	}

	return writePos
}

func capitalize(b byte) byte {
	return b - 32
}
