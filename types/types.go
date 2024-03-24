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
	IinRanges: []IinRange{IinRange{Min: 560221, Max: 560225}},
	Lengths:   []int{16},
}

var ChinaTUnion IssuerConstraints = IssuerConstraints{
	IinExacts: []int{31},
	Lengths:   []int{19},
}

var MasterCard IssuerConstraints = IssuerConstraints{
	IinRanges: []IinRange{IinRange{Min: 2221, Max: 2720}, IinRange{Min: 51, Max: 55}},
	Lengths:   []int{16},
}

var Discover IssuerConstraints = IssuerConstraints{
	IinExacts: []int{6011, 65},
	IinRanges: []IinRange{IinRange{Min: 644, Max: 649}, IinRange{Min: 622126, Max: 622925}},
	Lengths:   []int{16, 17, 18, 19},
}
