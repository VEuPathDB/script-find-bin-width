package vpdb

import "strconv"

type AttributeStableID string

func ParseAttributeStableID(raw string) (AttributeStableID, error) {
	return AttributeStableID(raw), nil
}

func (a AttributeStableID) Compare(other AttributeStableID) int {
	switch true {
	case a > other:
		return 1
	case a < other:
		return -1
	default:
		return 0
	}
}

func (a AttributeStableID) String() string {
	return string(a)
}

type EntityTypeID uint32

func ParseEntityTypeID(raw string) (EntityTypeID, error) {
	if r, err := strconv.ParseUint(raw, 10, 32); err != nil {
		return EntityTypeID(0), err
	} else {
		return EntityTypeID(uint32(r)), nil
	}
}

func (e EntityTypeID) Compare(other EntityTypeID) int {
	switch true {
	case e > other:
		return 1
	case e < other:
		return -1
	default:
		return 0
	}
}

func (e EntityTypeID) String() string {
	return strconv.FormatUint(uint64(e), 10)
}
