package dbu

import (
	"bytes"
	"database/sql/driver"

	"github.com/pborman/uuid"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// UUID type to wrap uuid package
type UUID uuid.UUID

// Scan is the scanner for UUID
func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		u = nil
		return nil
	}
	str, err := ConvertToString(value)
	if err != nil {
		return err
	}
	*u = UUID(uuid.Parse(str))
	return nil
}

// Value is the valuer for UUID
func (u UUID) Value() (driver.Value, error) {
	if u.Empty() {
		return nil, nil
	}
	return u.String(), nil
}

// Empty checks whether UUID holds any value
func (u UUID) Empty() bool {
	return u == nil
}

// MarshalJSON implements Marshaller interface
func (u UUID) MarshalJSON() ([]byte, error) {
	return uuid.UUID(u).MarshalBinary()
}

// UnmarshalJSON implements Unmarshaller interface
func (u *UUID) UnmarshalJSON(data []byte) error {
	var uu *uuid.UUID
	if err := uu.UnmarshalBinary(data); err != nil {
		u = nil
		return err
	}
	*u = UUID(*uu)
	return nil
}

// MarshalYAML implements Marshaler interface of YAML
func (u UUID) MarshalYAML() (interface{}, error) {
	return u.MarshalJSON()
}

// UnmarshalYAML implements Unmarshaler interface of YAML
func (u *UUID) UnmarshalYAML(unmarshaler func(interface{}) error) error {
	var s string
	err := unmarshaler(&s)
	if err != nil {
		return err
	}
	*u, err = Parse(s)
	return err
}

// String converts UUID into string
func (u UUID) String() string {
	return uuid.UUID(u).String()
}

// Equals check if two UUIDs are equal
func (u UUID) Equals(u2 UUID) bool {
	return bytes.Equal(u, u2)
}

// Parse parses UUID from string
func Parse(s string) (UUID, errstack.E) {
	u := UUID(uuid.Parse(s))
	if u == nil {
		return u, errstack.NewReqF("Failed to parse %s as UUID", s)
	}
	return u, nil
}

// MustParse parses UUID from string. If `s` is not a proper UUID the the function will panic
func MustParse(s string) UUID {
	u, err := Parse(s)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return u
}

// RandomUUID returns a new random uuid.
func RandomUUID() UUID {
	return UUID(uuid.NewRandom())
}

// UUIDIterator is an interface used by `ExtractUUIDs` function
type UUIDIterator interface {
	Get() UUID
	Next() bool
}

// ExtractUUIDs returns a list of unique ids from given interator
func ExtractUUIDs(i UUIDIterator, optLength ...int) []UUID {
	var length = 0
	if len(optLength) > 0 {
		length = optLength[0]
	}
	set := make(map[string]bool, length)
	res := []UUID{}
	for i.Next() {
		v := i.Get()
		s := v.String()
		if !set[s] {
			set[s] = true
			res = append(res, v)
		}
	}
	return res
}

// UUIDs is a slice of UUID
type UUIDs []UUID

// Scan implements sql Scanner interface
func (ls *UUIDs) Scan(src interface{}) error {
	bs, err := ConvertToBytes(src)
	if err != nil {
		return err
	}
	*ls, err = ParseUUIDArray(bs)
	return err
}

// Value implements sql Valuer interface
func (ls UUIDs) Value() (driver.Value, error) {
	length := 2 // for {}
	if len(ls) > 0 {
		length += 37*len(ls) - 1 // = 36*len(ls) + len(ls)-1;; 36 = uuid str len, len(ls)-1 = amount of ','
	}
	out := make([]byte, length)
	out[0] = '{'
	i := 1
	for _, id := range ls {
		for j, s := 0, id.String(); j < len(s); j++ {
			out[i] = s[j]
			i++
		}
		out[i] = ','
		i++
	}
	out[length-1] = '}' // if len(ls) >0 this will overwrite the final ','

	return bat.UnsafeByteArrayToStr(out), nil
}

// ParseUUIDArray parses UUID array
func ParseUUIDArray(src []byte) (UUIDs, error) {
	if bytes.Equal(src, EmptyArray) {
		return UUIDs{}, nil
	}
	vals := SplitArray(src)
	var results = make(UUIDs, len(vals))
	var err errstack.E
	for i := range vals {
		if results[i], err = Parse(bat.UnsafeByteArrayToStr(vals[i])); err != nil {
			return nil, err
		}
	}
	return results, nil
}
