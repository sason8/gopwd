package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/atotto/clipboard"
)

const (
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	similarChars = "l1Io0O|"
)

func generatePassword(length int, useUpper, useLower, useNumbers, useSpecial, excludeSimilar bool) (string, error) {
	var charset string
	if useUpper {
		charset += upperChars
	}
	if useLower {
		charset += lowerChars
	}
	if useNumbers {
		charset += numberChars
	}
	if useSpecial {
		charset += specialChars
	}

	if excludeSimilar {
		for _, char := range similarChars {
			charset = strings.ReplaceAll(charset, string(char), "")
		}
	}

	if charset == "" {
		return "", fmt.Errorf("at least one character set must be selected and not fully excluded")
	}

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func calculateEntropy(length int, useUpper, useLower, useNumbers, useSpecial, excludeSimilar bool) float64 {
	poolSize := 0
	if useUpper {
		poolSize += len(upperChars)
	}
	if useLower {
		poolSize += len(lowerChars)
	}
	if useNumbers {
		poolSize += len(numberChars)
	}
	if useSpecial {
		poolSize += len(specialChars)
	}

	if excludeSimilar {
		// Calculate how many of the similar characters were in the selected sets
		excludedCount := 0
		for _, char := range similarChars {
			inPool := false
			if useUpper && strings.ContainsRune(upperChars, char) {
				inPool = true
			}
			if useLower && strings.ContainsRune(lowerChars, char) {
				inPool = true
			}
			if useNumbers && strings.ContainsRune(numberChars, char) {
				inPool = true
			}
			if useSpecial && strings.ContainsRune(specialChars, char) {
				inPool = true
			}
			if inPool {
				excludedCount++
			}
		}
		poolSize -= excludedCount
	}

	if poolSize <= 0 {
		return 0.0
	}

	// Entropy = L * log2(R)
	return float64(length) * math.Log2(float64(poolSize))
}

func getStrengthLabel(entropy float64) string {
	if entropy < 40 {
		return "Weak (Słabe) 🔴"
	} else if entropy < 80 {
		return "Medium (Średnie) 🟡"
	}
	return "Strong (Silne) 🟢"
}

func main() {
	length := flag.Int("l", 16, "Length of the password")
	upper := flag.Bool("u", true, "Include uppercase letters (A-Z)")
	lower := flag.Bool("lower", true, "Include lowercase letters (a-z)")
	numbers := flag.Bool("n", true, "Include numbers (0-9)")
	special := flag.Bool("s", true, "Include special characters")
	excludeSimilar := flag.Bool("e", false, "Exclude similar characters (l, 1, I, o, 0, O, |)")
	copyToClip := flag.Bool("c", false, "Copy generated password to clipboard (applies to first password only)")
	info := flag.Bool("info", false, "Show strength and entropy metrics")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	if *length <= 0 {
		fmt.Println("Error: Password length must be greater than 0")
		return
	}
	if *count <= 0 {
		fmt.Println("Error: Count must be greater than 0")
		return
	}

	var passwords []string
	for i := 0; i < *count; i++ {
		password, err := generatePassword(*length, *upper, *lower, *numbers, *special, *excludeSimilar)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		passwords = append(passwords, password)
		fmt.Println(password)
	}

	if *info {
		entropy := calculateEntropy(*length, *upper, *lower, *numbers, *special, *excludeSimilar)
		fmt.Printf("\n--- Password Strength Analysis ---\n")
		fmt.Printf("Length: %d\n", *length)
		fmt.Printf("Entropy: %.2f bits\n", entropy)
		fmt.Printf("Strength: %s\n", getStrengthLabel(entropy))
	}

	if *copyToClip && len(passwords) > 0 {
		err := clipboard.WriteAll(passwords[0])
		if err != nil {
			fmt.Printf("\nError copying to clipboard: %v\n", err)
		} else {
			fmt.Println("\n📋 Password copied to clipboard successfully!")
		}
	}
}
