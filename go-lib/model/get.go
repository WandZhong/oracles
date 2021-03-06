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
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
)

// CheckPgNoRows wraps pg error into errstack.E
func CheckPgNoRows(title string, err error) errstack.E {
	if err == pg.ErrNoRows {
		return errstack.WrapAsReqF(err, "Can't get %q from DB", title)
	}
	return errstack.WrapAsInfF(err, "Can't get %q from DB", title)
}

// ErrNotNoRows check if errors is not Nil and is not ErrNoRows
func ErrNotNoRows(title string, err error) errstack.E {
	if err == nil || err == pg.ErrNoRows {
		return nil
	}
	return errstack.WrapAsInfF(err, "Can't select %q from DB", title)
}
