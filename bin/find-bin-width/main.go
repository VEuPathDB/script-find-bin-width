package main

import (
	"bufio"
	"os"
	"strconv"

	"find-bin-width/pkg/cmd"
	"find-bin-width/pkg/output"
	"find-bin-width/pkg/stats"
)

func main() {
	config := cmd.ParseCliArgs(os.Args)

	var result stats.ResultIterator
	var err error
	var file *os.File

	if config.InputFile == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(config.InputFile)
		if err != nil {
			badExit(err.Error())
		}
		defer func(file *os.File) { _ = file.Close() }(file)
	}

	if config.InputsAreSorted {
		result, err = stats.CalculateSorted(file, config.RemoveNAValues)
	} else {
		result, err = stats.CalculateUnsorted(file, config.RemoveNAValues)
	}

	if err != nil {
		badExit(err.Error())
	}

	var formatter output.Formatter

	writer := bufio.NewWriter(os.Stdout)

	switch config.OutputFormat {
	case output.FormatTSV:
		formatter = output.DSVFormatter(writer, output.DSVFormatConfig{
			Delimiter:     "\t",
			LineSeparator: "\n",
			WriteHeaders:  config.PrintHeaders,
		})
	case output.FormatCSV:
		formatter = output.DSVFormatter(writer, output.DSVFormatConfig{
			Delimiter:     ",",
			LineSeparator: "\n",
			WriteHeaders:  config.PrintHeaders,
		})
	case output.FormatJSON:
		formatter = output.JsonFormatter(writer, output.JsonFormatConfig{
			FieldNameFormat: output.FieldNameFormatCamel,
		})
	case output.FormatJSONL:
		formatter = output.JsonLinesFormatter(writer, output.JsonLinesFormatterConfig{
			FieldNameFormat:     output.FieldNameFormatCamel,
			LineSeparator:       "\n",
			PrintFieldNameArray: config.PrintHeaders,
		})
	default:
		badExit("illegal state: unrecognized config format " + strconv.Itoa(int(config.OutputFormat)))
	}

	formatter.Open()
	for result.HasNext() {
		if res, err := result.Next(); err != nil {
			badExit(err.Error())
		} else {
			formatter.Write(res)
		}
	}

	formatter.Finalize()
}

func badExit(msg string) {
	_, _ = os.Stderr.WriteString(msg)
	os.Exit(1)
}
