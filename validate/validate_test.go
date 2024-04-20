package validate

import (
	"strconv"
	"testing"
)

const (
	amex    = 371720518195098
	amexIin = 37
)

func TestLuhnCheckDigit(t *testing.T) {
	expected := true
	got := LuhnCheckDigit(amex)
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}
func TestCardLength(t *testing.T) {
	card := 1234567890
	expected := true
	got := CardLength(card, len(strconv.Itoa(card)))
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

func TestIinExact(t *testing.T) {
	expected := true
	got := IinExact(amex, amexIin)
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}
