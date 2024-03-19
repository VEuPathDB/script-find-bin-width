package output

import (
	"bufio"
	"strconv"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xutil"
)

type JsonOutputStyle uint8

const (
	JsonOutputStyleArrayOfObjects JsonOutputStyle = iota
	JsonOutputStyleArrayOfArraysNoHeaderRow
	JsonOutputStyleArrayOfArraysWithHeaderRow
)

type JsonFormatConfig struct {
	OutputStyle     JsonOutputStyle
	FieldNameFormat FieldNameFormat
}

func JsonFormatter(out *bufio.Writer, config JsonFormatConfig) Formatter {
	switch config.OutputStyle {
	case JsonOutputStyleArrayOfObjects:
		return jsonAOO{
			jsonFormatCommon{
				nameFormat: config.FieldNameFormat,
				stream:     out,
			},
		}
	case JsonOutputStyleArrayOfArraysNoHeaderRow:
		return jsonAOANHR{
			jsonFormatCommon{
				nameFormat: config.FieldNameFormat,
				stream:     out,
			},
		}
	case JsonOutputStyleArrayOfArraysWithHeaderRow:
		return jsonAOAWHR{
			jsonFormatCommon: jsonFormatCommon{
				nameFormat: config.FieldNameFormat,
				stream:     out,
			},
			writeHeader: true,
		}
	default:
		panic("unrecognized json output style " + strconv.Itoa(int(config.OutputStyle)))
	}
}

type jsonAOO struct{ jsonFormatCommon }

func (j jsonAOO) Open() {
	xos.BufWriteByte(j.stream, '[')
}

func (j jsonAOO) Write(result Formattable) {
	j.writeObject(result)
}

func (j jsonAOO) Finalize() {
	xos.BufWriteByte(j.stream, ']')
	xutil.Must(j.stream.Flush())
}

type jsonAOANHR struct{ jsonFormatCommon }

func (j jsonAOANHR) Open() {
	xos.BufWriteByte(j.stream, '[')
}

func (j jsonAOANHR) Write(row Formattable) {
	j.writeValueArray(row)
}

func (j jsonAOANHR) Finalize() {
	xos.BufWriteByte(j.stream, ']')
	xutil.Must(j.stream.Flush())
}

type jsonAOAWHR struct {
	jsonFormatCommon
	writeHeader bool
}

func (j jsonAOAWHR) Open() {
	xos.BufWriteByte(j.stream, '[')
}

func (j jsonAOAWHR) Write(result Formattable) {
	if j.writeHeader {
		j.writeHeaderArray(result)
		j.writeHeader = false
	}
	j.writeValueArray(result)
}

func (j jsonAOAWHR) Finalize() {
	xos.BufWriteByte(j.stream, ']')
	xutil.Must(j.stream.Flush())
}
