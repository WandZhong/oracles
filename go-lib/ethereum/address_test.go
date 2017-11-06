package ethereum

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

func checkIsZeroAddr(testcases []string, expected bool, c *C) {
	for _, a := range testcases {
		h := common.HexToAddress(a)
		c.Check(IsZeroAddr(h), Equals, expected,
			Commentf("IsZeroAddr should return %v for: %q (%s)", expected, a, h.Hex()))
	}
}

type AddressSuite struct{}

func (e AddressSuite) TestIsZeroAddr(c *C) {
	c.Assert(IsZeroAddr(ZeroAddress), IsTrue)

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
	checkIsZeroAddr(as, true, c)

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
	checkIsZeroAddr(as, false, c)
}

func (e AddressSuite) TestPgtAddress(c *C) {
	var err error
	addr, err := ToAddress("0xce0d46d924cc8437c806721496599fc3ffa268b9")
	c.Assert(err, IsNil, Commentf("Can't convert correct address"))
	hex := addr.Hex()

	paddr := PgtAddress{addr}
	b, err := paddr.Value()
	c.Assert(err, IsNil, Commentf("Can't serialize PgtAddress"))

	var paddr2 PgtAddress
	c.Assert(paddr2.Scan(b), IsNil, Commentf("Can't deserialize PgtAddress"))
	c.Check(paddr2.Hex(), Equals, hex)

	var paddr3 = new(PgtAddress)
	c.Assert(paddr3.Scan(b), IsNil, Commentf("Can't deserialize PgtAddress to pointer"))
	c.Check(paddr3.Hex(), Equals, hex)
}

func (e AddressSuite) TestHashToAddress(c *C) {
	s := "0x000000000000000000000000d435bbbaa004889f95f54e8232575d87793b42df"
	addr, err := HashToAddress(common.HexToHash(s))
	c.Assert(err, IsNil)
	c.Assert(strings.ToLower(addr.Hex()), Equals, "0x"+s[26:])

	// check wrong address
	s = "0x0000000000000000000000000001bbbaa004889f95f54e8232575d87793b42zz"
	addr, err = HashToAddress(common.HexToHash(s))
	c.Assert(err, Not(IsNil))
}
