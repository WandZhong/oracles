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

func failOnErr(err error, msg string, t *testing.T) {
	if err != nil {
		t.Fatal(msg, err)
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

func TestPgtAddress(t *testing.T) {
	var err error
	addr, err := ToAddress("0xce0d46d924cc8437c806721496599fc3ffa268b9")
	failOnErr(err, "Can't convert correct address", t)
	hex := addr.Hex()

	paddr := PgtAddress{addr}
	b, err := paddr.Value()
	failOnErr(err, "Can't serialize PgtAddress", t)

	var paddr2 PgtAddress
	failOnErr(paddr2.Scan(b), "Can't deserialize PgtAddress", t)
	if paddr2.Hex() != hex {
		t.Errorf("Scan * Value is not identity. orig=%v obtained=%v", hex, paddr2.Hex())
	}

	var paddr3 = new(PgtAddress)
	failOnErr(paddr3.Scan(b), "Can't deserialize PgtAddress to nil pointer", t)
	if paddr3.Hex() != hex {
		t.Errorf("Scan * Value is not identity. orig=%v obtained=%v", hex, paddr2.Hex())
	}
}
