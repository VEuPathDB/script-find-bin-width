package xmath_test

import (
	"testing"

	"find-bin-width/pkg/xmath"
)

func TestSturges(t *testing.T) {
	intTests := []struct {
		size   int
		expect float64
	}{
		{31, 6},
		{512, 10},
		{1024, 11},
		{2048, 12},
		{4096, 13},
		{185158, 19},
	}

	for i := range intTests {
		input := make([]int64, intTests[i].size)
		actual := xmath.Sturges(input)

		if actual != intTests[i].expect {
			t.Error("Sturges test failed, expected", intTests[i].expect, "got", actual)
		}
	}
}
