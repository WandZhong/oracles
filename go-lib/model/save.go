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

package model

import (
	"io"

	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Validated represents model entity which can be validated
type Validated interface {
	Validate() errstack.Builder
}

// DecodeAndSave is handy model function which decodes, validates and save into DB an object.
// `dest` must be a pointer.
func DecodeAndSave(src io.ReadCloser, dest Validated, db *pg.DB) errstack.E {
	if err := bat.DecodeJSON(src, dest); err != nil {
		return err
	}
	if errb := dest.Validate(); errb.NotNil() {
		return errb.ToReqErr()
	}
	err := errstack.WrapAsInf(db.Insert(dest), "Can't insert token")
	return err
}
