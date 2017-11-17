package ethereumt

import (
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/check.v1"
)

// ParseLog parses ethereum log JSON
func ParseLog(data string, c *check.C) *types.Log {
	var log = new(types.Log)
	err := log.UnmarshalJSON([]byte(data))
	c.Assert(err, check.IsNil)
	return log
}
