package validate

import (
	"github.com/idomath/payments/types"
	"math"
	"strconv"
)

func LuhnCheckDigit(cardNumber int) bool {
	checkDigit := cardNumber % 10

	return checkDigit == computeLuhnCheckDigit(cardNumber/10)
}

func CardLength(cardNumber, expectedLength int) bool {
	return expectedLength == len(strconv.Itoa(cardNumber))
}

func IinExact(cardNumber, Iin int) bool {
	firstNDigits := getFirstNDigits(cardNumber, len(strconv.Itoa(Iin)))
	return firstNDigits == Iin
}

func IinRange(cardNumber int, iinRange types.IinRange) bool {
	firstNDigits := getFirstNDigits(cardNumber, len(strconv.Itoa(iinRange.Max)))

	return firstNDigits >= iinRange.Min && firstNDigits <= iinRange.Max
}

func Issuer(cardNumber int, constraints types.IssuerConstraints) bool {
	isValidLength, isValidIinExact, isValidIinRange := false, false, false
	for _, length := range constraints.Lengths {
		if CardLength(cardNumber, length) {
			isValidLength = true
			break
		}
	}

	for _, exact := range constraints.IinExacts {
		if IinExact(cardNumber, exact) {
			isValidIinExact = true
			break
		}
	}

	for _, iinRange := range constraints.IinRanges {
		if IinRange(cardNumber, iinRange) {
			isValidIinRange = true
			break
		}
	}

	if len(constraints.Lengths) == 0 {
		isValidLength = true
	}
	if len(constraints.IinExacts) == 0 {
		isValidIinExact = true
	}
	if len(constraints.IinRanges) == 0 {
		isValidIinRange = true
	}

	return isValidLength && (isValidIinExact || isValidIinRange)
}

func getFirstNDigits(cardNumber, n int) int {
	cardLength := len(strconv.Itoa(cardNumber))

	return cardNumber / int(math.Pow10(cardLength-n))
}

// ComputeLuhnCheckDigit takes an int and performs algorithm to find the last digit of the card being sent.
func computeLuhnCheckDigit(cardNumber int) int {
	var (
		digits       []int
		sum          = 0
		shouldDouble = true
	)
	for num := cardNumber; num > 0; num = num / 10 {
		digit := num % 10
		digits = append(digits, digit)
	}

	for _, val := range digits {
		sumDigit := val
		if shouldDouble {
			sumDigit *= 2
		}
		shouldDouble = !shouldDouble

		ones := sumDigit % 10
		tens := sumDigit / 10
		sum += ones + tens
	}

	return 10 - sum%10
}
