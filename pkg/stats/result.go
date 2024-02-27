package stats

import (
	"find-bin-width/pkg/output"
	"find-bin-width/pkg/vpdb"
)

const (
	fieldNameAttributeStableID = "attribute_stable_id"
	fieldNameEntityTypeID      = "entity_type_id"
)

type Result struct {
	AttributeStableID vpdb.AttributeStableID
	EntityTypeID      vpdb.EntityTypeID
	Stats             Stats
}

func (r Result) FieldCount() int {
	return r.Stats.FieldCount() + 2
}

func (r Result) WriteFieldName(index int, buf *output.FieldNameBuffer) int {
	switch index {
	case 0:
		return copy(buf[:], fieldNameAttributeStableID)
	case 1:
		return copy(buf[:], fieldNameEntityTypeID)
	default:
		return r.Stats.WriteFieldName(index-2, buf)
	}
}

func (r Result) WriteFieldValue(index int, buf *output.FieldValueBuffer) int {
	switch index {
	case 0:
		return copy(buf[:], r.AttributeStableID.String())
	case 1:
		return copy(buf[:], r.EntityTypeID.String())
	default:
		return r.Stats.WriteFieldValue(index-2, buf)
	}
}

func (r Result) GetFieldType(index int) output.FieldType {
	switch index {
	case 0:
		return output.FieldTypeText
	case 1:
		return output.FieldTypeNumeric
	default:
		return r.Stats.GetFieldType(index - 2)
	}
}
