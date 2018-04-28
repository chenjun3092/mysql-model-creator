package libs

import (
	"os"
	"path/filepath"
)

func GetCode(table MysqlTable, packageName string) string {
	var fileContent = ""

	fileContent += "package " + packageName + "\n\n"
	fileContent += "import (\n"
	fileContent += "\t\"database/sql\"\n"
	//fileContent += "\t\"github.com/jmoiron/sqlx\"\n"
	fileContent += "\t\"github.com/laixyz/database/mysql\"\n"
	fileContent += "\t\"github.com/laixyz/database/npager\"\n"
	fileContent += "\t\"github.com/laixyz/database/sqlxyz\"\n"
	if table.HasCreated || table.HasDeleted || table.HasUpdated || table.HasTime {
		fileContent += "\t\"time\"\n"
	}
	fileContent += ")\n"
	fileContent += "\n"
	fileContent += table.Doc
	return fileContent
}

/*
通过实际路径，得到包名
*/
func GetPackageName(Path string) (string, error) {
	path, err := filepath.Abs(Path)
	if err == nil {
		fileinfo, err := os.Stat(path)
		if err == nil {
			return fileinfo.Name(), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func FileWrite(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
