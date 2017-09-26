package dbu

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	nullbytes = []byte("null")
)

// String is a database.sql.NullString wrapper
type String sql.NullString

// Strings is a slice of strings for valuer interface
type Strings []string

// MarshalJSON implements Marshaler interface
func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return nullbytes, nil
}

// UnmarshalJSON implements Unmarshaler interface
func (s *String) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullbytes) {
		s.String, s.Valid = "", false
		return nil
	}
	err := json.Unmarshal(data, &s.String)
	s.Valid = err == nil
	return err
}

// Scan implements sql.Scanner for the String type
func (s *String) Scan(src interface{}) error {
	if src == nil {
		s.String, s.Valid = "", false
		return nil
	}
	err := UnmarshalString(src, &s.String)
	s.Valid = err == nil
	return err
}

// Value is the valuer for String type. The error is always nil.
func (s String) Value() (driver.Value, error) {
	if s.Valid {
		return s.String, nil
	}
	return nil, nil
}

// NewString created from the specified string
func NewString(s string, emptyToNull bool) String {
	if emptyToNull && s == "" {
		return String{}
	}
	return String{String: s, Valid: true}
}

// FilterValidStrings returns a slice of only valid strings
func FilterValidStrings(ss ...String) []string {
	res := []string{}
	for _, x := range ss {
		if x.Valid {
			res = append(res, x.String)
		}
	}
	return res
}

// Scan implements sql.Scanner for the String slice type
// Scanners take the database value (in this case as a byte slice)
// and sets the value of the type.  Here we cast to a string and
// do a regexp based parse
func (s *Strings) Scan(src interface{}) error {
	str, err := ConvertToString(src)
	if err != nil {
		return err
	}
	parsed, err := parseArray(str)
	if err != nil {
		return err
	}
	*s = Strings(parsed)
	return nil
}

// Value is the valuer for string slice which serializes it via comma
// Returns error if one of strings contain comma
func (s Strings) Value() (driver.Value, error) {
	quoted := make([]string, len(s), len(s))
	for i := range s {
		escapedBackslash := strings.Replace(s[i], "\\", "\\\\", -1)
		escapedQuote := strings.Replace(escapedBackslash, "\"", "\\\"", -1)
		quoted[i] = fmt.Sprintf("\"%s\"", escapedQuote)
	}
	return fmt.Sprintf("{%s}", strings.Join(quoted, ",")), nil
}

// Equals compares if two string slices are equal
func (s Strings) Equals(s2 Strings) bool {
	if len(s) != len(s2) {
		return false
	}
	for i := range s {
		if s[i] != s2[i] {
			return false
		}
	}
	return true
}

// Exclude removes existing items from source
func (s Strings) Exclude(ss ...string) Strings {
	newS := Strings{}
	for _, e := range s {
		if !Strings(ss).Contains(e) {
			newS = append(newS, e)
		}
	}
	return newS
}

// Contains checks if s contains val
func (s Strings) Contains(val string) bool {
	for i := range s {
		if s[i] == val {
			return true
		}
	}
	return false
}

// ContainsAll checks if s contains all vals
func (s Strings) ContainsAll(vals []string) bool {
	for i := range vals {
		if !s.Contains(vals[i]) {
			return false
		}
	}
	return true
}

// ContainsAllSorted checks if s contains all elements from `s2`.
// We assume that both slices are sorted.
func (s Strings) ContainsAllSorted(s2 Strings) bool {
	lens, lens2 := len(s), len(s2)
	if lens2 > lens {
		return false
	}
	var i, j int
	var ok = true
	for ; j < lens2; j++ {
		ok = false
		for ; i < lens; i++ {
			if s[i] == s2[j] {
				i++
				ok = true
				break
			}
		}
		if i == lens || !ok {
			break
		}
	}
	return ok && j >= lens2-1
}

// EmptyToNil returns nil slice if source contains
// only empty elements. Otherwise original slice is returned
func (s Strings) EmptyToNil() Strings {
	var result Strings
	for _, elem := range s {
		if len(elem) > 0 {
			result = append(result, elem)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return s
}

// ForEachPair executes function for each pair inside Strings. See tests for details
func (s Strings) ForEachPair(f func(string, string)) {
	if len(s) < 2 {
		return
	}
	first := s[0]
	for _, second := range s[1:] {
		f(first, second)
		first = second
	}
}

// Map using the specified function
func (s Strings) Map(f func(string) string) Strings {
	var result Strings
	for _, elem := range s {
		result = append(result, f(elem))
	}
	return result
}

// TrimSpace returns Strings instance with all elements transformed by strings.TrimSpace
func (s Strings) TrimSpace() Strings {
	return s.Map(strings.TrimSpace)
}

// TrimMap maps elements which are inside map to empty strings
func (s Strings) TrimMap(m map[string]bool) Strings {
	trim := func(s string) string {
		if m[strings.ToLower(s)] {
			return ""
		}
		return s
	}
	return s.Map(trim)
}

// GetOrEmpty returns n-th element or empty string if i is out of range
func (s Strings) GetOrEmpty(i int) string {
	if i >= len(s) {
		return ""
	}
	return s[i]
}

// Distinct omits every non-first occurrence of a element
func (s Strings) Distinct() Strings {
	if s == nil {
		return nil
	}
	result := Strings{}
	seen := map[string]bool{}
	for _, elem := range s {
		if !seen[elem] {
			seen[elem] = true
			result = append(result, elem)
		}
	}
	return result
}
