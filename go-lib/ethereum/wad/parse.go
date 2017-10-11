package wad

import (
	"regexp"
	"strings"

	"github.com/robert-zaremba/errstack"
)

var decimalSuffixes [19]string
var reNumber = regexp.MustCompile(`^\d+(\.\d+)?$`)

func init() {
	s := ""
	for i := len(decimalSuffixes) - 1; i >= 0; i-- {
		decimalSuffixes[i] = s
		s += "0"
	}
}

// afToCoinStr converts float amount to giga integer
func afToCoinStr(amount string) (string, errstack.E) {
	if !reNumber.MatchString(amount) {
		return "", errstack.NewReq("Malformed decimal number")
	}
	commaIdx := strings.IndexRune(amount, '.')
	if len(amount)-commaIdx-1 > 18 {
		return "", errstack.NewReq(
			"Too many decimal places. Maximum 18 after comma is allowed.")
	}

	if commaIdx < 0 { // no decimal part
		intPart := dropLeadingZeros(amount)
		if intPart == "" {
			return "0", nil
		}
		return intPart + decimalSuffixes[0], nil
	}

	intPart := dropLeadingZeros(amount[:commaIdx])
	decPart := dropLastZeros(amount[commaIdx+1:])
	if decPart == "" {
		if intPart == "" {
			return "0", nil
		}
		decPart = decimalSuffixes[0]
	} else {
		decPart += decimalSuffixes[len(decPart)]
	}
	return intPart + decPart, nil
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
