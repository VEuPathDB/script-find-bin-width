package output

type FieldNameBuffer = [64]byte
type FieldValueBuffer = [256]byte

type FieldType uint8

const (
	FieldTypeNA FieldType = iota
	FieldTypeBoolean
	FieldTypeNumeric
	FieldTypeText
)

type Formattable interface {
	// FieldCount returns the number of printable fields contained in this
	// Formattable object.
	FieldCount() int

	// GetFieldName writes the name or key for the field at the given index to the
	// given buffer, returning the number of bytes written.
	//
	// Field name values should always be in snake case.
	GetFieldName(index int, buf *FieldNameBuffer) int

	// GetFieldValue writes the value for the field at the given index to the
	// given buffer, returning the number of bytes written.
	GetFieldValue(index int, buf *FieldValueBuffer) int

	GetFieldType(index int) FieldType
}
