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

package encoding

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/robert-zaremba/errstack"
)

// NewCSVReader creates csv.Reader. It assumes that
// * fields are separated with `,`
// * `#` is used for comments
// * accepts trailing commas
// Furthermore it reads and skips given number of records
func NewCSVReader(r io.Reader, skip int) (*csv.Reader, errstack.E) {
	var reader = csv.NewReader(r)
	reader.Comment = '#'
	reader.TrailingComma = true
	for i := 0; i < skip; i++ {
		if _, err := reader.Read(); err != nil {
			return nil, errstack.WrapAsReq(err, "Can't read the CSV header")
		}
	}
	return reader, nil
}

// NewCSVFileReader opens the `fname` file and calls NewCSVReader.
// It also returns file close function as the second field.
func NewCSVFileReader(fname string, skip int) (*csv.Reader, func() error, errstack.E) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, nil, errstack.WrapAsReq(err, "Can't open csv file: "+fname)
	}
	r, err2 := NewCSVReader(f, skip)
	return r, f.Close, err2
}
