package helper

import (
	"testing"
)

func TestIsValidBankName(t *testing.T) {
	tests := []struct {
		bankName string
		expected bool
	}{
		{"BCA", true},
		{"bca", true},
		{"BNI", true},
		{"bni", true},
		{"Mandiri", true},
		{"mandiri", true},
		{"OCBC", true},
		{"ocbc", true},
		{"Maybank", true},
		{"maybank", true},
		{"Panin", true},
		{"panin", true},
		{"BRI", true},
		{"bri", true},
		{"JAGO", true},
		{"jago", true},
		{"BTPN", true},
		{"btpn", true},
		{"UnknownBank", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.bankName, func(t *testing.T) {
			result := IsValidBankName(tt.bankName)
			if result != tt.expected {
				t.Errorf("IsValidBankName(%v) = %v; want %v", tt.bankName, result, tt.expected)
			}
		})
	}
}
