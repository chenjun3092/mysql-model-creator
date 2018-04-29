package libs

import (
	"flag"
	"fmt"
	"github.com/laixyz/database/mysql"
	"github.com/laixyz/database/mysqlconfig"
	"github.com/laixyz/ini"
	"os"
	"strings"
)

func Exec() {
	var arg []string = os.Args
	if arg[1] == "version" {
		fmt.Println("\n " + ProjectName + " v " + Version)
		fmt.Println("\n 使用:")
		fmt.Println(" 生成所有表:\n db2const -conf=./test.conf -dist=../model -connect=default")
		fmt.Println(" 只生成members表:\n db2const -conf=./test.conf -dist=../model -connect=default -table=members")
		fmt.Println(" 只生成members和members_messages表:\n db2const -conf=./test.conf -dist=../model -connect=default -table=members,members_messages \n")
		fmt.Println(" 配置文件范例test.conf:")
		fmt.Println("[mysql]")
		fmt.Println("host=localhost")
		fmt.Println("user=test")
		fmt.Println("password=test")
		fmt.Println("db=test")
		fmt.Println("port=3306")
		fmt.Println("charset=utf8")
		return
	}
	var config_file string
	flag.StringVar(&config_file, "conf", "./config.conf", "请使用指定配置文件目录")
	var dist_path string
	flag.StringVar(&dist_path, "dist", "./dist", "指定输出文件目录")

	var ConnectID string
	flag.StringVar(&ConnectID, "connect", "default", "指定使用数据库的配置的ConnectID")

	var dest_table string
	flag.StringVar(&dest_table, "table", "", "表名, 缺省时则生成所有表, 指定表名则只生成指定表的文件,多个表时以半角逗号隔开")

	flag.Parse()
	config, err := ini.Load(config_file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if ConnectID == "" {
		ConnectID = "default"
	}
	var destTable map[string]bool = make(map[string]bool)
	if dest_table != "" {
		tablenames := strings.Split(dest_table, ",")
		for _, t := range tablenames {
			t = strings.TrimSpace(t)
			if t != "" {
				if _, ok := destTable[t]; !ok {
					destTable[t] = true
				}
			}
		}
	}
	var packageName string
	packageName, err = GetPackageName(dist_path)
	if err != nil {
		fmt.Println("发生错误: ", err.Error())
		return
	}

	var cfg mysqlconfig.Config
	if err = config.Cfg.Section("mysql").MapTo(&cfg); err != nil {
		fmt.Println("发生错误: ", err.Error())
		return
	}
	cfg.ParseTime = true
	err = mysql.Register("default", cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db, err := mysql.Using("default")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tables, err := GetTables(db, ConnectID)
	if err == nil {
		var fileContent string
		for _, table := range tables {
			if dest_table != "" {
				if _, ok := destTable[table.TableName]; !ok {
					continue
				}
			}
			fileContent = GetCode(table, packageName)
			err = FileWrite(dist_path+"/"+table.Alias+".go", fileContent)
			if err != nil {
				fmt.Println("发生错误: ", err.Error())
			}
		}
		fmt.Println("完成，请查看 " + dist_path + " 并在该目录下执行go install 或go build 检查是否发生错误.")
	} else {
		fmt.Println("发生错误: ", err.Error())
	}
}
