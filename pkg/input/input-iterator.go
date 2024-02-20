package input

import (
	"bufio"
)

const LineElementCount = 3

type LineSplitResult [LineElementCount]string

func (l *LineSplitResult) split() (string, string, string) {
	return l[0], l[1], l[2]
}

type LineSplitFunction = func(line string) (LineSplitResult, error)

type inputIterator struct {
	line    uint
	scanner *bufio.Scanner
	err     error
	splitFn LineSplitFunction
}

func (i *inputIterator) hasNext() bool {
	if !i.scanner.Scan() {
		i.err = i.scanner.Err()
		if i.err != nil {
			i.err = inputError{i.line + 1, i.err}
		}
		return false
	}

	return true
}

func (i *inputIterator) next() (LineSplitResult, error) {
	if i.err != nil {
		return LineSplitResult{}, i.err
	}

	i.line++

	if out, err := i.splitFn(i.scanner.Text()); err != nil {
		return LineSplitResult{}, inputError{i.line, err}
	} else {
		return out, nil
	}
}
