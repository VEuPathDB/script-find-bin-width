package xos

import (
	"bufio"
)

func BufWriteString(stream *bufio.Writer, text string) {
	if _, err := stream.WriteString(text); err != nil {
		panic(err)
	}
}

func BufWriteBytes(stream *bufio.Writer, value []byte) {
	if _, err := stream.Write(value); err != nil {
		panic(err)
	}
}

func BufWriteByte(stream *bufio.Writer, value byte) {
	if err := stream.WriteByte(value); err != nil {
		panic(err)
	}
}
