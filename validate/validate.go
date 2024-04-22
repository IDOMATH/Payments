package validate

import (
	"errors"
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
	isValidLength, isValidIinExact, matchedIinExact, isValidIinRange := false, false, false, false

	if len(constraints.Lengths) != 0 {
		for _, length := range constraints.Lengths {
			if CardLength(cardNumber, length) {
				isValidLength = true
				break
			}
		}
	} else {
		isValidLength = true
	}

	if len(constraints.IinExacts) != 0 {
		for _, exact := range constraints.IinExacts {
			if IinExact(cardNumber, exact) {
				matchedIinExact = true
				isValidIinExact = true
				break
			}
		}
	} else {
		isValidIinExact = true
	}

	if !matchedIinExact && len(constraints.IinRanges) != 0 {
		for _, iinRange := range constraints.IinRanges {
			if IinRange(cardNumber, iinRange) {
				isValidIinRange = true
				break
			}
		}
	} else {
		isValidIinRange = true
	}

	return isValidLength && isValidIinExact && isValidIinRange
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

func GetCardIssuer(cardNumber int) (string, error) {
	var cardIssuers = make([]types.CardIssuer, 4)
	cardIssuers[0] = types.AmericanExpress
	cardIssuers[1] = types.MasterCard
	cardIssuers[2] = types.Discover
	cardIssuers[3] = types.Visa

	if !LuhnCheckDigit(cardNumber) {
		return "", errors.New("Luhn Check digit failed")
	}

	for _, issuer := range cardIssuers {
		if Issuer(cardNumber, issuer.Constraints) {
			return issuer.Issuer, nil
		}
	}

	return "", errors.New("Card Issuer not found")
}
