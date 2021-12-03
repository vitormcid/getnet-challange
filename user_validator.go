package main

import "unicode"

func validatePassword(s string) bool {
	var (
		hasMinLen     = false
		hasUpper      = false
		hasLower      = false
		hasNumber     = false
		hasSpecial    = false
		lastCharacter rune
	)

	if len(s) >= 9 {
		hasMinLen = true
	}

	for _, char := range s {

		if char == lastCharacter {
			return false
		}

		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}

		lastCharacter = char
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
