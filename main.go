package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func generatePassword(length int, useUpper, useLower, useNumbers, useSpecial bool) (string, error) {
	var charset string
	if useUpper { charset += upperChars }
	if useLower { charset += lowerChars }
	if useNumbers { charset += numberChars }
	if useSpecial { charset += specialChars }

	if charset == "" {
		return "", fmt.Errorf("at least one character set must be selected")
	}

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil { return "", err }
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	length := flag.Int("l", 16, "Length of the password")
	upper := flag.Bool("u", true, "Include uppercase letters")
	lower := flag.Bool("lower", true, "Include lowercase letters")
	numbers := flag.Bool("n", true, "Include numbers")
	special := flag.Bool("s", true, "Include special characters")

	flag.Parse()

	password, err := generatePassword(*length, *upper, *lower, *numbers, *special)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(password)
}
