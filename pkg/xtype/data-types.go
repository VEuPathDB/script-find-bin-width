package xtype

type DataType uint8

const (
	DataTypeUndecided DataType = 0
	DataTypeUnknown   DataType = 1
	DataTypeInteger   DataType = 2
	DataTypeFloat     DataType = 3
	DataTypeDate      DataType = 4
	DataTypeBoolean   DataType = 5
)

func (t DataType) String() string {
	switch t {
	case DataTypeUndecided:
		return "undecided"
	case DataTypeUnknown:
		return "unknown"
	case DataTypeInteger:
		return "integer"
	case DataTypeFloat:
		return "float"
	case DataTypeDate:
		return "date"
	case DataTypeBoolean:
		return "boolean"
	default:
		return "invalid data type"
	}
}

func (t DataType) IsValid() bool {
	return t < 6
}

func (t DataType) IsNumeric() bool {
	return t == DataTypeFloat || t == DataTypeInteger
}
