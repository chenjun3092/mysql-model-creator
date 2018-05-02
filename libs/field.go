package libs

func getFieldMethodCode(table MysqlTable) string {
	var strMethod string
	strMethod += "\n// Select 指定为true的字段生成select sql语句\n"
	strMethod += "func (self " + table.Alias + "Fields) Select() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_SELECT, self, false)\n"
	strMethod += "}\n"
	strMethod += "\n// SelectAll 所有字段生成select sql语句\n"
	strMethod += "func (self " + table.Alias + "Fields) SelectAll() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_SELECT, self, true)\n"
	strMethod += "}\n"
	strMethod += "\n// Update 指定为true的字段生成update sql语句\n"
	strMethod += "func (self " + table.Alias + "Fields) Update() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_UPDATE, self, false)\n"
	strMethod += "}\n"
	strMethod += "\n// UpdateAll 所有字段生成update sql语句\n"
	strMethod += "func (self " + table.Alias + "Fields) UpdateAll() string { \n"
	strMethod += "\treturn sqlxyz.SQLCreator(\"" + table.TableName + "\", sqlxyz.SQL_UPDATE, self, true)\n"
	strMethod += "}\n"
	return strMethod
}
