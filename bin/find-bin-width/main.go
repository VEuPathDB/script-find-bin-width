package main

import (
	"log"
	"os"

	"find-bin-width/pkg/bw"
	"find-bin-width/pkg/xos"
	"find-bin-width/pkg/xtype"
)

func main() {
	values := xos.ReadStdinWords()
	dataType := xtype.FindDataType(values)

	var out string

	switch dataType {
	case xtype.DataTypeUndecided:
		fallthrough
	case xtype.DataTypeInteger:
		ints := xtype.ToIntegers(values)
		values = nil
		out = bw.FindIntegerBinWidth(ints, true)
		break
	case xtype.DataTypeFloat:
		floats := xtype.ToFloats(values)
		values = nil
		out = bw.FindFloatBinWidth(floats, true)
		break
	case xtype.DataTypeDate:
		dates := xtype.ToDates(values)
		values = nil
		out = bw.FindDateBinWidth(dates, true)
		break
	case xtype.DataTypeBoolean:
		log.Fatalln("boolean data types not yet implemented")
	default:
		log.Fatalln("unrecognized or mixed data types passed on stdin")
	}

	xos.RequireWriteString(os.Stdout, out)
}
