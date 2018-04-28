package libs

func getFieldMethodCode(table MysqlTable) string {
	var strMethod string
	strMethod += "\n// 指定为true的字段生成select sql语句\n"
	strMethod += "func (this " + table.Alias + "Fields) Select() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_SELECT, this, false)\n"
	strMethod += "}\n"
	strMethod += "\n// 所有字段生成select sql语句\n"
	strMethod += "func (this " + table.Alias + "Fields) SelectAll() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_SELECT, this, true)\n"
	strMethod += "}\n"
	strMethod += "\n// 指定为true的字段生成update sql语句\n"
	strMethod += "func (this " + table.Alias + "Fields) Update() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_UPDATE, this, false)\n"
	strMethod += "}\n"
	strMethod += "\n// 所有字段生成update sql语句\n"
	strMethod += "func (this " + table.Alias + "Fields) UpdateAll() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_UPDATE, this, true)\n"
	strMethod += "}\n"
	return strMethod
}
