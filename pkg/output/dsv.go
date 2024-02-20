package output

import (
	"bufio"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xutil"
)

type DSVFormatConfig struct {
	Delimiter     string
	WriteHeaders  bool
	HeaderFormat  FieldNameFormat
	LineSeparator string
}

func DSVFormatter(stream *bufio.Writer, config DSVFormatConfig) Formatter {
	return &dsvFormatter{
		out:        stream,
		writeHead:  config.WriteHeaders,
		headFormat: config.HeaderFormat,
		delim:      config.Delimiter,
		nl:         config.LineSeparator,
	}
}

type dsvFormatter struct {
	out        *bufio.Writer
	writeHead  bool
	headFormat FieldNameFormat
	delim      string
	nl         string
	keyBuf     FieldNameBuffer
	valBuf     FieldValueBuffer
}

func (d *dsvFormatter) Open() {}

func (d *dsvFormatter) Write(result Formattable) {
	fc := result.FieldCount()

	if fc < 1 {
		if d.writeHead {
			xos.BufWriteString(d.out, d.nl)
			d.writeHead = false
		}
		xos.BufWriteString(d.out, d.nl)
		return
	}

	if d.writeHead {
		for i := 0; i < fc; i++ {
			xos.BufWriteBytes(d.out, d.keyBuf[:d.headFormat.format(&d.keyBuf, result.GetFieldName(i, &d.keyBuf))])
		}
		xos.BufWriteString(d.out, d.nl)
		d.writeHead = false
	}

	xos.BufWriteBytes(d.out, d.valBuf[:result.GetFieldValue(0, &d.valBuf)])
	for i := 1; i < fc; i++ {
		xos.BufWriteString(d.out, d.delim)
		xos.BufWriteBytes(d.out, d.valBuf[:result.GetFieldValue(0, &d.valBuf)])
	}
	xos.BufWriteString(d.out, d.nl)
}

func (d *dsvFormatter) Finalize() {
	xutil.Must(d.out.Flush())
}
