// Copyright (c) 2017 Sweetbridge Inc.
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
	"encoding/json"
	"strings"

	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Currency type represent currency ISO name
type Currency string

// currencies
var (
	currencies = map[string]Currency{}

	// The Smart Contract defines currency as [3]byte.
	// But here, for simplicity we ware using string (encoding)
	CurrUSD  = Currency("USD")
	CurrcETH = Currency("xETH")
)

func init() {
	var currenciesList = []Currency{CurrUSD, CurrcETH}
	for _, c := range currenciesList {
		currencies[string(c)] = c
	}
}

// ParseCurrencyErrp converts currency string into Currency type
func ParseCurrencyErrp(curr string, errp errstack.Putter) Currency {
	if curr == "" {
		errp.Put("expecting 3-characters currency ISO code. Got empty code.")
		return ""
	}
	if len(curr) > 3 {
		// non fiat currencies has prefix
		curr = curr[:1] + strings.ToUpper(curr[1:])
	}
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

// Bytes converts Currency to bytes array which is expected by the SmartContract
func (curr Currency) Bytes() [3]byte {
	return [3]byte{curr[0], curr[1], curr[2]}
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

// MarshalJSON implements Marshaller interface
func (curr Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(curr[:]))
}

// UnmarshalJSON implements Unmarshaller interface
func (curr *Currency) UnmarshalJSON(data []byte) (err error) {
	var l = len(data)
	if l < 2 || data[0] != '"' || data[l-1] != '"' {
		return errstack.NewReqF("JSON string is expected, got: %s", string(data))
	}
	*curr, err = ParseCurrency(string(data[1 : l-1]))
	return err
}
