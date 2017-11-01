package swcq

import (
	"time"

	pgt "github.com/robert-zaremba/go-pgt"
)

// Pledge represents swc_queue entry
type Pledge struct {
	tableName struct{} `sql:"swc_queue"`

	ID        pgt.UUID   `sql:"swc_queue_id,pk"`
	UserID    pgt.UUID   `sql:"user_id"`
	Wad       pgt.BigInt `sql:",notnull"`
	CreatedOn time.Time  `sql:"created_on"`
	Currency  string     `sql:",notnull"`
	Direct    bool       `sql:",notnull"`
}
