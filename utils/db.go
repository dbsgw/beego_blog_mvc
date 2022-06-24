package utils

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	mysqladmin, _ := beego.AppConfig.String("mysqladmin")
	mysqlpwd, _ := beego.AppConfig.String("mysqlpwd")
	mysqldb, _ := beego.AppConfig.String("mysqldb")
	mysqlurl, _ := beego.AppConfig.String("mysqlurl")
	mysqlport, _ := beego.AppConfig.String("mysqlport")
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "" + mysqladmin + ":" + mysqlpwd + "@tcp(" + mysqlurl + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                                                                 // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	fmt.Println(err)

	if err != nil {
		fmt.Println("mysql数据库链接失败")
	}else {
		fmt.Println("mysql数据库链接成功")
	}
}
