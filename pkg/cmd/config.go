package cmd

import (
	"os"

	cli "github.com/Foxcapades/Argonaut"

	"find-bin-width/pkg/output"
)

type Config struct {
	RemoveNAValues  bool
	InputsAreSorted bool
	OutputFormat    output.Format
	PrintHeaders    bool
	InputFile       string
}

func ParseCliArgs(args []string) (config Config) {
	rDesc := "Whether NA values (empty strings on input) should be ignored.  If this is not set, or is set to false, " +
		"data sets containing NA values will result in an NA value being returned.\n\nFor JSON output types, NA values " +
		"are represented as null.  For CSV/TSV output types, NA values are represented as empty strings."

	siDesc := "Whether the input is pre-sorted.  If not set, or set to false, input will be fully consumed and sorted " +
		"in memory before processing."

	fDesc := "Output format.  Valid options are tsv, csv, json, or jsonl"

	iDesc := "File to read data from.  If omitted, data will be read from stdin."

	tDesc := "Whether the header/title line should be included in the output.  This option does not apply to json " +
		"output format."

	cli.Command().
		WithFlag(cli.ComboFlag('r', "rm-na").
			WithDescription(rDesc).
			WithBinding(&config.RemoveNAValues, false)).
		WithFlag(cli.ComboFlag('s', "sorted-inputs").
			WithDescription(siDesc).
			WithBindingAndDefault(&config.InputsAreSorted, false, true)).
		WithFlag(cli.ComboFlag('f', "format").
			WithDescription(fDesc).
			WithBindingAndDefault(func(val string) (err error) {
				config.OutputFormat, err = output.ParseFormat(val)
				return
			}, "tsv", true)).
		WithFlag(cli.ComboFlag('t', "headers").
			WithDescription(tDesc).
			WithBinding(&config.PrintHeaders, false)).
		WithArgument(cli.Argument().
			WithName("file").
			WithDescription(iDesc).
			WithBinding(func(path string) (err error) {
				config.InputFile = path
				_, err = os.Stat(path)
				return
			})).
		MustParse(args)

	return
}
