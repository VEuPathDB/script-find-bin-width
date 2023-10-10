package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/Foxcapades/Argonaut/v0"

	"find-bin-width/pkg/stats"
	"find-bin-width/pkg/xos"
)

func main() {
	config := parseArgs()
	config.validate()

	var result stats.Stats

	if config.file == "" {
		result = stats.Calculate(os.Stdin, config.rmNa)
	} else {
		file := xos.RequireOpen(config.file)
		defer file.Close()

		result = stats.Calculate(file, config.rmNa)
	}

	switch config.format {
	case "tsv":
		printTSV(&result, &config)
	case "csv":
		printCSV(&result, &config)
	case "json":
		printJSON(&result)
	default:
		fmt.Print(result)
	}
}

func printCSV(s *stats.Stats, c *cliConfig) {
	if c.header {
		fmt.Println("Min,Max,BinWidth,Mean,Median,Q1,Q3")
	}
	fmt.Printf("%s,%s,%s,%s,%s,%s,%s", s.Min, s.Max, s.BinWidth, s.Mean, s.Median, s.LowerQuartile, s.UpperQuartile)
}

func printTSV(s *stats.Stats, c *cliConfig) {
	if c.header {
		fmt.Println("Min\tMax\tBinWidth\tMean\tMedian\tQ1\tQ3")
	}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s", s.Min, s.Max, s.BinWidth, s.Mean, s.Median, s.LowerQuartile, s.UpperQuartile)
}

func printJSON(s *stats.Stats) {
	if s.IsText {
		fmt.Printf(
			`{"min":"%s","max":"%s","binWidth":"%s","mean":"%s","median":"%s","q1":"%s","q3":"%s"}`,
			s.Min,
			s.Max,
			s.BinWidth,
			s.Mean,
			s.Median,
			s.LowerQuartile,
			s.UpperQuartile,
		)
	} else {
		fmt.Printf(
			`{"min":%s,"max":%s,"binWidth":%s,"mean":%s,"median":%s,"q1":%s,"q3":%s}`,
			s.Min,
			s.Max,
			s.BinWidth,
			s.Mean,
			s.Median,
			s.LowerQuartile,
			s.UpperQuartile,
		)
	}
}

type cliConfig struct {
	rmNa   bool
	file   string
	format string
	header bool
}

func (c cliConfig) validate() {
	switch c.format {
	case "tsv", "csv", "json":
		break
	default:
		log.Fatal("Unrecognized format value.  Must be one of tsv, csv, or json")
	}
}

func parseArgs() (conf cliConfig) {
	cli.NewCommand().
		Flag(cli.NewFlag().
			Long("rm-na").
			Short('r').
			Description("Whether NA values (empty strings on input) should be ignored.  If this is not set, or is set to false, data sets containing NA values will result in an NA value being returned.").
			Bind(&conf.rmNa, false)).
		Flag(cli.NewFlag().
			Long("format").
			Short('f').
			Description("Output format.  Valid options are tsv, csv, or json").
			Default("tsv").
			Bind(&conf.format, true)).
		Flag(cli.NewFlag().
			Long("headers").
			Short('t').
			Description("Whether the header/title line should be included in the output.  Only applies to tsv and csv formats, ignored for json.").
			Bind(&conf.header, false)).
		Arg(cli.NewArg().
			Name("file").
			Description("File to read data from.  If omitted, data will be read from stdin.").
			Bind(&conf.file)).
		MustParse()

	return
}
