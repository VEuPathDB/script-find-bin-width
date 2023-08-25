package xmath_test

import (
	"find-bin-width/pkg/xmath"
	"testing"
)

func TestHardRoundToInt(t *testing.T) {
	if xmath.HardRoundToInt(1.1) != 1 {
		t.Error("Bug in HardRoundToInt! (a)")
	}

	if xmath.HardRoundToInt(1.5) != 2 {
		t.Error("Bug in HardRoundToInt! (b)")
	}

	if xmath.HardRoundToInt(-0.1) != 0 {
		t.Error("Bug in HardRoundToInt! (c)")
	}

	if xmath.HardRoundToInt(-1.6) != -2 {
		t.Error("Bug in HardRoundToInt! (d)")
	}
}

func TestHardRound(t *testing.T) {
	if xmath.HardRound(1.1) != 1 {
		t.Error("Bug in HardRound! (a)")
	}

	if xmath.HardRound(1.5) != 2 {
		t.Error("Bug in HardRound! (b)")
	}

	if xmath.HardRound(-0.1) != 0 {
		t.Error("Bug in HardRound! (c)")
	}

	if xmath.HardRound(-1.6) != -2 {
		t.Error("Bug in HardRound! (d)")
	}
}
