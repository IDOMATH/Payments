package main

import (
	"fmt"
	"github.com/idomath/payments/types"
	"math"
	"strconv"
)

// These are all test cards generated from Paypal's developer website.
// DO NOT TRY TO USE THESE.
const (
	Amex = 371720518195098
)

func main() {
	fmt.Println("hello world")
	fmt.Println(ValidateCardNumber(Amex))
	fmt.Println(isAmericanExpress(Amex))
}

func ValidateCardNumber(cardNumber int) bool {
	checkDigit := cardNumber % 10

	return checkDigit == ComputeLuhnCheckDigit(cardNumber/10)
}

// ComputeLuhnCheckDigit takes an int and performs algorithm to find the last digit of the card being sent.
func ComputeLuhnCheckDigit(cardNumber int) int {
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

func CheckNumberLength(cardNumber, expectedLength int) bool {
	return expectedLength == len(strconv.Itoa(cardNumber))
}

func GetFirstNDigits(cardNumber, n int) int {
	cardLength := len(strconv.Itoa(cardNumber))

	return cardNumber / int(math.Pow10(cardLength-n))
}

func CheckIinExact(cardNumber, Iin int) bool {
	firstNDigits := GetFirstNDigits(cardNumber, len(strconv.Itoa(Iin)))
	return firstNDigits == Iin
}

func CheckIinRange(cardNumber int, iinRange types.IinRange) bool {
	firstNDigits := GetFirstNDigits(cardNumber, len(strconv.Itoa(iinRange.Max)))

	return firstNDigits >= iinRange.Min && firstNDigits <= iinRange.Max
}

// Cards issues by American Express have an IIN range of 34, 37 and are 15 digits long.
func isAmericanExpress(cardNumber int) bool {
	for _, length := range types.AmericanExpressContraints.Lengths {
		if CheckNumberLength(cardNumber, length) {
			break
		}
		return false
	}
	for _, exact := range types.AmericanExpressContraints.IinExacts {
		if CheckIinExact(cardNumber, exact) {
			return true
		}
	}

	return false
}

// Cards issued by Bankcard have an IIN 5610, 560221–560225 and are 16 digits long
func isBankcard(cardNumer int) bool {
	for _, length := range types.BankCardConstraints.Lengths {
		if CheckNumberLength(cardNumer, length) {
			break
		}
		return false
	}
	return false
}
