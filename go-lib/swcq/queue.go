package swcq

import "time"

//go:generate reform

// Pledge represents swc_queue entry
//reform:swc_queue
type Pledge struct {
	ID        string `reform:"id,pk"`
	UserID    string
	Wad       int
	CreatedOn time.Time `reform:"created_on"`
}
