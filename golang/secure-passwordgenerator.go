package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"
)

const (
	minLength = 8
	maxLength = 128
)

var (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

type Password struct {
	Value   string
	Created time.Time
	Strength string
}

func main() {
	fmt.Println("=== SecurePass Generator ===")
	
	for {
		fmt.Println("\n1. Generate Password")
		fmt.Println("2. Analyze Password Strength")
		fmt.Println("3. Exit")
		fmt.Print("Choose option (1-3): ")
		
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		
		switch choice {
		case 1:
			generateAndPrintPassword()
		case 2:
			analyzePassword()
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func generateAndPrintPassword() {
	length := getLength()
	includeSymbols := getIncludeSymbols()
	
	password := generateSecurePassword(length, includeSymbols)
	strength := calculateStrength(password)
	
	fmt.Printf("Generated Password: %s\n", password)
	fmt.Printf("Strength: %s\n", strength)
}

func getLength() int {
	for {
		fmt.Print("Enter password length (8-128): ")
		var length int
		_, err := fmt.Scanf("%d", &length)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		
		if length >= minLength && length <= maxLength {
			return length
		}
		fmt.Println("Length must be between 8 and 128")
	}
}

func getIncludeSymbols() bool {
	for {
		fmt.Print("Include symbols? (y/n): ")
		var response string
		_, err := fmt.Scanf("%s", &response)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		
		switch strings.ToLower(response) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Please enter 'y' or 'n'")
		}
	}
}

func generateSecurePassword(length int, includeSymbols bool) string {
	charset := uppercase + lowercase + digits
	if includeSymbols {
		charset += symbols
	}
	
	// Ensure at least one character from each required set
	password := make([]byte, length)
	
	// Add one character from each required set
	password[0] = getRandomChar(uppercase)
	password[1] = getRandomChar(lowercase)
	password[2] = getRandomChar(digits)
	
	if includeSymbols {
		password[3] = getRandomChar(symbols)
	}
	
	// Fill the rest with random characters
	for i := 4; i < length; i++ {
		password[i] = getRandomChar(charset)
	}
	
	// Shuffle the password to avoid predictable patterns
	shuffle(password)
	
	return string(password)
}

func getRandomChar(charset string) byte {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		// Fallback to less secure random if crypto rand fails
		n = big.NewInt(int64(time.Now().UnixNano() % int64(len(charset))))
	}
	
	return charset[n.Int64()]
}

func shuffle(slice []byte) {
	for i := len(slice) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			// Fallback to less secure random if crypto rand fails
			j = big.NewInt(int64(time.Now().UnixNano() % int64(i+1)))
		}
		
		slice[i], slice[j.Int64()] = slice[j.Int64()], slice[i]
	}
}

func analyzePassword() {
	fmt.Print("Enter password to analyze: ")
	var password string
	_, err := fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	
	strength := calculateStrength(password)
	fmt.Printf("Password strength: %s\n", strength)
	
	if strength == "Weak" {
		fmt.Println("Tips:")
		fmt.Println("- Use at least 12 characters")
		fmt.Println("- Include uppercase, lowercase, numbers, and symbols")
		fmt.Println("- Avoid common patterns")
	}
}

func calculateStrength(password string) string {
	if len(password) < 8 {
		return "Very Weak"
	}
	
	score := 0
	
	// Length check
	if len(password) >= 12 {
		score++
	}
	if len(password) >= 16 {
		score++
	}
	
	// Character variety checks
	if regexp.MustCompile(`[a-z]`).MatchString(password) {
		score++
	}
	if regexp.MustCompile(`[A-Z]`).MatchString(password) {
		score++
	}
	if regexp.MustCompile(`\d`).MatchString(password) {
		score++
	}
	if regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password) {
		score++
	}
	
	// Complexity check
	if hasRepeatingChars(password) {
		score--
	}
	if hasSequentialChars(password) {
		score--
	}
	
	switch {
	case score <= 1:
		return "Very Weak"
	case score <= 2:
		return "Weak"
	case score <= 3:
		return "Medium"
	case score <= 4:
		return "Strong"
	default:
		return "Very Strong"
	}
}

func hasRepeatingChars(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			return true
		}
	}
	return false
}

func hasSequentialChars(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
			return true
		}
		if password[i+1] == password[i]-1 && password[i+2] == password[i]-2 {
			return true
		}
	}
	return false
}
