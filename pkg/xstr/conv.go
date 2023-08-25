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

func MustParseDate(value string) time.Time {
	var tmp string

	if len(value) == 10 {
		tmp = value + "T00:00:00.0Z"
	} else if len(value) == 19 {
		tmp = value + "Z"
	} else {
		tmp = value
	}

	out, err := time.Parse(time.RFC3339, tmp)

	if err != nil {
		log.Fatal("failed to convert string " + value + " to date")
	}

	return out
}
