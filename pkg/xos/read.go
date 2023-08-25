package xos

import (
	"os"
	"strings"

	"find-bin-width/pkg/xstr"
)

func ReadStdinWords() []string {
	readBuffer := [4096]byte{}
	holdBuffer := new(strings.Builder)
	values := make([]string, 0, 2048)
	lwcr := false

	for true {
		red := RequireReadBytes(os.Stdin, readBuffer[:])
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
