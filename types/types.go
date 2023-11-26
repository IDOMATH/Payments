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

// Would love to have these constraints as const, but slices can't be made constants.

// AmericanExpressContraints houses the IIN ranges and possible lengths for
// cards issued by American Express
var AmericanExpressContraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{34, 37},
	Lengths:   []int{15},
}

var BankCardConstraints IssuerConstraints = IssuerConstraints{
	IinExacts: []int{5610},
	IinRanges: []IinRange{IinRange{min: 560221, max: 560225}},
	Lengths:   []int{16},
}

var ChinaTUnion IssuerConstraints = IssuerConstraints{
	IinExacts: []int{31},
}
