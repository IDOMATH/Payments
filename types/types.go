package types

type IinRange struct {
	Min int
	Max int
}

type IssuerConstraints struct {
	IinExacts []int
	IinRanges []IinRange
	Lengths   []int
}

type CardIssuer struct {
	Issuer      string
	Constraints IssuerConstraints
}

// Would love to have these constraints as const, but slices can't be made constants.

// AmericanExpressConstraints houses the IIN ranges and possible lengths for
// cards issued by American Express
var AmericanExpressConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{34, 37},
	Lengths:   []int{15},
}

var BankCardConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{5610},
	IinRanges: []IinRange{IinRange{Min: 560221, Max: 560225}},
	Lengths:   []int{16},
}

var ChinaTUnionConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{31},
	Lengths:   []int{19},
}

var MasterCardConstraints IssuerConstraints = IssuerConstraints{
	IinRanges: []IinRange{IinRange{Min: 2221, Max: 2720}, IinRange{Min: 51, Max: 55}},
	Lengths:   []int{16},
}

var DiscoverConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{6011, 65},
	IinRanges: []IinRange{IinRange{Min: 644, Max: 649}, IinRange{Min: 622126, Max: 622925}},
	Lengths:   []int{16, 17, 18, 19},
}

var VisaConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{4},
	Lengths:   []int{13, 16, 19},
}

var AmericanExpress CardIssuer = CardIssuer{Issuer: "American Express", Constraints: AmericanExpressConstraints}
var MasterCard CardIssuer = CardIssuer{Issuer: "Master Card", Constraints: MasterCardConstraints}
var Discover CardIssuer = CardIssuer{Issuer: "Discover", Constraints: DiscoverConstraints}
var Visa CardIssuer = CardIssuer{Issuer: "Visa", Constraints: VisaConstraints}
