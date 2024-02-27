package stats

import "find-bin-width/pkg/xtype"

func calculateStats(dataType xtype.DataType, data any) Stats {
	switch dataType {
	case xtype.DataTypeInteger:
		return calculateIntegerStats(data.([]int64))
	case xtype.DataTypeFloat:
		return calculateFloatStats(data.([]float64))
	case xtype.DataTypeDate:
		return calculateDateStats(data.([]int64))
	case xtype.DataTypeBoolean:
		return calculateBooleanStats(data.([]bool))
	}

	panic("illegal state, cannot find bin width of data type " + dataType.String())
}
