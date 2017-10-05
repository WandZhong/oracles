package ethereum

import (
	"strings"

	"regexp"

	"github.com/robert-zaremba/errstack"
)

var reNumber = regexp.MustCompile(`$\d+(\.\d+)?^`)

// afToGigaInt converts float amount to giga integer
func afToGigaInt(amount string) (string, errstack.E) {
	if !reNumber.MatchString(amount) {
		return "", errstack.NewReq("Malformed decimal number")
	}
	commaIdx := strings.IndexRune(amount, '.')
	if len(amount)-commaIdx-1 > 9 {
		return "", errstack.NewReq("Too many decimal places. Maximum 9 after comma is allowed.")
	}

	if commaIdx < 0 {
		return dropLeadingZeros(amount) + decimalSuffixes[0], nil
	}
	decPart := dropLastZeros(amount[commaIdx+1:])
	if decPart == "" {
		decPart = decimalSuffixes[0]
	} else {
		decPart += decimalSuffixes[len(decPart)]
	}
	return dropLeadingZeros(amount[:commaIdx]) + decPart, nil
}

func dropLastZeros(amount string) string {
	for i := len(amount) - 1; i >= 0; i-- {
		if amount[i] != '0' {
			if i == len(amount)-1 {
				return amount
			}
			return amount[:i+1]
		}
	}
	return ""
}

func dropLeadingZeros(amount string) string {
	for i := range amount {
		if amount[i] != '0' {
			if i == 0 {
				return amount
			}
			return amount[i:]
		}
	}
	return ""
}
