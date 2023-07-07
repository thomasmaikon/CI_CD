package tests

import (
	"CI/math"
	"testing"
)

func TestSum(t *testing.T) {
	expected := 14

	received := math.Sum(8, 6)

	if received != expected {
		t.Errorf("Value expected [%d] has differente that received [%d]", expected, received)
	}
}

func TestSubstract(t *testing.T) {
	expected := 0

	received := math.Subtract(10, 10)

	if received != expected {
		t.Errorf("Value expected [%d] has differente that received [%d]", expected, received)
	}
}

func TestDivision(t *testing.T) {
	expected := 2

	received := math.Subtract(16, 8)

	if received != expected {
		t.Errorf("Value expected [%d] has differente that received [%d]", expected, received)
	}
}
