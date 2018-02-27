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

package model

import (
	"github.com/go-pg/pg/orm"
	"github.com/robert-zaremba/errstack"
)

// CheckRowsAffected asserts that expected number of rows has been affected in the
// SQL operation.
func CheckRowsAffected(title string, expected int, res orm.Result, err error) errstack.E {
	if err != nil {
		return errstack.WrapAsDomainF(err, "Query execution error %q", title)
	}
	if res.RowsAffected() != expected {
		return errstack.NewDomainF("Expected to affect %d rows, got: %d", expected,
			res.RowsAffected())
	}
	return nil
}
