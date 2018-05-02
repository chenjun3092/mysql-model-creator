package libs

import (
	"database/sql"
)

const (
	ProjectName = "MySQL-Model-Creator"
	Version     = "0.9"
	ProjectURL  = "https://github.com/laixyz/mysql-model-creator"
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

var FieldTypes map[string]bool = map[string]bool{"ArrayString": true}
