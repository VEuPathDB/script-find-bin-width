package stats

import "find-bin-width/pkg/xtype"

func calculateSummary(dataType xtype.DataType, data any) Summary {
	switch dataType {
	case xtype.DataTypeInteger:
		return calculateIntegerSummary(data.([]int64))
	case xtype.DataTypeFloat:
		return calculateFloatSummary(data.([]float64))
	case xtype.DataTypeDate:
		return calculateDateSummary(data.([]int64))
	case xtype.DataTypeBoolean:
		return calculateBooleanSummary(data.([]bool))
	}

	panic("illegal state, cannot find bin width of data type " + dataType.String())
}
