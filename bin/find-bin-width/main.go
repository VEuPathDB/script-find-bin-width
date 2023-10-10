package main

import (
	"os"

	cli "github.com/Foxcapades/Argonaut/v0"

	"find-bin-width/pkg/stats"
	"find-bin-width/pkg/xos"
)

func main() {
	config := parseArgs()
	if config.file == "" {
		result := stats.Calculate(os.Stdin, config.rmNa)
		xos.RequireWriteString(os.Stdout, result.String())
	} else {
		file := xos.RequireOpen(config.file)
		defer file.Close()

		result := stats.Calculate(file, config.rmNa)
		xos.RequireWriteString(os.Stdout, result.String())
	}
}

type cliConfig struct {
	rmNa bool
	file string
}

func parseArgs() (conf cliConfig) {
	cli.NewCommand().
		Flag(cli.NewFlag().
			Long("rm-na").
			Short('r').
			Description("Whether NA values (empty strings on input) should be ignored.  If this is not set, or is set to false, data sets containing NA values will result in an NA value being returned.").
			Bind(&conf.rmNa, true)).
		Arg(cli.NewArg().
			Name("file").
			Description("File to read data from.  If omitted, data will be read from stdin.").
			Bind(&conf.file)).
		MustParse()

	return
}
