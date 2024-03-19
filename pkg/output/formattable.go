package output

// FieldNameBuffer defines the type used as a buffer to hold field names to be
// written to an output stream.
type FieldNameBuffer = [64]byte

// FieldValueBuffer defines the type used as a buffer to hold field values to be
// written to an output stream.
type FieldValueBuffer = [512]byte

// FieldType used to indicate the output type for a specified field.
type FieldType uint8

const (
	// FieldTypeNA represents the field type of a field which contains no value.
	FieldTypeNA FieldType = iota

	// FieldTypeBoolean represents the field type of a field which contains a
	// boolean value.
	FieldTypeBoolean

	// FieldTypeNumeric represents the field type of a field which contains a
	// numeric value.
	FieldTypeNumeric

	// FieldTypeText represents the field type of a field which contains a string
	// value.
	FieldTypeText
)

// Formattable represents a value that may be written to an output stream via a
// Formatter instance.
type Formattable interface {
	// FieldCount returns the number of printable fields contained in this
	// Formattable object.
	FieldCount() int

	// WriteFieldName writes the name or key for the field at the given index to
	// the given buffer, returning the number of bytes written.
	//
	// Field name values should always be in snake case.
	WriteFieldName(index int, buf *FieldNameBuffer) int

	// WriteFieldValue writes the value for the field at the given index to the
	// given buffer, returning the number of bytes written.
	WriteFieldValue(index int, buf *FieldValueBuffer) int

	// GetFieldType returns the FieldType type indicator for the target field.
	GetFieldType(index int) FieldType
}
