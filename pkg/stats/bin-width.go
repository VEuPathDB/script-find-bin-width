package stats

import "find-bin-width/pkg/xtype"

func findBinWidth(dataType xtype.DataType, data any) Stats {
	switch dataType {
	case xtype.DataTypeInteger:
		return findIntegerBinWidth(data.([]int64))
	case xtype.DataTypeFloat:
		return findFloatBinWidth(data.([]float64))
	case xtype.DataTypeDate:
		return findDateBinWidth(data.([]int64))
	case xtype.DataTypeBoolean:
		return findBooleanBinWidth(data.([]bool))
	}

	panic("illegal state, cannot find bin width of data type " + dataType.String())
}
