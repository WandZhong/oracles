package liquidity

import (
	"strings"

	"github.com/robert-zaremba/errstack"
)

// Currency type represent currency ISO name
type Currency [3]byte

var (
	currenciesList = []string{"usd"}
	currencies     = map[string]Currency{}
)

func init() {
	for _, c := range currenciesList {
		currencies[c] = Currency{c[0], c[1], c[2]}
	}
}

// ToCurrency converts currency string into Currency type
func ToCurrency(curr string, errp errstack.Putter) Currency {
	if curr == "" {
		errp.Put("expecting 3-characters currency ISO code. Got empty code.")
		return [3]byte{}
	}
	curr = strings.ToLower(curr)
	c, ok := currencies[curr]
	if !ok {
		errp.Put("Unknown Currency: " + curr)
	}
	return c
}
