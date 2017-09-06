package ethereum

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func checkIsZeroAddr(testcases []string, expected bool, t *testing.T) {
	for _, a := range testcases {
		h := common.HexToAddress(a)
		if IsZeroAddr(h) != expected {
			t.Errorf("IsZeroAddr should return %v for: %q (%s)", expected, a, h.Hex())
		}
	}
}

func TestIsZeroAddr(t *testing.T) {
	// check invalid or zero addresses
	var as = []string{"", "0", "0x0",
		"0x0000000000000000000000000000000000000000",
		"0000000000000000000000000000000000000000",
		"x", "yyy", "0xy", "0y",
		// single digits are invalid addresses
		"1", "2", "9",
		// spaces are not removed
		"12 34", " 1234", "1234 ",
		// even so long addresses are trimmed, it checks the character range
		"yyce0d46d924cc8437c806721496599fc3ffa268b9123"}
	checkIsZeroAddr(as, true, t)

	// check correct addresses
	as = []string{
		"0x0000000000000000000000000000000000000001",
		"0000000000000000000000000000000000000001",
		"0xce0d46d924cc8437c806721496599fc3ffa268b9",
		"ce0d46d924cc8437c806721496599fc3ffa268b9",
		// long addresses are left-trimmed
		"0xce0d46d924cc8437c806721496599fc3ffa268b9123",
		"ce0d46d924cc8437c806721496599fc3ffa268b9123",
		// short addresses are are left-padded with 0
		"0x123", "123", "12", "99", "099"}
	checkIsZeroAddr(as, false, t)
}
