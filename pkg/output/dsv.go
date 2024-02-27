package output

import (
	"bufio"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xutil"
)

// DSVFormatConfig contains configuration values for initializing a Formatter
// instance that writes delimiter separated values out to a target stream.
type DSVFormatConfig struct {
	// Delimiter to use between written output values.
	Delimiter string

	// WriteHeaders indicates whether a header row should be printed at the start
	// of the DSV output.
	WriteHeaders bool

	// HeaderFormat defines the FieldNameFormat to use when writing out
	// Formattable field names.
	HeaderFormat FieldNameFormat

	// LineSeparator to use between written rows.
	LineSeparator string
}

// DSVFormatter returns a new Formatter instance that writes out Formattable
// values as delimiter separated value rows.
//
// @param stream = Output stream that Formattable values will be written to.
//
// @param config = Configuration for the new Formatter instance.
//
// @return A new Formatter instance that writes DSV rows to the given output
// stream.
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
		xos.BufWriteBytes(d.out, d.keyBuf[:d.headFormat.format(&d.keyBuf, result.WriteFieldName(0, &d.keyBuf))])
		for i := 1; i < fc; i++ {
			xos.BufWriteString(d.out, d.delim)
			xos.BufWriteBytes(d.out, d.keyBuf[:d.headFormat.format(&d.keyBuf, result.WriteFieldName(i, &d.keyBuf))])
		}
		xos.BufWriteString(d.out, d.nl)
		d.writeHead = false
	}

	xos.BufWriteBytes(d.out, d.valBuf[:result.WriteFieldValue(0, &d.valBuf)])
	for i := 1; i < fc; i++ {
		xos.BufWriteString(d.out, d.delim)
		xos.BufWriteBytes(d.out, d.valBuf[:result.WriteFieldValue(i, &d.valBuf)])
	}
	xos.BufWriteString(d.out, d.nl)
}

func (d *dsvFormatter) Finalize() {
	xutil.Must(d.out.Flush())
}
