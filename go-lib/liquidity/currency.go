// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package liquidity

import (
	"database/sql/driver"
	"strings"

	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Currency type represent currency ISO name
type Currency [3]byte

// currencies
var (
	currencies = map[string]Currency{}

	CurrUSD = Currency{'U', 'S', 'D'}
)

func init() {
	var currenciesList = []string{"USD"}
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
