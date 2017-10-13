package swcq

import "github.com/ethereum/go-ethereum/crypto"

// Event titles
var (
	TitleLogSWCqueueDirectPledge = crypto.Keccak256Hash(
		[]byte("LogSWCqueueDirectPledge(address,uint128,bytes3)"))
)
