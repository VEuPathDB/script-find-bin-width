package xtype

// PseudoBool represents a boolean value that may also be NA.
type PseudoBool = uint8

const (
	BoolFalse PseudoBool = 0
	BoolTrue  PseudoBool = 1
	BoolNA    PseudoBool = 3
)
