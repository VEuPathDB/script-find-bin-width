package input

import (
	"fmt"

	"find-bin-width/pkg/vpdb"
	"find-bin-width/pkg/xstr"
)

func makeLineKey(key1, key2 string) (key lineKey, err error) {
	key.attributeSourceID, err = vpdb.ParseAttributeStableID(key1)
	if err != nil {
		return
	}

	key.entityTypeID, err = vpdb.ParseEntityTypeID(key2)

	return
}

type lineKey struct {
	attributeSourceID vpdb.AttributeStableID
	entityTypeID      vpdb.EntityTypeID
}

func (k lineKey) String() string {
	return "(attributeSourceId=" + k.attributeSourceID.String() +
		", entityTypeId=" + k.entityTypeID.String() + ")"
}

func TabDelimiterSplitFn() LineSplitFunction {
	return CharDelimiterSplitFn(xstr.AsciiHorizontalTab)
}

func CharDelimiterSplitFn(delim byte) LineSplitFunction {
	return func(line string) (LineSplitResult, error) {
		next := -1
		result := LineSplitResult{}
		for i := 0; i < LineElementCount; i++ {
			result[i], next = xstr.CharSplitNextSegment(delim, next+1, line)

			if next == -1 && i+1 < LineElementCount {
				return LineSplitResult{}, fmt.Errorf("input line contained %d field(s) when %d were expected", i+1, LineElementCount)
			}
		}

		if next != -1 {
			return LineSplitResult{}, fmt.Errorf("input line contained more than %d fields", LineElementCount)
		}

		return result, nil
	}
}
