package crawlers

import (
	. "gopkg.in/check.v1"
)

type FixedAddressesSuite struct{}

// TestInitOK checks a proper initialisation of a filter
func (r FixedAddressesSuite) TestInitOK(c *C) {
	_, err := NewFixedAddressesFilter([]string{"0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"})
	c.Assert(err, IsNil) // "Creating a filter with no entry should raise an error"
}

// TestInitWithoutAddress checks initialization of a filter with no address provided
func (r FixedAddressesSuite) TestInitWithoutAddress(c *C) {
	_, err := NewFixedAddressesFilter(nil)
	c.Assert(err, NotNil) // "Creating a filter with no entry should raise an error"
}

// TestInitWithBadAddress checks initialization of a filter with a invalid address value
func (r FixedAddressesSuite) TestInitWithBadAddress(c *C) {
	_, err := NewFixedAddressesFilter([]string{"ABC"})
	c.Assert(err, NotNil) // "Creating a filter with an invalid address"
}

// TestCheckWithoutAddress checks the function invoqued with no address raises an error
func (r FixedAddressesSuite) TestCheckWithoutAddress(c *C) {
	filter, _ := NewFixedAddressesFilter([]string{"0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"})
	_, err := filter.MatchesNone()
	c.Assert(err, NotNil) // "Address(es) to be checked are missing"
}

// TestCheckValidaddress checks the function invoked with bad address raises an error
func (r FixedAddressesSuite) TestCheckValidaddress(c *C) {
	filter, _ := NewFixedAddressesFilter([]string{"0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"})
	_, err := filter.MatchesNone("ABC")
	c.Assert(err, NotNil) // "Address(es) to be checked is invalid"
}

func (r FixedAddressesSuite) TestMatchesNone(c *C) {
	var (
		scenarii = []struct {
			desc     string
			init     []string
			input    []string
			expected bool
		}{
			{
				"1 Address to check not in filter",
				[]string{"0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"},
				[]string{"0x52bc44d5378309EE2abF1539BF71dE1b7d7bE3b5", "0xFBb1b73C4f0BDa4f67dcA266ce6Ef42f520fBB98"},
				true,
			}, {
				"Addresses to check not in filter",
				[]string{"0x52bc44d5378309EE2abF1539BF71dE1b7d7bE3b5", "0xFBb1b73C4f0BDa4f67dcA266ce6Ef42f520fBB98"},
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D"},
				true,
			}, {
				"Address in filter",
				[]string{"0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D"},
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D"},
				false,
			}, {
				"Addresses in filter",
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D"},
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D", "0x876EabF441B2EE5B5b0554Fd502a8E0600950cFa"},
				false,
			}, {
				"Addresses some in filter some out of filter",
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xB6AaC3b56FF818496B747EA57fCBe42A9aae6218"},
				[]string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0xFE92a3cf1843B5eC7CCf27b2AE753faC1289Fa9D", "0x876EabF441B2EE5B5b0554Fd502a8E0600950cFa"},
				false,
			},
		}
	)

	for _, sc := range scenarii {
		filter, _ := NewFixedAddressesFilter(sc.init)
		actual, _ := filter.MatchesNone(sc.input...)
		c.Check(actual, Equals, sc.expected, Commentf("Input:%s and Filter:%s", sc.input, sc.init))
	}
}
