package stats

func findBooleanBinWidth(_ []bool) Stats {
	// According to the original R implementation, boolean inputs always result in
	// an NA output.
	return naStats(0)
}
