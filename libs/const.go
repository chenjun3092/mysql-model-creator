package libs

import (
	"database/sql"
)

type MysqlTable struct {
	TableName       string
	Alias           string
	Comment         string
	Fields          []MysqlTableField
	FiledName       []string
	IsOnlyPrimary   bool
	PrimaryKeyField string
	Doc             string
	HasTime         bool
	IsModel         bool
	HasState        bool
	HasCreated      bool
	HasUpdated      bool
	HasDeleted      bool
	ConnectID       string
}

type MysqlTableField struct {
	Field      string         `db:"Field"`
	Type       string         `db:"Type"`
	Collation  sql.NullString `db:"Collation"`
	Null       string         `db:"Null"`
	Key        string         `db:"Key"`
	Default    sql.NullString `db:"Default"`
	Extra      string         `db:"Extra"`
	Privileges string         `db:"Privileges"`
	Comment    string         `db:"Comment"`
}
