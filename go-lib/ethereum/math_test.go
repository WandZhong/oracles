package ethereum

import (
	"testing"
)

func TestDecimalSuffixes(t *testing.T) {
	var cases = []struct {
		idx      int
		expected string
	}{
		{0, "000000000"},
		{7, "00"},
		{9, ""}}
	for _, c := range cases {
		if decimalSuffixes[c.idx] != c.expected {
			t.Error("Suffix ", c.idx, " must equal ", c.expected)
		}
	}
}

func TestToWei(t *testing.T) {
	var cases = []struct {
		Amount uint64
		Str    string
	}{
		{0, "0"},
		{1, "1000000000000000000"},
		{999, "999000000000000000000"}}
	for _, c := range cases {
		w := ToWei(c.Amount)
		wstr := w.String()
		if wstr != c.Str {
			t.Error("Expected ", wstr, " to equal ", c.Str)
		}

		if a := WeiToInt(w); a != c.Amount {
			t.Error("Convertion back to int64 doesn't work. Expected ", a, " to equal ", c.Amount)
		}
	}
}
