package stats

func calculateBooleanSummary(_ []bool) Summary {
	// According to the original R implementation, boolean inputs always result in
	// an NA output.
	return naSummary(0)
}
