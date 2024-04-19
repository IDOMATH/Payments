package validate

import (
	"strconv"
	"testing"
)

func TestLuhnCheckDigit(t *testing.T) {
	amex := 371720518195098
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
