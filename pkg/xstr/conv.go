package xstr

import (
	"log"
	"strconv"
	"time"
)

func MustParseInt64(value string) int64 {
	out, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		log.Fatal("Failed to convert string " + value + " to int")
	}

	return out
}

func MustParseFloat64(value string) float64 {
	out, err := strconv.ParseFloat(value, 64)

	if err != nil {
		log.Fatal("failed to convert string " + value + " to float")
	}

	return out
}

const (
	isoOrder = "2006-01-02"
)

func MustParseDate(value string) time.Time {
	t, err := time.Parse(isoOrder, value)

	if err != nil {
		log.Fatal("failed to convert string " + value + " to date")
	}

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
}

func MustParseBool(value string) bool {
	switch value[0] {

	case 't', 'T', 'y', 'Y', '1':
		return true

	case 'f', 'F', 'n', 'N', '0':
		return false

	default:
		panic("unexpected non-boolean value: " + value)
	}
}
