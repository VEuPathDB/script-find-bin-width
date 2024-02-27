package stats

import (
	"io"

	"find-bin-width/pkg/input"
)

func CalculateUnsorted(stream io.Reader, rmNa bool) (ResultIterator, error) {
	if it, err := input.CollectAndSortGroups(stream, rmNa, input.TabDelimiterSplitFn()); err != nil {
		return ResultIterator{}, err
	} else {
		return ResultIterator{it}, nil
	}
}

func CalculateSorted(stream io.Reader, rmNa bool) (ResultIterator, error) {
	if it, err := input.StreamGroups(stream, rmNa, input.TabDelimiterSplitFn()); err != nil {
		return ResultIterator{}, err
	} else {
		return ResultIterator{it}, nil
	}
}
