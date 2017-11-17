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

package ethereum

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// MustParseABI parses abi string
func MustParseABI(name string, ctrABI string) abi.ABI {
	a, err := abi.JSON(strings.NewReader(ctrABI))
	if err != nil {
		logger.Fatal("Can't parse contract abi", "name", name, err)
	}
	return a
}

// MustHaveEvents ensures that the ABI object has listed events
func MustHaveEvents(e abi.ABI, eventNames ...string) {
	for _, s := range eventNames {
		if _, ok := e.Events[s]; !ok {
			logger.Fatal("Contract doesn't have requested event", "event_name", s)
		}
	}
}
