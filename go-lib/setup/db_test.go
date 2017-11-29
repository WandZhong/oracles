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

package setup

import (
	chk "gopkg.in/check.v1"
)

type DBSuite struct{}

func (suite *DBSuite) TestNewPgFlagsCheck(c *chk.C) {
	pf := PgFlags{User: strPtr(""), Addr: strPtr("")}
	c.Assert(pf.Check(), chk.ErrorMatches, "pg-user must be specifed")

	pf = PgFlags{User: strPtr("u1"), Addr: strPtr("")}
	c.Assert(pf.Check(), chk.IsNil)

	pf = PgFlags{User: strPtr("u1"), Addr: strPtr("example.com:123")}
	c.Assert(pf.Check(), chk.IsNil)

	pf = PgFlags{User: strPtr("u1"), Addr: strPtr("example.com")}
	c.Assert(pf.Check(), chk.ErrorMatches, "missing port in address")
}
