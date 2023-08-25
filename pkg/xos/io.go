package xos

import (
	"io"
	"log"
	"os"
)

func RequireWriteString(w *os.File, v string) int {
	written, err := w.WriteString(v)

	if err != nil {
		log.Fatal("Failed to write string to io stream " + err.Error())
	}

	return written
}

func RequireReadBytes(r io.Reader, buffer []byte) int {
	red, err := r.Read(buffer)

	if err == io.EOF {
		return red
	}

	if err != nil {
		log.Fatal("Failed to read bytes from reader " + err.Error())
	}

	return red
}
