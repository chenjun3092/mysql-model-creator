# mysql-model-creator

功能: 针对mysql数据库的所有表创建golang需要的结构体及查询全表全字段的SQL语句常量

使用:

	mysql-model-creator -conf=./test.conf -dist=../model -connect=default  生成所有表
	mysql-model-creator -conf=./test.conf -dist=../model -connect=default -table=members 只生成members表
	mysql-model-creator -conf=./test.conf -dist=../model -connect=default -table=members,members_messages 只生成members和members_messages表

数据库配置文件格式 test.conf

	[mysql]
	host=localhost
	user=test
	password=test
	db=test
	port=3306
	charset=utf8


生成代码如下

	package model

	import (
		"database/sql"
		"github.com/laixyz/database/mysql"
		"github.com/laixyz/database/npager"
		"github.com/laixyz/database/sqlxyz"
		"time"
	)

	// 数据库表名
	const Member_TableName = "`members`"

	/*
	说明:
		针对数据库的用户结构体 Member 的定义及常用方法, 由db2const工具自动生成, 详细使用请查看: https://github.com/laixyz/db2const
	常用SQL:
		SELECT `UserID`,`Email`,`MobilePhone`,`Password`,`PasswordCreated`,`AuthorizedIP`,`Token`,`TokenExpired`,`UserGroupID`,`RegisterDate`,`RegisterIP`,`ClientAppID`,`ClientHashID`,`State`,`Created`,`Updated`,`Deleted` FROM `members`
		INSERT INTO `members` SET `UserID`=?,`Email`=?,`MobilePhone`=?,`Password`=?,`PasswordCreated`=?,`AuthorizedIP`=?,`Token`=?,`TokenExpired`=?,`UserGroupID`=?,`RegisterDate`=?,`RegisterIP`=?,`ClientAppID`=?,`ClientHashID`=?,`State`=?,`Created`=?,`Updated`=?,`Deleted`=?
		UPDATE `members` SET `UserID`=?,`Email`=?,`MobilePhone`=?,`Password`=?,`PasswordCreated`=?,`AuthorizedIP`=?,`Token`=?,`TokenExpired`=?,`UserGroupID`=?,`RegisterDate`=?,`RegisterIP`=?,`ClientAppID`=?,`ClientHashID`=?,`State`=?,`Created`=?,`Updated`=?,`Deleted`=?
		DELETE FROM `members` WHERE
	*/
	type Member struct {
		mysql.SQLXYZ_MODEL
		UserID          int64     `db:"UserID"`          // 用户ID 类型: int(10) unsigned 主健字段（Primary Key） 自增长字段
		Email           string    `db:"Email"`           // 账户Email 类型: varchar(64)
		MobilePhone     string    `db:"MobilePhone"`     // 手机号码 类型: varchar(15)
		Password        string    `db:"Password"`        // 账户密码 类型: varchar(32)
		PasswordCreated time.Time `db:"PasswordCreated"` // 密码创建时间 类型: datetime
		AuthorizedIP    string    `db:"AuthorizedIP"`    // 授权IP 类型: varchar(128)
		Token           string    `db:"Token"`           // TOKEN 类型: varchar(32)
		TokenExpired    time.Time `db:"TokenExpired"`    // token 过期时间 类型: datetime
		UserGroupID     int64     `db:"UserGroupID"`     // 用户分组 类型: int(10) unsigned 默认值: 0
		RegisterDate    time.Time `db:"RegisterDate"`    // 注册日期 类型: datetime
		RegisterIP      string    `db:"RegisterIP"`      // 注册ip 类型: varchar(128)
		ClientAppID     int64     `db:"ClientAppID"`     // 渠道商ID 类型: int(10) unsigned 默认值: 0
		ClientHashID    string    `db:"ClientHashID"`    // 客户唯一标识 类型: varchar(64)
		State           int64     `db:"State"`           // 用户状态 类型: tinyint(3) 默认值: 0
		Created         time.Time `db:"Created"`         // 创建时间 类型: datetime
		Updated         time.Time `db:"Updated"`         // 更新时间 类型: datetime
		Deleted         time.Time `db:"Deleted"`         // 删除时间 类型: datetime
	}

	// 新建一个Member 对像，并指定默认值
	func NewMember() Member {
		var this Member
		this.ConnectID = "default"
		this.UserID = 0
		this.Email = ""
		this.MobilePhone = ""
		this.Password = ""
		this.PasswordCreated = time.Now()
		this.AuthorizedIP = ""
		this.Token = ""
		this.TokenExpired = time.Now()
		this.UserGroupID = 0
		this.RegisterDate = time.Now()
		this.RegisterIP = ""
		this.ClientAppID = 0
		this.ClientHashID = ""
		this.State = 0
		this.Created = time.Now()
		this.Updated = time.Unix(0, 0)
		this.Deleted = time.Unix(0, 0)
		return this
	}

	// 检查数据库连接是否正常
	func (this *Member) Ping() (err error) {
		this.ConnectID = "default"
		this.DB, err = mysql.Using("default")
		if err != nil {
			return err
		}
		return nil
	}

	// 根据条件查找一条记录
	func FindMember(Where string) (this Member, exists bool, err error) {
		err = this.Ping()
		if err != nil {
			return this, false, err
		}
		var query = "SELECT `UserID`,`Email`,`MobilePhone`,`Password`,`PasswordCreated`,`AuthorizedIP`,`Token`,`TokenExpired`,`UserGroupID`,`RegisterDate`,`RegisterIP`,`ClientAppID`,`ClientHashID`,`State`,`Created`,`Updated`,`Deleted` FROM `members` WHERE " + Where
		err = this.DB.QueryRow(query).Scan(&this.UserID, &this.Email, &this.MobilePhone, &this.Password, &this.PasswordCreated, &this.AuthorizedIP, &this.Token, &this.TokenExpired, &this.UserGroupID, &this.RegisterDate, &this.RegisterIP, &this.ClientAppID, &this.ClientHashID, &this.State, &this.Created, &this.Updated, &this.Deleted)
		if err == nil {
			return this, true, nil
		} else if err == sql.ErrNoRows {
			return this, false, nil
		} else {
			return this, false, err
		}
	}

	// 根据条件查找一条记录, 条件实例: Find("`State`!=-1")
	func (this *Member) Find(Where string) (exists bool, err error) {
		err = this.Ping()
		if err != nil {
			return false, err
		}
		var query = "SELECT `UserID`,`Email`,`MobilePhone`,`Password`,`PasswordCreated`,`AuthorizedIP`,`Token`,`TokenExpired`,`UserGroupID`,`RegisterDate`,`RegisterIP`,`ClientAppID`,`ClientHashID`,`State`,`Created`,`Updated`,`Deleted` FROM `members` WHERE " + Where
		err = this.DB.QueryRow(query).Scan(&this.UserID, &this.Email, &this.MobilePhone, &this.Password, &this.PasswordCreated, &this.AuthorizedIP, &this.Token, &this.TokenExpired, &this.UserGroupID, &this.RegisterDate, &this.RegisterIP, &this.ClientAppID, &this.ClientHashID, &this.State, &this.Created, &this.Updated, &this.Deleted)
		if err == nil {
			return true, nil
		} else if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	// 根据条件查询一个结果集, 条件实例: FindAll("`State`!=-1")
	func (this *Member) FindAll(Where string) (data []Member, total int, err error) {
		err = this.Ping()
		if err != nil {
			return data, 0, err
		}
		var query = "SELECT `UserID`,`Email`,`MobilePhone`,`Password`,`PasswordCreated`,`AuthorizedIP`,`Token`,`TokenExpired`,`UserGroupID`,`RegisterDate`,`RegisterIP`,`ClientAppID`,`ClientHashID`,`State`,`Created`,`Updated`,`Deleted` FROM `members` WHERE " + Where
		err = this.DB.Select(&data, query)
		if err == nil {
			return data, len(data), nil
		} else if err == sql.ErrNoRows {
			return data, 0, nil
		} else {
			return data, 0, err
		}
	}

	// 根据条件查询一个分页结果集, 条件实例: Pager("`State`!=-1", "ID DESC", 1, 50)
	func (this *Member) Pager(Where string, OrderBy string, Page, PageSize int64) (p npager.Pager, total int, err error) {
		err = this.Ping()
		if err != nil {
			return p, 0, err
		}
		var sqlTotal = "SELECT count(*) as Total FROM `members` WHERE " + Where
		var RecordCount int64
		err = this.DB.QueryRow(sqlTotal).Scan(&RecordCount)
		if err != nil {
			return p, 0, err
		}
		p = npager.NewPager(Page, RecordCount, PageSize)
		var Data []Member
		if RecordCount > 0 {
			var query = "SELECT `UserID`,`Email`,`MobilePhone`,`Password`,`PasswordCreated`,`AuthorizedIP`,`Token`,`TokenExpired`,`UserGroupID`,`RegisterDate`,`RegisterIP`,`ClientAppID`,`ClientHashID`,`State`,`Created`,`Updated`,`Deleted` FROM `members` WHERE " + Where + " ORDER BY " + OrderBy
			err = this.DB.Select(&Data, query+" limit ?,?", p.Offset, p.PageSize)
			if err == sql.ErrNoRows {
				return p, 0, nil
			} else if err != nil {
				return p, 0, err
			}
			p.Data = Data
		}
		return p, len(Data), nil
	}

	// 写入一条完整记录
	func (this *Member) Save() (result sql.Result, err error) {
		err = this.Ping()
		if err != nil {
			return result, err
		}
		var query = "INSERT INTO `members` SET Email=?, MobilePhone=?, Password=?, PasswordCreated=?, AuthorizedIP=?, Token=?, TokenExpired=?, UserGroupID=?, RegisterDate=?, RegisterIP=?, ClientAppID=?, ClientHashID=?, State=?, Created=?, Updated=?, Deleted=?"
		result, err = this.DB.Exec(query, this.Email, this.MobilePhone, this.Password, this.PasswordCreated, this.AuthorizedIP, this.Token, this.TokenExpired, this.UserGroupID, this.RegisterDate, this.RegisterIP, this.ClientAppID, this.ClientHashID, this.State, this.Created, this.Updated, this.Deleted)
		return result, err
	}

	// 更新一条完整记录，如果是单一主键会自动忽略主键值的更新
	func (this *Member) Update(Where string) (result sql.Result, err error) {
		err = this.Ping()
		if err != nil {
			return result, err
		}
		this.Updated = time.Now()
		var query = "UPDATE `members` SET Email=?, MobilePhone=?, Password=?, PasswordCreated=?, AuthorizedIP=?, Token=?, TokenExpired=?, UserGroupID=?, RegisterDate=?, RegisterIP=?, ClientAppID=?, ClientHashID=?, State=?, Created=?, Updated=?, Deleted=?` WHERE " + Where
		result, err = this.DB.Exec(query, this.Email, this.MobilePhone, this.Password, this.PasswordCreated, this.AuthorizedIP, this.Token, this.TokenExpired, this.UserGroupID, this.RegisterDate, this.RegisterIP, this.ClientAppID, this.ClientHashID, this.State, this.Created, this.Updated, this.Deleted)
		return result, err
	}

	// 标注记录删除状态及时间 State=-1 作为删除状态
	func (this *Member) Delete(Where string) (result sql.Result, err error) {
		err = this.Ping()
		if err != nil {
			return result, err
		}
		var query = "UPDATE `members`SET `State`=-1, `Deleted`=? WHERE " + Where
		result, err = this.DB.Exec(query, time.Now())
		return result, err
	}

	// 根据条件物理删除一条记录，删除后无法恢复
	func (this *Member) PhysicallyDelete(Where string) (result sql.Result, err error) {
		err = this.Ping()
		if err != nil {
			return result, err
		}
		var query = "DELETE FROM `members` WHERE " + Where
		result, err = this.DB.Exec(query)
		return result, err
	}

	type MemberFields struct {
		UserID          bool `db:"UserID"`          // 用户ID
		Email           bool `db:"Email"`           // 账户Email
		MobilePhone     bool `db:"MobilePhone"`     // 手机号码
		Password        bool `db:"Password"`        // 账户密码
		PasswordCreated bool `db:"PasswordCreated"` // 密码创建时间
		AuthorizedIP    bool `db:"AuthorizedIP"`    // 授权IP
		Token           bool `db:"Token"`           // TOKEN
		TokenExpired    bool `db:"TokenExpired"`    // token 过期时间
		UserGroupID     bool `db:"UserGroupID"`     // 用户分组
		RegisterDate    bool `db:"RegisterDate"`    // 注册日期
		RegisterIP      bool `db:"RegisterIP"`      // 注册ip
		ClientAppID     bool `db:"ClientAppID"`     // 渠道商ID
		ClientHashID    bool `db:"ClientHashID"`    // 客户唯一标识
		State           bool `db:"State"`           // 用户状态
		Created         bool `db:"Created"`         // 创建时间
		Updated         bool `db:"Updated"`         // 更新时间
		Deleted         bool `db:"Deleted"`         // 删除时间
	}

	// 指定为true的字段生成select sql语句
	func (this MemberFields) Select() string {
		return sqlxyz.SQLCreator("members", sqlxyz.SQL_SELECT, this, false)
	}

	// 所有字段生成select sql语句
	func (this MemberFields) SelectAll() string {
		return sqlxyz.SQLCreator("members", sqlxyz.SQL_SELECT, this, true)
	}

	// 指定为true的字段生成update sql语句
	func (this MemberFields) Update() string {
		return sqlxyz.SQLCreator("members", sqlxyz.SQL_UPDATE, this, false)
	}

	// 所有字段生成update sql语句
	func (this MemberFields) UpdateAll() string {
		return sqlxyz.SQLCreator("members", sqlxyz.SQL_UPDATE, this, true)
	}
