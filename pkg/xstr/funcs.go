package xstr

func IsWhitespaceOrBreak(b byte) bool {
	return b == AsciiHorizontalTab || b == AsciiSpace || b == AsciiLineFeed || b == AsciiCarriageReturn
}
