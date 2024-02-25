package cmd

import (
	"fmt"
	"os"

	cli "github.com/Foxcapades/Argonaut"

	"find-bin-width/pkg/output"
)

type Config struct {
	RemoveNAValues  bool
	InputsAreSorted bool
	OutputFormat    output.Format
	PrintHeaders    bool
	InputFiles      []string
}

func ParseCliArgs(args []string) (config Config) {
	rDesc := "Whether NA values (empty strings on input) should be ignored.  If this is not set, or is set to false, " +
		"data sets containing NA values will result in an NA value being returned.\n\nFor JSON output types, NA values " +
		"are represented as null.  For CSV/TSV output types, NA values are represented as empty strings."

	siDesc := "Whether the input is pre-sorted.  If not set, or set to false, input will be fully consumed and sorted " +
		"in memory before processing."

	fDesc := "Output format.  Valid options are tsv, csv, json, or jsonl"

	tDesc := "Whether the header/title line should be included in the output.  This option does not apply to json " +
		"output format."

	com := cli.Command().
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
				if err != nil {
					_, _ = fmt.Fprint(os.Stderr, err.Error())
				}
				return
			}, "tsv", true)).
		WithFlag(cli.ComboFlag('t', "headers").
			WithDescription(tDesc).
			WithBinding(&config.PrintHeaders, false)).
		WithUnmappedLabel("input files").
		MustParse(args)

	config.InputFiles = com.UnmappedInputs()

	return
}
