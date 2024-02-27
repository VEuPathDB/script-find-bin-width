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
	overallDesc := "Calculates various stats about a given stream of input data.\n\n" +
		"Input is expected to be in a headerless 3 column TSV format with the columns 'attribute_stable_id', 'entity_id', " +
		"and 'value'.\n\n" +
		"Output will contain the fields 'min', 'max', 'bin_width', 'mean', 'median', 'lower_quartile', and " +
		"'upper_quartile'.\n\n" +
		"Input may be passed either on stdin, or via a list of 1 or more file paths which will be read and processed in " +
		"the order they are passed."

	rDesc := "Whether NA values (empty strings on input) should be ignored.  If this is not set, or is set to false, " +
		"data sets containing NA values will result in an NA value being returned.\n\nFor JSON output types, NA values " +
		"are represented as null.  For CSV/TSV output types, NA values are represented as empty strings."

	siDesc := "Whether the input is pre-sorted.  If not set, or set to false, input will be fully consumed and sorted " +
		"in memory before processing."

	fDesc := "Output format.  Valid options are tsv, csv, json, or jsonl"

	tDesc := "Whether the header/title line should be included in the output.  This option does not apply to json " +
		"output format."

	com := cli.Command().
		WithDescription(overallDesc).
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
		WithUnmappedLabel("input files...").
		MustParse(args)

	config.InputFiles = com.UnmappedInputs()

	return
}
