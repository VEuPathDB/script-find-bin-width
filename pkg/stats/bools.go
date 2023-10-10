package stats

import "find-bin-width/pkg/xtype"

func FindBooleanBinWidth(_ []xtype.PseudoBool, _ bool) Stats {
	// According to the original R implementation, boolean inputs always result in
	// an NA output.
	return Stats{}
}
