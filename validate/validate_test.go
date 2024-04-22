package validate

import (
	"github.com/idomath/payments/types"
	"strconv"
	"testing"
)

const (
	amex       = 371720518195098
	masterCard = 5110925019961261
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
	got := IinExact(amex, types.AmericanExpressConstraints.IinExacts[1])
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

func TestIinRange(t *testing.T) {
	expected := true
	got := IinRange(masterCard, types.MasterCardConstraints.IinRanges[1])
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}
