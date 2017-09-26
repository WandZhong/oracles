package dbu

import (
	//	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type UUIDSuite struct{}

func (suite *UUIDSuite) TestUUIDsScan(c *C) {
	// 1. Check empty slices
	var ls UUIDs
	var lsout = new(UUIDs)
	s, err := ls.Value()
	c.Assert(err, IsNil)
	c.Assert(s, Equals, "{}")

	ls = UUIDs{}
	s, _ = ls.Value()
	c.Assert(s, Equals, "{}")
	c.Assert(lsout.Scan(s), IsNil)
	c.Check(*lsout, DeepEquals, ls)

	// 2. Check 1 element slice
	id := RandomUUID()
	ls = UUIDs{id}
	s, _ = ls.Value()
	c.Assert(s, Equals, "{"+id.String()+"}")
	c.Assert(lsout.Scan(s), IsNil)
	c.Check(*lsout, DeepEquals, ls)

	// 3. Check multiple elements
	ls = make(UUIDs, 10)
	for i := 0; i < 10; i++ {
		ls[i] = RandomUUID()
	}
	s, _ = ls.Value()
	c.Assert(lsout.Scan(s), IsNil)
	c.Check(*lsout, DeepEquals, ls)
}
