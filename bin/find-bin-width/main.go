package main

import (
	"bufio"
	"io"
	"os"
	"strconv"

	"find-bin-width/pkg/cmd"
	"find-bin-width/pkg/output"
	"find-bin-width/pkg/stats"
)

func main() {
	config := cmd.ParseCliArgs(os.Args)

	formatter := buildFormatter(os.Stdout, &config)

	formatter.Open()

	if len(config.InputFiles) > 0 {
		for _, path := range config.InputFiles {
			if err := processInputFile(path, formatter, &config); err != nil {
				badExit(err.Error())
			}
		}
	} else {
		if err := processInput(os.Stdin, formatter, &config); err != nil {
			badExit(err.Error())
		}
	}

	formatter.Finalize()
}

func processInputFile(path string, formatter output.Formatter, config *cmd.Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	return processInput(file, formatter, config)
}

func processInput(input io.Reader, formatter output.Formatter, config *cmd.Config) error {

	var result stats.ResultIterator
	var err error

	if config.InputsAreSorted {
		result, err = stats.CalculateSorted(input, config.RemoveNAValues)
	} else {
		result, err = stats.CalculateUnsorted(input, config.RemoveNAValues)
	}

	if err != nil {
		return err
	}

	for result.HasNext() {
		if res, err := result.Next(); err != nil {
			return err
		} else {
			formatter.Write(res)
		}
	}

	return nil
}

func buildFormatter(out io.Writer, config *cmd.Config) output.Formatter {
	writer := bufio.NewWriter(out)

	switch config.OutputFormat {
	case output.FormatTSV:
		return output.DSVFormatter(writer, output.DSVFormatConfig{
			Delimiter:     "\t",
			LineSeparator: "\n",
			WriteHeaders:  config.PrintHeaders,
		})
	case output.FormatCSV:
		return output.DSVFormatter(writer, output.DSVFormatConfig{
			Delimiter:     ",",
			LineSeparator: "\n",
			WriteHeaders:  config.PrintHeaders,
		})
	case output.FormatJSON:
		return output.JsonFormatter(writer, output.JsonFormatConfig{
			FieldNameFormat: output.FieldNameFormatCamel,
		})
	case output.FormatJSONL:
		return output.JsonLinesFormatter(writer, output.JsonLinesFormatterConfig{
			FieldNameFormat:     output.FieldNameFormatCamel,
			LineSeparator:       "\n",
			PrintFieldNameArray: config.PrintHeaders,
		})
	default:
		badExit("illegal state: unrecognized config format " + strconv.Itoa(int(config.OutputFormat)))
		panic(nil)
	}
}

func badExit(msg string) {
	_, _ = os.Stderr.WriteString(msg)
	os.Exit(1)
}
