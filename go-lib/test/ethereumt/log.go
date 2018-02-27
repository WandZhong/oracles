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

package ethereumt

import (
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/check.v1"
)

// ParseLog parses ethereum log JSON
func ParseLog(data string, c *check.C) *types.Log {
	var log = new(types.Log)
	err := log.UnmarshalJSON([]byte(data))
	c.Assert(err, check.IsNil)
	return log
}
