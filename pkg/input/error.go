package input

import "fmt"

type inputError struct {
	line uint
	err  error
}

func (i inputError) Unwrap() error {
	return i.err
}

func (i inputError) Error() string {
	return fmt.Sprintf("error encountered on input line %d: %s", i.line, i.err.Error())
}
