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

	fmt.Println(validate.GetCardIssuer(Amex))
	fmt.Println(validate.GetCardIssuer(Visa))
	fmt.Println(validate.GetCardIssuer(MasterCard))
	fmt.Println(validate.GetCardIssuer(Discover))

	fmt.Println(validate.LuhnCheckDigit(Amex))

	fmt.Println(validate.Issuer(Amex, types.AmericanExpressConstraints))
	fmt.Println(validate.Issuer(MasterCard, types.MasterCardConstraints))
	fmt.Println(validate.Issuer(Discover, types.DiscoverConstraints))
	fmt.Println(validate.Issuer(Visa, types.VisaConstraints))
}
