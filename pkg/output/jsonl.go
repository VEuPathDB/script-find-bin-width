package output

import (
	"bufio"
	"strconv"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xutil"
)

type JsonLineFormat uint8

const (
	JsonLineFormatObjects JsonLineFormat = iota
	JsonLineFormatArrays
)

type JsonLinesFormatterConfig struct {
	FieldNameFormat     FieldNameFormat
	OutputFormat        JsonLineFormat
	PrintFieldNameArray bool
	LineSeparator       string
}

func JsonLinesFormatter(stream *bufio.Writer, config JsonLinesFormatterConfig) Formatter {
	if config.OutputFormat == JsonLineFormatObjects {
		return &jsonlObjectFormatter{
			jsonFormatCommon: jsonFormatCommon{
				nameFormat: config.FieldNameFormat,
				stream:     stream,
			},
			lineSep: config.LineSeparator,
		}
	}

	if config.OutputFormat == JsonLineFormatArrays {
		return &jsonlArrayFormatter{
			jsonFormatCommon: jsonFormatCommon{
				nameFormat: config.FieldNameFormat,
				stream:     stream,
			},
			writeHead: config.PrintFieldNameArray,
			lineSep:   config.LineSeparator,
		}
	}

	panic("unrecognized JsonLineFormat value: " + strconv.Itoa(int(config.OutputFormat)))
}

type jsonlObjectFormatter struct {
	jsonFormatCommon
	lineSep string
}

func (o jsonlObjectFormatter) Open() {}

func (o jsonlObjectFormatter) Write(result Formattable) {

	fc := result.FieldCount()

	if fc < 1 {
		xos.BufWriteString(o.stream, "{}")
		xos.BufWriteString(o.stream, o.lineSep)
		return
	}

	xos.BufWriteByte(o.stream, '{')

	o.writeKey(0, result)
	xos.BufWriteByte(o.stream, ':')
	o.writeValue(0, result)

	for i := 1; i < fc; i++ {
		xos.BufWriteByte(o.stream, ',')
		o.writeKey(i, result)
		xos.BufWriteByte(o.stream, ':')
		o.writeValue(i, result)
	}

	xos.BufWriteByte(o.stream, '}')
	xos.BufWriteString(o.stream, o.lineSep)
}

func (o jsonlObjectFormatter) Finalize() {
	// Do nothing.
}

type jsonlArrayFormatter struct {
	jsonFormatCommon
	writeHead bool
	lineSep   string
}

func (j *jsonlArrayFormatter) Open() {}

func (j *jsonlArrayFormatter) Write(result Formattable) {
	fc := result.FieldCount()

	if j.writeHead {
		xos.BufWriteByte(j.stream, '[')
		if fc > 0 {
			j.writeKey(0, result)
			for i := 1; i < fc; i++ {
				xos.BufWriteByte(j.stream, ',')
				j.writeKey(i, result)
			}
		}
		xos.BufWriteByte(j.stream, ']')
		xos.BufWriteString(j.stream, j.lineSep)
		j.writeHead = false
	}

	xos.BufWriteByte(j.stream, '[')
	if fc > 0 {
		j.writeValue(0, result)
		for i := 1; i < fc; i++ {
			xos.BufWriteByte(j.stream, ',')
			j.writeValue(i, result)
		}
	}
	xos.BufWriteByte(j.stream, ']')
	xos.BufWriteString(j.stream, j.lineSep)
}

func (j *jsonlArrayFormatter) Finalize() {
	xutil.Must(j.stream.Flush())
}
