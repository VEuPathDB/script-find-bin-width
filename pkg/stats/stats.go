package stats

import (
	"find-bin-width/pkg/output"
)

// Stats is a container for the data set stat calculation results.
type Stats interface {
	output.Formattable

	// Min returns the minimum value from the input data.
	Min() string

	// Max returns the maximum value from the input data.
	Max() string

	// BinWidth returns the histogram bin width calculated for the input data set.
	BinWidth() string

	// Mean returns the mean value for the input data.
	Mean() string

	// Median returns the median value for the input data.
	Median() string

	// LowerQuartile returns the lower quartile value calculated for the input
	// data.
	LowerQuartile() string

	// UpperQuartile returns the upper quartile value calculated for the input
	// data.
	UpperQuartile() string
}
