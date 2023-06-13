package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func IsStrongPassword(password string) (bool, string) {
	// Check the length
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long."
	}

	// Check for at least one uppercase letter
	hasUppercase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
			break
		}
	}
	if !hasUppercase {
		return false, "Password must contain at least one uppercase letter."
	}

	// Check for at least one lowercase letter
	hasLowercase := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowercase = true
			break
		}
	}
	if !hasLowercase {
		return false, "Password must contain at least one lowercase letter."
	}

	// Check for at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return false, "Password must contain at least one digit."
	}

	// Check for at least one special character
	hasSpecialChar, _ := regexp.MatchString(`[^a-zA-Z0-9]`, password)
	if !hasSpecialChar {
		return false, "Password must contain at least one special character."
	}

	// Check for consecutive characters
	lowerPwd := strings.ToLower(password)
	for i := 0; i < len(password)-2; i++ {
		if lowerPwd[i]+1 == lowerPwd[i+1] && lowerPwd[i+1]+1 == lowerPwd[i+2] {
			return false, "Password cannot have consecutive characters in alphabetical order."
		}
	}

	return true, "Password is strong."
}
