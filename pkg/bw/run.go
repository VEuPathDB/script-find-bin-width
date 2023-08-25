package bw

import (
	"io"
	"log"

	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xtype"
)

func FindBinWidth(input io.Reader, rmNa bool) string {
	values := xos.ReadWords(input)
	dataType := xtype.FindDataType(values)

	var out string

	switch dataType {

	case xtype.DataTypeUndecided:
		fallthrough
	case xtype.DataTypeInteger:
		ints := xtype.ToIntegers(values)
		values = nil
		out = FindIntegerBinWidth(ints, rmNa)
		break

	case xtype.DataTypeFloat:
		floats := xtype.ToFloats(values)
		values = nil
		out = FindFloatBinWidth(floats, rmNa)
		break

	case xtype.DataTypeDate:
		dates := xtype.ToDates(values)
		values = nil
		out = FindDateBinWidth(dates, rmNa)
		break

	case xtype.DataTypeBoolean:
		bools := xtype.ToBools(values)
		values = nil
		out = FindBooleanBinWidth(bools, rmNa)
		break

	default:
		log.Fatalln("unrecognized or mixed data types passed on stdin")
	}

	return out
}
