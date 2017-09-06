package ethereum

import (
	"testing"
)

func TestToWei(t *testing.T) {
	var cases = []struct {
		Amount int64
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
