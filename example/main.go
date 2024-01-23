package main

import (
	"fmt"
	"github.com/joe-l-mathew/validogo" // Replace with the actual path to your validogo package
)

func main() {
	// Example for phone number validation
	phoneNumber := "1234567890"
	isValidPhoneNumber := validogo.ValidatePhoneNumber(phoneNumber)
	fmt.Printf("Is %s a valid phone number? %t\n", phoneNumber, isValidPhoneNumber)

	// Example for email validation
	email := "user@example.com"
	isValidEmail := validogo.ValidateEmailId(email)
	fmt.Printf("Is %s a valid email address? %t\n", email, isValidEmail)

	// Example for URL validation
	url := "https://www.example.com"
	isValidURL := validogo.ValidateURL(url)
	fmt.Printf("Is %s a valid URL? %t\n", url, isValidURL)

	// Example for IP address validation
	ipAddress := "192.168.1.1"
	isValidIPAddress := validogo.ValidateIPAddress(ipAddress)
	fmt.Printf("Is %s a valid IP address? %t\n", ipAddress, isValidIPAddress)

	// Example for password strength validation
	password := "SecureP@ssw0rd"
	passwordStrength := validogo.PasswordStrength{
		MinLength:        8,
		RequireUppercase: true,
		RequireLowercase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
	}
	isValidPassword := validogo.ValidatePasswordStrength(password, passwordStrength)
	fmt.Printf("Is %s a valid password? %t\n", password, isValidPassword)

	// Example for file type detection
	filename := "example.jpg"
	fileType := validogo.DetectFileType(filename)
	fmt.Printf("The file type of %s is: %s\n", filename, fileType)

	// Example for hex color code validation
	colorCode := "#1a2b3c"
	isValidColorCode := validogo.ValidateHexColorCode(colorCode)
	fmt.Printf("Is %s a valid hex color code? %t\n", colorCode, isValidColorCode)

	// Example for latitude and longitude validation
	latitude := "37.7749"
	longitude := "-122.4194"
	isValidLatitude := validogo.ValidateLatitude(latitude)
	isValidLongitude := validogo.ValidateLongitude(longitude)
	fmt.Printf("Is %s a valid latitude? %t\n", latitude, isValidLatitude)
	fmt.Printf("Is %s a valid longitude? %t\n", longitude, isValidLongitude)

	// Example for age validation within a range
	age := "25"
	isValidAge := validogo.ValidateAge(age, nil, nil) // No specific age range
	fmt.Printf("Is %s a valid age? %t\n", age, isValidAge)
}
