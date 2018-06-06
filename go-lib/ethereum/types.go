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

package ethereum

import (
	"database/sql/driver"

	"github.com/ethereum/go-ethereum/common"
	bat "github.com/robert-zaremba/go-bat"
)

// PgtHash is an ethereum common.Hash wrapper to provide DB interface
type PgtHash struct {
	common.Hash
}

// Scan implements sql.Sanner interface
func (a *PgtHash) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	s, err := bat.UnsafeToBytes(src)
	if err != nil {
		return err
	}
	return a.UnmarshalText(s)
}

// Value implements sql/driver.Valuer
func (a *PgtHash) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return a.Hex(), nil
}
