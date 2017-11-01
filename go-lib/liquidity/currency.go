package liquidity

import (
	"database/sql/driver"
	"strings"

	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Currency type represent currency ISO name
type Currency [3]byte

var (
	currenciesList = []string{"USD"}
	currencies     = map[string]Currency{}
)

func init() {
	for _, c := range currenciesList {
		currencies[c] = Currency{c[0], c[1], c[2]}
	}
}

// ParseCurrencyErrp converts currency string into Currency type
func ParseCurrencyErrp(curr string, errp errstack.Putter) Currency {
	if curr == "" {
		errp.Put("expecting 3-characters currency ISO code. Got empty code.")
		return [3]byte{}
	}
	curr = strings.ToUpper(curr)
	c, ok := currencies[curr]
	if !ok {
		errp.Put("Unknown Currency: " + curr)
	}
	return c
}

// ParseCurrency  converts currency string into Currency type
func ParseCurrency(curr string) (Currency, errstack.E) {
	errb := errstack.NewBuilder()
	return ParseCurrencyErrp(curr, errb.Putter("currency")), errb.ToReqErr()
}

// String implements fmt.Stringer interface
func (curr Currency) String() string {
	return string(curr[:])
}

// Scan implements sql.Sanner interface
func (curr *Currency) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	s, err := bat.UnsafeToString(src)
	if err != nil {
		return err
	}
	*curr, err = ParseCurrency(s)
	return err
}

// Value implements sql/driver.Valuer
func (curr Currency) Value() (driver.Value, error) {
	return string(curr[:]), nil
}
