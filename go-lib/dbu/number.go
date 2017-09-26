package dbu

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	bat "github.com/robert-zaremba/go-bat"
)

// Float64 is a database.sql.NullString Float64
type Float64 sql.NullFloat64

// Float64s is a slice of floats for valuer interface
type Float64s []float64

// Ints is a slice of long integers for valuer interface
type Ints []int64

// Int64 is a database.sql.Int64 wrapper
type Int64 sql.NullInt64

// Scan implements scan methods for scanner
func (f *Float64s) Scan(src interface{}) error {
	bs, err := ConvertToBytes(src)
	if err != nil {
		return err
	}
	*f, err = ParseFloatArray(bs)
	return err
}

// Value is the valuer for float slice
func (f Float64s) Value() (driver.Value, error) {
	ff := []float64(f)
	res := make([]string, len(ff), len(ff))
	for i, v := range ff {
		res[i] = fmt.Sprintf("%.2f", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(res, ",")), nil
}

// MarshalJSON implements Marshaler interface
func (s Float64) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return bat.UnsafeStrToByteArray(bat.F64toa(s.Float64)), nil
	}
	return nullbytes, nil
}

// UnmarshalJSON implements Unmarshaler interface
func (s *Float64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullbytes) {
		s.Valid = false
		return nil
	}
	f, err := strconv.ParseFloat(bat.UnsafeByteArrayToStr(data), 64)
	if err != nil {
		return err
	}
	s.Float64 = f
	s.Valid = true
	return nil
}

// Scan implements sql.Scanner for the Float64 type
func (s *Float64) Scan(src interface{}) error {
	if src == nil {
		s.Float64, s.Valid = 0, false
		return nil
	}
	if asFloat, ok := src.(float64); ok {
		s.Float64, s.Valid = asFloat, true
		return nil
	}
	return error(errors.New("Scan source was not a float64"))
}

// Value is the valuer for Float64 type. The error is always nil.
func (s Float64) Value() (driver.Value, error) {
	if s.Valid {
		return s.Float64, nil
	}
	return nil, nil
}

// MarshalYAML implements Marshaler
func (s Float64) MarshalYAML() (interface{}, error) {
	if s.Valid {
		return s.Float64, nil
	}
	return "", nil
}

// UnmarshalYAML implements Unmarshaler
func (s *Float64) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var v float64
	if err := unmarshal(&v); err != nil {
		return err
	}
	s.Valid = true
	s.Float64 = v
	return nil
}

// String is the stringer implementation for nullable flaot64
func (s Float64) String() string {
	if s.Valid {
		return bat.F64toa(s.Float64)
	}
	return ""
}

// Equals compares if two Float64 slices are equal
func (s Float64) Equals(other *Float64) bool {
	if other == nil {
		return false
	}
	if !s.Valid {
		return !other.Valid
	}
	return s.Float64 == other.Float64
}

// Add adds given float64 into reveiver and returns copy
func (s Float64) Add(other Float64) Float64 {
	sum := s
	if other.Valid {
		if s.Valid {
			sum.Float64 += other.Float64
		} else {
			sum.Float64, sum.Valid = other.Float64, true
		}
	}
	return sum
}

// Scan implements scan methods for scanner
func (ls *Ints) Scan(src interface{}) error {
	bs, err := ConvertToBytes(src)
	if err != nil {
		return err
	}
	*ls, err = ParseInt64Array(bs)
	return err
}

// Value is the valuer for integer slice
func (ls Ints) Value() (driver.Value, error) {
	res := make([]string, len(ls))
	for i, v := range ls {
		res[i] = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("{%s}", strings.Join(res, ",")), nil
}

// MarshalJSON implements Marshaler interface
func (s Int64) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return []byte(fmt.Sprintf("%d", s.Int64)), nil
	}
	return nullbytes, nil
}

// UnmarshalJSON implements Unmarshaler interface
func (s *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullbytes) {
		s.Valid = false
		return nil
	}
	i, err := bat.Atoi64(bat.UnsafeByteArrayToStr(data))
	if err != nil {
		return err
	}
	s.Int64 = i
	s.Valid = true
	return nil
}

// Scan implements sql.Scanner for the if type
func (s *Int64) Scan(src interface{}) error {
	if src == nil {
		s.Int64, s.Valid = 0, false
		return nil
	}
	if asInt, ok := src.(int64); ok {
		s.Int64, s.Valid = asInt, true
		return nil
	}
	if asInt, ok := src.(int32); ok {
		s.Int64, s.Valid = int64(asInt), true
		return nil
	}
	return fmt.Errorf("Scan source was not int64 or int32, but %T", src)
}

// Value is the valuer for if type. The error is always nil.
func (s Int64) Value() (driver.Value, error) {
	if s.Valid {
		return s.Int64, nil
	}
	return nil, nil
}
