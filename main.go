package validogo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

// Validate Password
type PasswordStrength struct {
	MinLength        int
	RequireUppercase bool
	RequireLowercase bool
	RequireNumbers   bool
	RequireSpecial   bool
}

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

// File Type
// FileType represents the type of the file.
type FileType string

const (
	Image    FileType = "image"
	Video    FileType = "video"
	PDF      FileType = "pdf"
	Document FileType = "document"
	Audio    FileType = "audio"
	Unknown  FileType = "unknown"
)

// DetectFileType determines the type of the file based on its extension.
func DetectFileType(filename string) FileType {
	// Convert the filename to lowercase for case-insensitive matching
	lowercaseFilename := strings.ToLower(filename)

	imageRegex := regexp.MustCompile(`\.(jpg|jpeg|png|gif|bmp|webp)$`)
	videoRegex := regexp.MustCompile(`\.(mp4|mov|avi|mkv|flv)$`)
	pdfRegex := regexp.MustCompile(`\.pdf$`)
	documentRegex := regexp.MustCompile(`\.(doc|docx|ppt|pptx|xls|xlsx)$`)
	audioRegex := regexp.MustCompile(`\.(mp3|wav|ogg|flac)$`)

	switch {
	case imageRegex.MatchString(lowercaseFilename):
		return Image
	case videoRegex.MatchString(lowercaseFilename):
		return Video
	case pdfRegex.MatchString(lowercaseFilename):
		return PDF
	case documentRegex.MatchString(lowercaseFilename):
		return Document
	case audioRegex.MatchString(lowercaseFilename):
		return Audio
	default:
		return Unknown
	}
}

// Color code
func ValidateHexColorCode(colorCode string) bool {
	hexColorRegex := regexp.MustCompile(`^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$`)
	return hexColorRegex.MatchString(colorCode)
}

// Latitude and Longitude
// ValidateLatitude checks if the given string is a valid latitude.
func ValidateLatitude(latitude string) bool {
	latitudeRegex := regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`)
	return latitudeRegex.MatchString(latitude)
}

// ValidateLongitude checks if the given string is a valid longitude.
func ValidateLongitude(longitude string) bool {
	longitudeRegex := regexp.MustCompile(`^[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
	return longitudeRegex.MatchString(longitude)
}

// ValidateAge checks if the given string is a valid age within the optional range.
func ValidateAge(age string, lowerLimit, upperLimit *int) bool {
	// Validate that age is a non-negative integer
	ageRegex := regexp.MustCompile(`^\d+$`)
	if !ageRegex.MatchString(age) {
		return false
	}

	// Convert age to an integer
	ageValue, err := strconv.Atoi(age)
	if err != nil {
		return false
	}

	// Check optional lower limit
	if lowerLimit != nil && ageValue < *lowerLimit {
		return false
	}

	// Check optional upper limit
	if upperLimit != nil && ageValue > *upperLimit {
		return false
	}

	return true
}
