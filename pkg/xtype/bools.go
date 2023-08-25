package xtype

type PseudoBool = uint8

const (
	BoolFalse PseudoBool = 0
	BoolTrue  PseudoBool = 1
	BoolNA    PseudoBool = 3
)
