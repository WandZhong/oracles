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

package members

import (
	"time"

	pgt "github.com/robert-zaremba/go-pgt"
)

// Individual represents the members individual DB record
type Individual struct {
	tableName struct{} `sql:"individual"`

	ID            pgt.UUID `sql:"id,pk"`
	AuthID        string   `sql:"auth_id"`
	Email         string   `sql:"email_address"`
	EmailVerified bool     `sql:"email_address_verified"`
	Country       string   `sql:"fk_country"`
	Language      string   `sql:"fk_language"`

	CreatedAt time.Time `sql:"created_at,notnull"`
	UpdatedAt time.Time `sql:"updated_at,notnull"`
}
