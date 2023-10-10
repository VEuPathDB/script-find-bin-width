package stats

import (
	"fmt"
	"io"
	"log"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xtype"
)

type Stats struct {
	Min           string
	Max           string
	BinWidth      string
	Mean          string
	Median        string
	LowerQuartile string
	UpperQuartile string
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s\t%s\t%s",
		s.Min,
		s.Max,
		s.BinWidth,
		s.Mean,
		s.Median,
		s.LowerQuartile,
		s.UpperQuartile,
	)
}

func Calculate(input io.Reader, rmNa bool) Stats {
	values := xos.ReadWords(input)
	dataType := xtype.FindDataType(values)

	switch dataType {

	case xtype.DataTypeUndecided:
		fallthrough
	case xtype.DataTypeInteger:
		ints := xtype.ToIntegers(values)
		values = nil
		return FindIntegerBinWidth(ints, rmNa)

	case xtype.DataTypeFloat:
		floats := xtype.ToFloats(values)
		values = nil
		return FindFloatBinWidth(floats, rmNa)

	case xtype.DataTypeDate:
		dates := xtype.ToDates(values)
		values = nil
		return FindDateBinWidth(dates, rmNa)

	case xtype.DataTypeBoolean:
		bools := xtype.ToBools(values)
		values = nil
		return FindBooleanBinWidth(bools, rmNa)

	default:
		log.Fatalln("unrecognized or mixed data types passed on stdin")
	}

	return Stats{}
}