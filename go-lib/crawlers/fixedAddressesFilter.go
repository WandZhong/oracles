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

package crawlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	"strings"
)

// NewFixedAddressesFilter returns a an instance of FixedAddressFilter setup with a list of addresses
func NewFixedAddressesFilter(addresses []string) (AddressFilter, errstack.E) {
	if len(addresses) == 0 {
		return nil, errstack.NewReq("Addresses must be provided")
	}
	if err := checkAddressesFormat(addresses); err != nil {
		return nil, err
	}

	table := make(map[string]struct{}, len(addresses))
	for _, v := range addresses {
		table[strings.ToUpper(v)] = struct{}{}
	}
	return &fixedAddressFilter{table}, nil
}

// fixedAddressFilter is an implementation of AddressFilter that uses a fixed list of addresses as a while list
type fixedAddressFilter struct {
	addresses map[string]struct{}
}

// MatchesNone returns true, if none of the addreses provided match the filter.
func (r *fixedAddressFilter) MatchesNone(addresses ...string) (bool, errstack.E) {
	if len(addresses) == 0 {
		return false, errstack.NewReq("Address(ses) are mandatory to initialize a filter")
	}
	if err := checkAddressesFormat(addresses); err != nil {
		return false, err
	}

	result := true
	for _, v := range addresses {
		if _, ok := r.addresses[strings.ToUpper(v)]; ok {
			result = result && !ok
		}
	}
	return result, nil
}

// Checks the format of the provided list of addresses
func checkAddressesFormat(addresses []string) errstack.E {
	var rejects []string
	for _, a := range addresses {
		if !common.IsHexAddress(a) {
			rejects = append(rejects, a)
		}
	}
	if len(rejects) > 0 {
		return errstack.NewReqF("Invalid address(es) format", rejects)
	}
	return nil
}
