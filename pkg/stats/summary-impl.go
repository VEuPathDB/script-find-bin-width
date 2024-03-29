package stats

import (
	"fmt"

	"find-bin-width/pkg/output"
	"find-bin-width/pkg/xtype"
)

const summaryFieldCount = 7

type summaryField uint8

const (
	summaryFieldMin summaryField = iota
	summaryFieldMax
	summaryFieldBinWidth
	summaryFieldMean
	summaryFieldMedian
	summaryFieldLowerQuartile
	summaryFieldUpperQuartile
)

var fieldNames = [...]string{
	"min",
	"max",
	"bin_width",
	"mean",
	"median",
	"lower_quartile",
	"upper_quartile",
}

// summary implements the Summary interface as a generic container for
// calculated data set stat values.
type summary[T float64 | int64 | bool] struct {
	min           T
	max           T
	binWidth      T
	mean          T
	median        T
	lowerQuartile NullablePrimitive[T]
	upperQuartile NullablePrimitive[T]
	stringifier   valueStringifier
	dataType      xtype.DataType
}

func (s summary[T]) GetFieldType(int) output.FieldType {
	switch s.dataType {
	case xtype.DataTypeFloat, xtype.DataTypeInteger:
		return output.FieldTypeNumeric
	case xtype.DataTypeBoolean:
		return output.FieldTypeBoolean
	case xtype.DataTypeDate:
		return output.FieldTypeText
	default:
		panic("illegal data type: " + s.dataType.String())
	}
}

func (s summary[T]) FieldCount() int {
	return summaryFieldCount
}

func (s summary[T]) WriteFieldName(index int, buf *output.FieldNameBuffer) int {
	if index < summaryFieldCount {
		return copy(buf[:], fieldNames[index])
	}

	panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, s.FieldCount()))
}

func (s summary[T]) WriteFieldValue(index int, buf *output.FieldValueBuffer) int {
	switch index {
	case 0:
		return s.stringifier(buf, s.min, summaryFieldMin)
	case 1:
		return s.stringifier(buf, s.max, summaryFieldMax)
	case 2:
		return s.stringifier(buf, s.binWidth, summaryFieldBinWidth)
	case 3:
		return s.stringifier(buf, s.mean, summaryFieldMean)
	case 4:
		return s.stringifier(buf, s.median, summaryFieldMedian)
	case 5:
		return s.stringifier(buf, s.lowerQuartile, summaryFieldLowerQuartile)
	case 6:
		return s.stringifier(buf, s.upperQuartile, summaryFieldUpperQuartile)
	default:
		panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, s.FieldCount()))
	}
}

func (s summary[T]) Min() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.min, summaryFieldMin)
	return string(buf[:num])
}

func (s summary[T]) Max() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.max, summaryFieldMax)
	return string(buf[:num])
}

func (s summary[T]) BinWidth() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.binWidth, summaryFieldBinWidth)
	return string(buf[:num])
}

func (s summary[T]) Mean() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.mean, summaryFieldMean)
	return string(buf[:num])
}

func (s summary[T]) Median() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.median, summaryFieldMedian)
	return string(buf[:num])
}

func (s summary[T]) LowerQuartile() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.lowerQuartile, summaryFieldLowerQuartile)
	return string(buf[:num])
}

func (s summary[T]) UpperQuartile() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.upperQuartile, summaryFieldUpperQuartile)
	return string(buf[:num])
}

// naSummary implements the Summary interface for an NA row.
type naSummary uint8

func (naSummary) FieldCount() int {
	return summaryFieldCount
}

func (naSummary) WriteFieldName(index int, buf *output.FieldNameBuffer) int {
	if index < summaryFieldCount {
		return copy(buf[:], fieldNames[index])
	}

	panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, summaryFieldCount))
}

func (n naSummary) WriteFieldValue(int, *output.FieldValueBuffer) int {
	return 0
}

func (n naSummary) GetFieldType(int) output.FieldType {
	return output.FieldTypeNA
}

func (n naSummary) Min() string {
	return ""
}

func (n naSummary) Max() string {
	return ""
}

func (n naSummary) BinWidth() string {
	return ""
}

func (n naSummary) Mean() string {
	return ""
}

func (n naSummary) Median() string {
	return ""
}

func (n naSummary) LowerQuartile() string {
	return ""
}

func (n naSummary) UpperQuartile() string {
	return ""
}
