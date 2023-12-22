package stats

import (
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

	IsText bool
}

func Calculate(input io.Reader, rmNa bool) Stats {
	values := xos.ReadWords(input)

	if len(values) < 4 {
		log.Fatalln("input must contain at least 4 elements")
	}

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
