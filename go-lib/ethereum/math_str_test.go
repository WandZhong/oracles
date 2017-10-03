package ethereum

import (
	"testing"
)

func TestDropLeadingZeros(t *testing.T) {
	var cases = []struct{ str, expected string }{
		{"", ""},
		{"0", ""},
		{"00000000000000000000000000", ""},
		{"1", "1"},
		{"3", "3"},
		{"7519820", "7519820"},
		{"0000002", "2"},
		{"00200", "200"},
		{"200", "200"}}
	for _, c := range cases {
		if s := dropLeadingZeros(c.str); s != c.expected {
			t.Errorf("%q should equal\n\t %q for %q", s, c.expected, c.str)
		}
	}
}

func TestDropLastZeros(t *testing.T) {
	var cases = []struct{ str, expected string }{
		{"", ""},
		{"0", ""},
		{"00000000000000000000000000", ""},
		{"1", "1"},
		{"3", "3"},
		{"7519820", "751982"},
		{"0000002", "0000002"},
		{"00200", "002"},
		{"200", "2"}}
	for _, c := range cases {
		if s := dropLastZeros(c.str); s != c.expected {
			t.Errorf("%q should equal\n\t %q for %q", s, c.expected, c.str)
		}
	}
}

func TestAfToGigaInt(t *testing.T) {
	var cases = []struct{ str, expected string }{
		{"1", "1000000000"},
		{"2", "2000000000"},
		{"20", "20000000000"},
		{"20.0", "20000000000"},
		{"0.1", "100000000"},
		{"1.2", "1200000000"},
		{"1230.00123", "1230001230000"},
		{"001230.0012300", "1230001230000"},
		{"0.123456789", "123456789"},
	}
	for _, c := range cases {
		s, err := afToGigaInt(c.str)
		if err != nil {
			t.Error("Case ", c.str, " should return no error. ", err)
		} else if s != c.expected {
			t.Errorf("%q should equal\n\t  %q for %q", s, c.expected, c.str)
		}
	}

	errCases = []string{"0.1234567891", "", " ", "1a", "1 2", ".", "0.1"}
	for _, s := range errCases {
		if _, err := afToGigaInt(s); err == nil {
			t.Error("Case ", s, " should return an error.")
		}
	}
}
