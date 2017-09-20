package swcq

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type pledgeTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pledgeTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("swc_queue").
func (v *pledgeTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pledgeTableType) Columns() []string {
	return []string{"id", "created_on"}
}

// NewStruct makes a new struct for that view or table.
func (v *pledgeTableType) NewStruct() reform.Struct {
	return new(Pledge)
}

// NewRecord makes a new record for that table.
func (v *pledgeTableType) NewRecord() reform.Record {
	return new(Pledge)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *pledgeTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PledgeTable represents swc_queue view or table in SQL database.
var PledgeTable = &pledgeTableType{
	s: parse.StructInfo{Type: "Pledge", SQLSchema: "", SQLName: "swc_queue", Fields: []parse.FieldInfo{{Name: "ID", PKType: "string", Column: "id"}, {Name: "CreatedOn", PKType: "", Column: "created_on"}}, PKFieldIndex: 0},
	z: new(Pledge).Values(),
}

// String returns a string representation of this struct or record.
func (s Pledge) String() string {
	res := make([]string, 2)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "CreatedOn: " + reform.Inspect(s.CreatedOn, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Pledge) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.CreatedOn,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Pledge) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.CreatedOn,
	}
}

// View returns View object for that struct.
func (s *Pledge) View() reform.View {
	return PledgeTable
}

// Table returns Table object for that record.
func (s *Pledge) Table() reform.Table {
	return PledgeTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Pledge) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Pledge) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Pledge) HasPK() bool {
	return s.ID != PledgeTable.z[PledgeTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Pledge) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = PledgeTable
	_ reform.Struct = new(Pledge)
	_ reform.Table  = PledgeTable
	_ reform.Record = new(Pledge)
	_ fmt.Stringer  = new(Pledge)
)

func init() {
	parse.AssertUpToDate(&PledgeTable.s, new(Pledge))
}
