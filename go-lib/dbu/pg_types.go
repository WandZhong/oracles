package dbu

import (
	"database/sql/driver"
	"errors"
)

// TruncInterval is a type for postgres date_trunc first argument - `field`
type TruncInterval string

// Valid TruncInterval values
const (
	Microseconds TruncInterval = "microseconds"
	Milliseconds TruncInterval = "milliseconds"
	Second       TruncInterval = "second"
	Minute       TruncInterval = "minute"
	Hour         TruncInterval = "hour"
	Day          TruncInterval = "day"
	Week         TruncInterval = "week"
	Month        TruncInterval = "month"
	Quarter      TruncInterval = "quarter"
	Year         TruncInterval = "year"
	Decade       TruncInterval = "decade"
	Century      TruncInterval = "century"
	Millennium   TruncInterval = "millennium"
)

// ParseTruncInterval converts s to valid TruncInternal or "" if s is empty and not required
func ParseTruncInterval(s string, required bool) (TruncInterval, error) {
	switch s {
	case "microseconds", "milliseconds", "second", "minute", "hour", "day", "week", "month", "quarter", "year", "decade", "century", "millennium":
		return TruncInterval(s), nil
	}
	if !required && s == "" {
		return "", nil
	}
	return "", errors.New("Wrong value")
}

// Scan implements sql.Scanner for the TruncInterval type
func (tc *TruncInterval) Scan(src interface{}) error {
	if asBytes, ok := src.([]byte); ok {
		*tc = TruncInterval(asBytes)
		return nil
	}
	return error(errors.New("Scan source was not a []byte"))
}

// Value is the valuer for TruncInterval type. The error is always nil.
func (tc TruncInterval) Value() (driver.Value, error) {
	return string(tc), nil
}
