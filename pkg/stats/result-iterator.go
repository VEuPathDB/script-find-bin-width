package stats

import (
	"find-bin-width/pkg/input"
)

type ResultIterator struct {
	sourceIterator input.LineGroupIterator
}

func (r *ResultIterator) HasNext() bool {
	return r.sourceIterator.HasNext()
}

func (r *ResultIterator) Next() (Result, error) {
	group, err := r.sourceIterator.Next()

	if err != nil {
		return Result{}, err
	}

	if group.IsNA {
		return Result{
			AttributeStableID: group.AttributeStableID,
			EntityTypeID:      group.EntityTypeID,
			Stats:             naSummary(0),
		}, nil
	}

	return Result{
		AttributeStableID: group.AttributeStableID,
		EntityTypeID:      group.EntityTypeID,
		Stats:             calculateSummary(group.DataType, group.Values),
	}, nil
}
