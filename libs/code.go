package libs

import (
	"strings"
)

func GetStructMethod(table MysqlTable) string {
	var strMethod string
	var tmp string
	var tmpInsert string
	var tmpSelect string
	strMethod += "\n// New" + table.Alias + " 新建一个" + table.Alias + " 对像，并指定默认值\n"
	strMethod += "func New" + table.Alias + "() " + table.Alias + " { \n"
	strMethod += "\tvar self " + table.Alias + "\n"
	strMethod += "\tself.ConnectID=\"" + table.ConnectID + "\"\n"
	for _, field := range table.Fields {

		field.Comment = strings.Replace(field.Comment, "，", ",", -1)
		var arrTmp = strings.Split(field.Comment, ",")
		//var fieldAlias = arrTmp[0]

		var filedType string
		if len(arrTmp) >= 2 {
			filedType = arrTmp[1]
		} else {
			filedType = ""
		}

		if tmpInsert != "" {
			tmpInsert += ", "
		}
		tmpInsert += "self." + field.Field
		if tmpSelect != "" {
			tmpSelect += ", "
		}
		tmpSelect += "&self." + field.Field

		if strings.HasPrefix(field.Type, "varchar") || strings.HasPrefix(field.Type, "text") || strings.HasPrefix(field.Type, "char") {
			if field.Default.String != "" {
				field.Default.String = strings.Replace(field.Default.String, "'", "", -1)
				field.Default.String = strings.Replace(field.Default.String, " ", "", -1)
				if field.Default.String == "NULL" {
					field.Default.String = ""
				}
				tmp += "\tself." + field.Field + " = \"" + field.Default.String + "\"\n"
			} else {
				if filedType == "ArrayString" {
					tmp += "\tself." + field.Field + " = sqlxyz.ArrayString{}\n"
				} else {
					tmp += "\tself." + field.Field + " = \"\"\n"
				}
			}

		} else if strings.HasPrefix(field.Type, "datetime") || strings.HasPrefix(field.Type, "date") {
			if field.Field == "Deleted" {
				tmp += "\tself." + field.Field + " = time.Unix(0, 0)\n"
			} else if field.Field == "Updated" {
				tmp += "\tself." + field.Field + " = time.Unix(0, 0)\n"
			} else {
				tmp += "\tself." + field.Field + " = time.Now()\n"
			}
		} else {
			if field.Default.String != "" {
				tmp += "\tself." + field.Field + " = " + field.Default.String + "\n"
			} else {
				tmp += "\tself." + field.Field + " = 0\n"
			}
		}
	}
	strMethod += tmp
	strMethod += "\treturn self\n"
	strMethod += "}\n"
	strMethod += "\n// Ping 检查数据库连接是否正常\n"
	strMethod += "func (self *" + table.Alias + ") Ping() (err error) {\n"
	strMethod += "\tself.ConnectID = \"" + table.ConnectID + "\"\n"
	strMethod += "\treturn self.SQLXYZ_MODEL.Ping()\n"
	strMethod += "}\n"

	tmp = "db *sqlx.DB"
	strMethod += "\n// Find 根据条件查找一条记录\n"
	strMethod += "func Find" + table.Alias + "(Where string) (self " + table.Alias + ", exists bool, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn self, false, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar query = \"SELECT " + strings.Join(table.FiledName, ",") + " FROM `" + table.TableName + "` WHERE \" + Where\n"
	strMethod += "\terr = self.DB.QueryRow(query).Scan(" + tmpSelect + ")\n"
	strMethod += "\tif err == nil { \n"
	strMethod += "\t\treturn self, true, nil\n"
	strMethod += "\t} else if err == sql.ErrNoRows {\n"
	strMethod += "\t\treturn self, false, nil\n"
	strMethod += "\t} else {\n"
	strMethod += "\t\treturn self, false, err\n"
	strMethod += "\t}\n"
	strMethod += "}\n"
	strMethod += "\n// Find 根据条件查找一条记录, 条件实例: Find(\"`State`!=-1\")\n"
	strMethod += "func (self *" + table.Alias + ") Find(Where string) (exists bool, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn false, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar query = \"SELECT " + strings.Join(table.FiledName, ",") + " FROM `" + table.TableName + "` WHERE \" + Where\n"
	strMethod += "\terr = self.DB.QueryRow(query).Scan(" + tmpSelect + ")\n"
	strMethod += "\tif err == nil { \n"
	strMethod += "\t\treturn true, nil\n"
	strMethod += "\t} else if err == sql.ErrNoRows {\n"
	strMethod += "\t\treturn false, nil\n"
	strMethod += "\t} else {\n"
	strMethod += "\t\treturn false, err\n"
	strMethod += "\t}\n"
	strMethod += "}\n"

	strMethod += "\n// FindAll 根据条件查询一个结果集, 条件实例: FindAll(\"`State`!=-1\")\n"
	strMethod += "func (self *" + table.Alias + ") FindAll(Where string) (data []" + table.Alias + ", total int, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn data, 0, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar query = \"SELECT " + strings.Join(table.FiledName, ",") + " FROM `" + table.TableName + "` WHERE \" + Where\n"
	strMethod += "\terr = self.DB.Select(&data, query)\n"
	strMethod += "\tif err == nil { \n"
	strMethod += "\t\treturn data, len(data), nil\n"
	strMethod += "\t} else if err == sql.ErrNoRows {\n"
	strMethod += "\t\treturn data, 0, nil\n"
	strMethod += "\t} else {\n"
	strMethod += "\t\treturn data, 0, err\n"
	strMethod += "\t}\n"
	strMethod += "}\n"

	strMethod += "\n// Pager 根据条件查询一个分页结果集, 条件实例: Pager(\"`State`!=-1\", \"ID DESC\", 1, 50)\n"
	strMethod += "func (self *" + table.Alias + ") Pager(Where string, OrderBy string, Page, PageSize int64) (p npager.Pager, total int, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn p, 0, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar sqlTotal = \"SELECT count(*) as Total FROM `" + table.TableName + "` WHERE \" + Where\n"
	strMethod += "\tvar RecordCount int64\n"
	strMethod += "\terr = self.DB.QueryRow(sqlTotal).Scan(&RecordCount)\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn p, 0, err\n"
	strMethod += "\t}\n"
	strMethod += "\tp = npager.NewPager(Page, RecordCount, PageSize)\n"
	strMethod += "\tvar Data []" + table.Alias + "\n"
	strMethod += "\tif RecordCount > 0 {\n"

	strMethod += "\t\tvar query = \"SELECT " + strings.Join(table.FiledName, ",") + " FROM `" + table.TableName + "` WHERE \" + Where + \" ORDER BY \" + OrderBy\n"
	strMethod += "\t\terr = self.DB.Select(&Data, query+\" limit ?,?\", p.Offset, p.PageSize)\n"
	strMethod += "\t\tif err == sql.ErrNoRows {\n"
	strMethod += "\t\t\treturn p, 0, nil\n"
	strMethod += "\t\t} else if err != nil {\n"
	strMethod += "\t\t\treturn p, 0, err\n"
	strMethod += "\t\t}\n"
	strMethod += "\t\tp.Data = Data\n"

	strMethod += "\t}\n"
	strMethod += "\treturn p, len(Data), nil\n"
	strMethod += "}\n"

	strMethod += GetSaveFunc(table)
	strMethod += GetUpdateFunc(table)

	strMethod += GetDeleteFunc(table)
	return strMethod
}
func GetSaveFunc(table MysqlTable) string {
	var strMethod string
	var tmpInsert string
	var fieldname []string
	for _, field := range table.Fields {
		if table.IsOnlyPrimary == true && field.Field == table.PrimaryKeyField && field.Extra == "auto_increment" {
			//跳过自增长主键更新
			continue
		}
		fieldname = append(fieldname, field.Field)
		if tmpInsert != "" {
			tmpInsert += ", "
		}
		tmpInsert += "self." + field.Field
	}
	strMethod += "\n// Save 写入一条完整记录\n"
	strMethod += "func (self *" + table.Alias + ") Save() (result sql.Result, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn result, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar query = \"INSERT INTO `" + table.TableName + "` SET " + strings.Join(fieldname, "=?, ") + "=?\"\n"
	strMethod += "\tresult, err = self.DB.Exec(query, " + tmpInsert + ")\n"
	strMethod += "\treturn result, err\n"
	strMethod += "}\n"
	return strMethod
}

