package output

import (
	"bufio"
	"strconv"

	"find-bin-width/pkg/xos"
)

type jsonFormatCommon struct {
	nameFormat FieldNameFormat
	stream     *bufio.Writer
	keyBuf     FieldNameBuffer
	valBuf     FieldValueBuffer
}

func (j jsonFormatCommon) writeHeaderArray(row Formattable) {
	fc := row.FieldCount()

	if fc < 1 {
		xos.BufWriteString(j.stream, "[]")
	} else {
		xos.BufWriteByte(j.stream, '[')
		j.writeKey(0, row)
		for i := 1; i < fc; i++ {
			xos.BufWriteByte(j.stream, ',')
			j.writeKey(i, row)
		}
		xos.BufWriteByte(j.stream, ']')
	}
}

func (j jsonFormatCommon) writeValueArray(row Formattable) {
	fc := row.FieldCount()

	if fc < 1 {
		xos.BufWriteString(j.stream, "[]")
	} else {
		xos.BufWriteByte(j.stream, '[')
		j.writeKey(0, row)
		for i := 1; i < fc; i++ {
			xos.BufWriteByte(j.stream, ',')
			j.writeValue(i, row)
		}
		xos.BufWriteByte(j.stream, ']')
	}
}

func (j jsonFormatCommon) writeObject(row Formattable) {
	fc := row.FieldCount()

	if fc < 1 {
		xos.BufWriteString(j.stream, "{}")
	} else {
		xos.BufWriteByte(j.stream, '{')
		j.writeKey(0, row)
		xos.BufWriteByte(j.stream, ':')
		j.writeValue(0, row)
		for i := 1; i < fc; i++ {
			xos.BufWriteByte(j.stream, ',')
			j.writeKey(i, row)
			xos.BufWriteByte(j.stream, ':')
			j.writeValue(i, row)
		}
		xos.BufWriteByte(j.stream, '}')
	}
}

func (j jsonFormatCommon) writeKey(index int, row Formattable) {
	j.writeQuoted(j.keyBuf[:j.nameFormat.format(&j.keyBuf, row.WriteFieldName(index, &j.keyBuf))])
}

func (j jsonFormatCommon) writeValue(index int, row Formattable) {
	switch row.GetFieldType(index) {
	case FieldTypeNumeric:
		xos.BufWriteBytes(j.stream, j.valBuf[:row.WriteFieldValue(index, &j.valBuf)])
	case FieldTypeText:
		j.writeQuoted(j.valBuf[:row.WriteFieldValue(index, &j.valBuf)])
	case FieldTypeBoolean:
		xos.BufWriteBytes(j.stream, j.valBuf[:row.WriteFieldValue(index, &j.valBuf)])
	case FieldTypeNA:
		xos.BufWriteString(j.stream, "null")
	default:
		panic("unrecognized field type: " + strconv.Itoa(int(row.GetFieldType(index))))
	}
}

func (j jsonFormatCommon) writeQuoted(buf []byte) {
	xos.BufWriteByte(j.stream, '"')
	xos.BufWriteBytes(j.stream, buf)
	xos.BufWriteByte(j.stream, '"')
}
