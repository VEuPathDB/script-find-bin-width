package output

import (
	"fmt"
	"strconv"
	"strings"
)

type Format uint8

const (
	FormatTSV Format = iota
	FormatCSV
	FormatJSON
	FormatJSONL
)

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

type Formatter interface {
	Open()
	Write(result Formattable)
	Finalize()
}

type FieldNameFormat uint8

const (
	FieldNameFormatSnake FieldNameFormat = iota
	FieldNameFormatCamel
	FieldNameFormatPascal
	FieldNameFormatTitle
	FieldNameFormatKebab
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