func GetUpdateFunc(table MysqlTable) string {
	var strMethod string
	var tmpInsert string
	var fieldname []string
	for _, field := range table.Fields {
		if table.IsOnlyPrimary == true && field.Field == table.PrimaryKeyField {
			//跳过唯一主键更新
			continue
		}
		fieldname = append(fieldname, field.Field)
		if tmpInsert != "" {
			tmpInsert += ", "
		}
		tmpInsert += "self." + field.Field
	}
	strMethod += "\n// Update 更新一条完整记录，如果是单一主键会自动忽略主键值的更新\n"
	strMethod += "func (self *" + table.Alias + ") Update(Where string) (result sql.Result, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn result, err\n"
	strMethod += "\t}\n"
	if table.HasUpdated {
		strMethod += "\tself.Updated = time.Now()\n"
	}
	strMethod += "\tvar query = \"UPDATE `" + table.TableName + "` SET " + strings.Join(fieldname, "=?, ") + "=?" + "` WHERE \" + Where\n"
	strMethod += "\tresult, err = self.DB.Exec(query, " + tmpInsert + ")\n"
	strMethod += "\treturn result, err\n"
	strMethod += "}\n"
	return strMethod
}
func GetDeleteFunc(table MysqlTable) string {
	var strMethod string
	if table.IsModel {
		strMethod += "\n// Delete 标注记录删除状态及时间 State=-1 作为删除状态\n"
		strMethod += "func (self *" + table.Alias + ") Delete(Where string) (result sql.Result, err error) { \n"
		strMethod += "\terr = self.Ping()\n"
		strMethod += "\tif err != nil {\n"
		strMethod += "\t\treturn result, err\n"
		strMethod += "\t}\n"
		strMethod += "\tvar query = \"UPDATE `" + table.TableName + "`SET `State`=-1, `Deleted`=? WHERE \" + Where\n"
		strMethod += "\tresult, err = self.DB.Exec(query, time.Now())\n"
		strMethod += "\treturn result, err\n"
		strMethod += "}\n"
	}
	strMethod += "\n// PhysicallyDelete 根据条件物理删除一条记录，删除后无法恢复\n"
	strMethod += "func (self *" + table.Alias + ") PhysicallyDelete(Where string) (result sql.Result, err error) { \n"
	strMethod += "\terr = self.Ping()\n"
	strMethod += "\tif err != nil {\n"
	strMethod += "\t\treturn result, err\n"
	strMethod += "\t}\n"
	strMethod += "\tvar query = \"DELETE FROM `" + table.TableName + "` WHERE \" + Where\n"
	strMethod += "\tresult, err = self.DB.Exec(query)\n"
	strMethod += "\treturn result, err\n"
	strMethod += "}\n"
	return strMethod
}
