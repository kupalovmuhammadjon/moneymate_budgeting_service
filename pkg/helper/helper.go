package helper

import (
	"fmt"
	"math/rand"
	"net/mail"
	"strconv"
	"unicode"
)

func CheckEmailAndPasswordValid(email, password string) error {

	// check email is valid
	if err := CheckEmailIsValid(email); err != nil {
		return fmt.Errorf("invalid email %s", err)
	}

	err := CheckPasswordIsStrong(password)
	if err != nil {
		return err
	}

	return nil
}

func CheckEmailIsValid(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email %s", err)
	}

	return nil
}

func CheckPasswordIsStrong(password string) error {

	var (
		special_chars  = map[rune]bool{}
		lowerCaseCount = 0
		upperCaseCount = 0
		digitCount     = 0
		signCount      = 0
	)
	special_chars = map[rune]bool{
		'!': true, '~': true,
		'@': true, '#': true,
		'$': true, '%': true,
		'^': true, '&': true,
		'*': true, '`': true,
		'(': true, ')': true,
		'-': true, '_': true,
		'+': true, '=': true,
		'{': true, '}': true,
		'[': true, ']': true,
		'|': true, '\\': true,
		';': true, ':': true,
		'"': true, '<': true,
		'?': true, '>': true,
		',': true, '.': true,
		'/': true,
	}

	if len(password) < 8 {
		return fmt.Errorf("password ning uzunligi 8 ta belgidan kam bo'lmasligi kerak")
	}

	for _, val := range password {
		if unicode.IsUpper(val) {
			upperCaseCount++
		} else if unicode.IsLower(val) {
			lowerCaseCount++
		} else if unicode.IsDigit(val) {
			digitCount++
		} else if _, ok := special_chars[val]; ok {
			signCount++
		}
	}

	if upperCaseCount == 0 {
		return fmt.Errorf("passwordda kamida 1 ta katta harf bo'lishi kerak")
	}
	if lowerCaseCount == 0 {
		return fmt.Errorf("passwordda kamida 1 ta kichik harf bo'lishi kerak")
	}
	if digitCount == 0 {
		return fmt.Errorf("passwordda kamida 1 ta raqam bo'lishi kerak")
	}
	if signCount == 0 {
		return fmt.Errorf("passwordda kamida 1 ta special character bo'lishi kerak")
	}
	return nil
}

func RandomCodeMaker() string {

	code := ""

	for i := 0; len(code) < 8; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}

	return code
}
