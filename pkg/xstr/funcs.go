package xstr

func IsWhitespaceOrBreak(b byte) bool {
	return b == AsciiHorizontalTab || b == AsciiSpace || b == AsciiLineFeed || b == AsciiCarriageReturn
}

func IndexOfByte(char byte, startFrom int, target string) int {
	l := len(target)

	for i := startFrom; i < l; i++ {
		if target[i] == char {
			return i
		}
	}

	return -1
}

func CharSplitNextSegment(char byte, startFrom int, target string) (segment string, charPos int) {
	i := IndexOfByte(char, startFrom, target)

	if i == -1 {
		return target[startFrom:], i
	}

	return target[startFrom:i], i
}
