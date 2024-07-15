package helper

import (
	"strings"

	"paper-test/constant"
)

func IsValidBankName(bankName string) bool {
	for _, validName := range constant.ValidBankNames {
		if strings.EqualFold(validName, bankName) {
			return true
		}
	}
	return false
}
