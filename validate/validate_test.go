package validate

import "testing"

func TestLuhnCheckDigit(t *testing.T) {
	amex := 371720518195098
	got := LuhnCheckDigit(amex)
	if got != true {
		t.Errorf("Expected true, got: %v", got)
	}
}
