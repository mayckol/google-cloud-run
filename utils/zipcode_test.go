package utils

import (
	"testing"
)

func TestZipCode_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  ZipCode
		expected bool
	}{
		{"ValidZipCode", ZipCode("12345678"), true},
		{"InvalidZipCodeWithLetters", ZipCode("1234abcd"), false},
		{"InvalidZipCodeTooShort", ZipCode("12345"), false},
		{"InvalidZipCodeTooLong", ZipCode("123456789"), false},
		{"ValidZipCodeWithDashes", ZipCode("12345-678"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zipCode.IsValid(); got != tt.expected {
				t.Errorf("IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestZipCode_Masked(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  ZipCode
		expected string
	}{
		{"MaskedZipCode", ZipCode("12345678"), "12345-678"},
		{"MaskedZipCodeWithDashes", ZipCode("12345-678"), "12345-678"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zipCode.Masked(); got != tt.expected {
				t.Errorf("Masked() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestZipCode_Raw(t *testing.T) {
	tests := []struct {
		name     string
		zipCode  ZipCode
		expected string
	}{
		{"RawZipCode", ZipCode("12345678"), "12345678"},
		{"RawZipCodeWithDashes", ZipCode("12345-678"), "12345678"},
		{"RawZipCodeWithLetters", ZipCode("1234abcd"), "1234"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zipCode.Raw(); got != tt.expected {
				t.Errorf("Raw() = %v, want %v", got, tt.expected)
			}
		})
	}
}
