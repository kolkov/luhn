package luhn

import (
	"errors"
)

var ErrValueNotLuhn = errors.New("luhn: must be valid number by Luhn algorithm")

// Validate check number is valid or not based on Luhn algorithm
func Validate(value interface{}) error {
	s, _ := value.(string)
	sum, err := checkSum(s)
	if err != nil {
		return err
	}
	if sum%10 != 0 {
		return ErrValueNotLuhn
	}
	return nil
}

// Generate return the number with check digit
func Generate(source string) (string, error) {
	sum, err := checkSum(source + "0")
	checkDigit := 10 - sum%10
	return source + string(checkDigit+48), err
}

func checkSum(code string) (int, error) {
	sum := 0
	strArr := []rune(code)
	nDigits := len(code)
	parity := nDigits % 2
	for i := 0; i < nDigits; i++ {
		digit, err := charToNum(strArr[i])
		if err != nil {
			return -1, err
		}
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum, nil
}

var ErrRuneNotInt = errors.New("luhn: rune was not int")

func charToNum(r rune) (int, error) {
	if '0' <= r && r <= '9' {
		return int(r) - '0', nil
	}
	return 0, ErrRuneNotInt
}
