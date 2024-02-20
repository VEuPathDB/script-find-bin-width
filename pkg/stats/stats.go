package stats

import (
	"fmt"
	"strconv"
	"time"

	"find-bin-width/pkg/output"
	"find-bin-width/pkg/xtype"
)

type Stats interface {
	output.Formattable

	Min() string
	Max() string
	BinWidth() string
	Mean() string
	Median() string
	LowerQuartile() string
	UpperQuartile() string
}

const statFieldCount = 7

type statField uint8

const (
	statFieldMin statField = iota
	statFieldMax
	statFieldBinWidth
	statFieldMean
	statFieldMedian
	statFieldLowerQuartile
	statFieldUpperQuartile
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

func intStringifier(buf *output.FieldValueBuffer, i float64, f statField) int {
	switch f {
	case statFieldMin, statFieldMax, statFieldBinWidth:
		return copy(buf[:], strconv.FormatInt(int64(i), 10))
	default:
		return copy(buf[:], strconv.FormatFloat(i, 'f', -1, 64))
	}
}

func floatStringifier(buf *output.FieldValueBuffer, f float64, _ statField) int {
	return copy(buf[:], strconv.FormatFloat(f, 'f', -1, 64))
}

const dateFormat = "2006-01-02"

type dateBinWidth = int64

const (
	dateBinWidthDay dateBinWidth = iota
	dateBinWidthWeek
	dateBinWidthMonth
	dateBinWidthYear
)

const sDay = "day"
const sWeek = "week"
const sMonth = "month"
const sYear = "year"

func dateStringifier(buf *output.FieldValueBuffer, i int64, f statField) int {
	if f == statFieldBinWidth {
		switch i {
		case dateBinWidthDay:
			return copy(buf[:], sDay)
		case dateBinWidthWeek:
			return copy(buf[:], sWeek)
		case dateBinWidthMonth:
			return copy(buf[:], sMonth)
		case dateBinWidthYear:
			return copy(buf[:], sYear)
		default:
			panic("invalid date bin width type")
		}
	}

	return copy(buf[:], time.Unix(i, 0).Format(dateFormat))
}

type stats[T any] struct {
	min           T
	max           T
	binWidth      T
	mean          T
	median        T
	lowerQuartile T
	upperQuartile T
	stringifier   func(buf *output.FieldValueBuffer, value T, field statField) int
	dataType      xtype.DataType
}

func (s stats[T]) GetFieldType(int) output.FieldType {
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

func (s stats[T]) FieldCount() int {
	return statFieldCount
}

func (s stats[T]) GetFieldName(index int, buf *output.FieldNameBuffer) int {
	if index < statFieldCount {
		return copy(buf[:], fieldNames[index])
	}

	panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, s.FieldCount()))
}

func (s stats[T]) GetFieldValue(index int, buf *output.FieldValueBuffer) int {
	switch index {
	case 0:
		return s.stringifier(buf, s.min, statFieldMin)
	case 1:
		return s.stringifier(buf, s.max, statFieldMax)
	case 2:
		return s.stringifier(buf, s.binWidth, statFieldBinWidth)
	case 3:
		return s.stringifier(buf, s.mean, statFieldMean)
	case 4:
		return s.stringifier(buf, s.median, statFieldMedian)
	case 5:
		return s.stringifier(buf, s.lowerQuartile, statFieldLowerQuartile)
	case 6:
		return s.stringifier(buf, s.upperQuartile, statFieldUpperQuartile)
	default:
		panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, s.FieldCount()))
	}
}

func (s stats[T]) Min() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.min, statFieldMin)
	return string(buf[:num])
}

func (s stats[T]) Max() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.max, statFieldMax)
	return string(buf[:num])
}

func (s stats[T]) BinWidth() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.binWidth, statFieldBinWidth)
	return string(buf[:num])
}

func (s stats[T]) Mean() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.mean, statFieldMean)
	return string(buf[:num])
}

func (s stats[T]) Median() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.median, statFieldMedian)
	return string(buf[:num])
}

func (s stats[T]) LowerQuartile() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.lowerQuartile, statFieldLowerQuartile)
	return string(buf[:num])
}

func (s stats[T]) UpperQuartile() string {
	buf := output.FieldValueBuffer{}
	num := s.stringifier(&buf, s.upperQuartile, statFieldUpperQuartile)
	return string(buf[:num])
}

type naStats uint8

func (naStats) FieldCount() int {
	return statFieldCount
}

func (naStats) GetFieldName(index int, buf *output.FieldNameBuffer) int {
	if index < statFieldCount {
		return copy(buf[:], fieldNames[index])
	}

	panic(fmt.Sprintf("index %d is out of bounds for field count %d", index, statFieldCount))
}

func (n naStats) GetFieldValue(int, *output.FieldValueBuffer) int {
	return 0
}

func (n naStats) GetFieldType(int) output.FieldType {
	return output.FieldTypeNA
}

func (n naStats) Min() string {
	return ""
}

func (n naStats) Max() string {
	return ""
}

func (n naStats) BinWidth() string {
	return ""
}

func (n naStats) Mean() string {
	return ""
}

func (n naStats) Median() string {
	return ""
}

func (n naStats) LowerQuartile() string {
	return ""
}

func (n naStats) UpperQuartile() string {
	return ""
}
