package libs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strings"
)

func GetTable(db *sqlx.DB, tableName string) (fields []MysqlTableField, err error) {
	var sqlQuery = "show full columns from "
	err = db.Select(&fields, sqlQuery+tableName)
	if err != nil {
		return
	}

	return fields, nil
}

func TableToConst(table MysqlTable) string {
	var sqlQuery = ""
	var structDoc = ""

	var structComment = ""
	var sqlInsert = ""
	structDoc = "\tmysql.SQLXYZ_MODEL\n"
	var fieldDoc string
	for _, field := range table.Fields {
		field.Comment = strings.Replace(field.Comment, "，", ",", -1)
		var arrTmp = strings.Split(field.Comment, ",")
		var fieldAlias = arrTmp[0]
		fieldDoc += "\t" + field.Field + "\tbool\t`db:\"" + field.Field + "\"`"
		fieldDoc += "\t// " + fieldAlias + " \n"
		structDoc += "\t" + field.Field + " "
		if strings.HasPrefix(field.Type, "bigint") || strings.HasPrefix(field.Type, "int") == true || strings.HasPrefix(field.Type, "tinyint") == true {
			structDoc += "\tint64"
		} else if strings.HasPrefix(field.Type, "float") == true {
			structDoc += "\tfloat64"
		} else if strings.HasPrefix(field.Type, "varchar") || strings.HasPrefix(field.Type, "text") || strings.HasPrefix(field.Type, "char") {
			structDoc += "\tstring"
		} else if strings.HasPrefix(field.Type, "datetime") || strings.HasPrefix(field.Type, "date") {
			structDoc += "\ttime.Time"
		}

		structDoc += "\t`db:\"" + field.Field + "\"`"
		structDoc += "\t// " + fieldAlias + " 类型: " + field.Type
		if field.Key == "PRI" {
			structDoc += " 主健字段（Primary Key）"
		}
		if field.Extra == "auto_increment" {
			structDoc += " 自增长字段 "
		}
		if field.Default.String != "" {
			structDoc += " 默认值: " + field.Default.String
		}

		structDoc += "\n"
	}
	var tablenameDoc string = "// 数据库表名\nconst " + table.Alias + "_TableName = \"`" + table.TableName + "`\"\n\n"
	structDoc = "type " + table.Alias + " struct {\n" + structDoc + "}\n"
	fieldDoc = "\ntype " + table.Alias + "Fields struct{\n" + fieldDoc + "}\n"
	sqlInsert = "\tINSERT INTO `" + table.TableName + "` SET " + strings.Join(table.FiledName, "=?,") + "=? \n"
	sqlInsert += "\tUPDATE `" + table.TableName + "` SET " + strings.Join(table.FiledName, "=?,") + "=? \n"
	sqlInsert += "\tDELETE FROM `" + table.TableName + "` WHERE \n"
	sqlQuery = "\tSELECT " + strings.Join(table.FiledName, ",") + " FROM `" + table.TableName + "`\n"
	fileContent := "\n说明:\n\t针对数据库的" + table.Comment + "结构体 " + table.Alias + " 的定义及常用方法, 由db2const工具自动生成, 详细使用请查看: https://github.com/laixyz/db2const\n"

	structComment = "常用SQL:\n" + sqlQuery + sqlInsert

	return tablenameDoc + "/*" + fileContent + structComment + "*/\n" + structDoc + GetStructMethod(table) + fieldDoc + getFieldMethodCode(table)
}

func GetTableComment(db *sqlx.DB, tableName string) (string, string) {
	var sqlQuery = "show create table " + tableName
	var sqlCreate string
	err := db.QueryRow(sqlQuery).Scan(&tableName, &sqlCreate)
	if err != nil {
		fmt.Println(err.Error())
		return "", ""
	}
	sqlCreate = strings.Replace(sqlCreate, "，", ",", -1)
	tmpArr := strings.SplitAfter(sqlCreate, "ENGINE=")
	if len(tmpArr) != 2 {
		return "", ""
	}
	var reg = regexp.MustCompile(`COMMENT='(.*)'$`)

	match := reg.FindStringSubmatch(tmpArr[1])
	if len(match) == 2 {
		if strings.Index(match[1], ",") > 0 {
			tmp := strings.Split(match[1], ",")
			if len(tmp) > 1 {
				return tmp[0], tmp[1]
			} else {
				return tmp[0], ""
			}
		} else {
			return match[1], ""
		}
	}
	return "", ""
}
func GetTables(db *sqlx.DB, ConnectID string) (tables []MysqlTable, err error) {
	var sqlQuery = "show tables"

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return
	}
	defer rows.Close()
	var tableName string

	for rows.Next() {
		var mysqlTable MysqlTable
		err = rows.Scan(&tableName)
		if err != nil {
			fmt.Println("show tables:" + err.Error())
			return
		}
		mysqlTable.Fields, err = GetTable(db, tableName)
		if err != nil {
			return
		}
		var primaryKeyTotal int = 0
		var fieldname []string
		for _, field := range mysqlTable.Fields {
			if field.Key == "PRI" {
				mysqlTable.PrimaryKeyField = field.Field
				primaryKeyTotal++
			}
			fieldname = append(fieldname, "`"+field.Field+"`")
			if field.Field == "State" {
				mysqlTable.HasState = true
			}
			if field.Field == "Created" {
				mysqlTable.HasCreated = true

			}
			if field.Field == "Updated" {
				mysqlTable.HasUpdated = true
			}
			if field.Field == "Deleted" {
				mysqlTable.HasDeleted = true
			}
			if strings.HasPrefix(field.Type, "datetime") || strings.HasPrefix(field.Type, "date") {
				mysqlTable.HasTime = true
			}
		}
		if primaryKeyTotal == 1 {
			mysqlTable.IsOnlyPrimary = true
		}

		if mysqlTable.HasState && mysqlTable.HasCreated && mysqlTable.HasUpdated && mysqlTable.HasDeleted {
			mysqlTable.IsModel = true
		}

		if tmp, tmp2 := GetTableComment(db, tableName); tmp != "" {
			mysqlTable.Alias = tmp
			if tmp2 != "" {
				mysqlTable.Comment = tmp2
			}
		} else {
			mysqlTable.Alias = tableName
		}
		mysqlTable.ConnectID = ConnectID
		mysqlTable.TableName = tableName
		mysqlTable.FiledName = fieldname

		mysqlTable.Doc = TableToConst(mysqlTable)

		tables = append(tables, mysqlTable)

	}
	return tables, nil
}
