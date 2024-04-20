package main

import (
	"fmt"
	"github.com/idomath/payments/types"
	"github.com/idomath/payments/validate"
)

// These are all test cards generated from Paypal's developer website.
// DO NOT TRY TO USE THESE.
const (
	Amex       = 371720518195098
	Visa       = 4032036811195625
	MasterCard = 5110925019961261
	Discover   = 6011397097510907
)

func main() {
	fmt.Println("hello world")
	fmt.Println(validate.LuhnCheckDigit(Amex))

	// Cards issues by American Express have an IIN range of 34, 37 and are 15 digits long.
	fmt.Println(validate.Issuer(Amex, types.AmericanExpressContraints))
	fmt.Println(validate.Issuer(MasterCard, types.MasterCardConstraints))
	fmt.Println(validate.Issuer(Discover, types.DiscoverConstraints))
	fmt.Println(validate.Issuer(Visa, types.VisaConstraints))
}
