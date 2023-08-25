package xos

import (
	"io"
	"strings"

	"find-bin-width/pkg/xstr"
)

// ReadWords reads the entirety of the given input stream and splits the
// contents into string words on whitespace.
//
// For example, given the input (note the extra space between 3 and 4):
//   1 2 3  4 5
// The output would be
//  ["1", "2", "3", "", "4", "5"]
func ReadWords(input io.Reader) []string {
	readBuffer := [4096]byte{}
	holdBuffer := new(strings.Builder)
	values := make([]string, 0, 2048)
	lwcr := false

	for true {
		red := RequireReadBytes(input, readBuffer[:])
		cur := byte(0)

		for i := 0; i < red; i++ {
			cur = readBuffer[i]

			if xstr.IsWhitespaceOrBreak(cur) {
				if cur == xstr.AsciiCarriageReturn {
					lwcr = true
				} else if lwcr && cur == xstr.AsciiLineFeed {
					lwcr = false
					continue
				} else {
					lwcr = false
				}

				values = append(values, holdBuffer.String())
				holdBuffer.Reset()
				continue
			} else {
				lwcr = false
			}

			holdBuffer.WriteByte(cur)
		}

		if red == 0 {
			break
		}
	}

	if holdBuffer.Len() > 0 {
		values = append(values, holdBuffer.String())
	}

	return values
}
