package main

import (
	"strings"
	"testing"
)

func TestGeneratePasswordLength(t *testing.T) {
	lengths := []int{8, 16, 32, 64}
	for _, l := range lengths {
		pw, err := generatePassword(l, true, true, true, true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(pw) != l {
			t.Errorf("expected length %d, got %d", l, len(pw))
		}
	}
}

func TestGeneratePasswordNoCharset(t *testing.T) {
	_, err := generatePassword(16, false, false, false, false, false)
	if err == nil {
		t.Error("expected error when no charsets are selected, got nil")
	}
}

func TestGeneratePasswordExcludeSimilar(t *testing.T) {
	// Generate a long password with similar exclusion enabled
	pw, err := generatePassword(100, true, true, true, true, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, char := range similarChars {
		if strings.ContainsRune(pw, char) {
			t.Errorf("password contains similar character %q, which should have been excluded", char)
		}
	}
}

func TestCalculateEntropy(t *testing.T) {
	// E = L * log2(R)
	// For length 10 and only numbers (R=10): 10 * log2(10) ≈ 10 * 3.3219 = 33.22
	entropy := calculateEntropy(10, false, false, true, false, false)
	expectedMin := 33.2
	expectedMax := 33.3
	if entropy < expectedMin || entropy > expectedMax {
		t.Errorf("expected entropy for 10 digits to be around 33.2, got %f", entropy)
	}
}

func TestGetStrengthLabel(t *testing.T) {
	tests := []struct {
		entropy float64
		want    string
	}{
		{30.0, "Weak (Słabe) 🔴"},
		{50.0, "Medium (Średnie) 🟡"},
		{90.0, "Strong (Silne) 🟢"},
	}

	for _, tt := range tests {
		got := getStrengthLabel(tt.entropy)
		if got != tt.want {
			t.Errorf("getStrengthLabel(%f) = %q, want %q", tt.entropy, got, tt.want)
		}
	}
}
