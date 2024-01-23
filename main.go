package validogo

import (
	"fmt"
	"regexp"
	"strings"
)

// PasswordStrength represents the criteria for password strength.
type PasswordStrength struct {
	MinLength        int
	RequireUppercase bool
	RequireLowercase bool
	RequireNumbers   bool
	RequireSpecial   bool
}

// validate phone number
func ValidatePhoneNumber(phoneNumber string, numberOfDigits ...int) bool {
	digits := 10
	if len(numberOfDigits) > 0 {
		digits = numberOfDigits[0]
	}
	phoneRegex := regexp.MustCompile(`^\d{` + fmt.Sprint(digits) + `}$`)
	return phoneRegex.MatchString(phoneNumber)
}

// validate email
func ValidateEmailId(emailId string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(emailId)
}

// validate url
func ValidateURL(url string) bool {
	urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$`)
	return urlRegex.MatchString(url)
}

// validate ip address
func ValidateIPAddress(ipAddress string) bool {
	ipRegex := regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
	return ipRegex.MatchString(ipAddress)
}

// ValidatePasswordStrength checks if the given password meets the specified strength criteria.
func ValidatePasswordStrength(password string, strength PasswordStrength) bool {
	if len(password) < strength.MinLength {
		return false
	}

	if strength.RequireUppercase && !containsUppercase(password) {
		return false
	}

	if strength.RequireLowercase && !containsLowercase(password) {
		return false
	}

	if strength.RequireNumbers && !containsNumbers(password) {
		return false
	}

	if strength.RequireSpecial && !containsSpecialCharacters(password) {
		return false
	}

	return true
}

func containsUppercase(s string) bool {
	return strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func containsLowercase(s string) bool {
	return strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz")
}

func containsNumbers(s string) bool {
	return strings.ContainsAny(s, "0123456789")
}

func containsSpecialCharacters(s string) bool {
	specialCharactersRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return specialCharactersRegex.MatchString(s)
}
