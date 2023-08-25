package bw

import "find-bin-width/pkg/xtype"

func FindBooleanBinWidth(_ []xtype.PseudoBool, _ bool) string {
	// According to the original R implementation, boolean inputs always result in
	// an NA output.
	return ""
}
