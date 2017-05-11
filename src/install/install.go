// install.go
// author:昌维 [github.com/cw1997]
// date:2017-05-11 00:15:59
package install

import (
	//	"fmt"
	"strings"

	"config"
	"db"
	"util"
)

func InitDatabase() error {
	sqlFile := util.ReadAllFile(util.GetCurrentDirectory() + "/tinyurl.sql.tpl")
	sql := string(sqlFile)
	//将sql生成脚本中的部分字段根据配置项进行动态修改
	sql = strings.Replace(sql, "_PREFIX_", config.Get("database.prefix"), 1)
	sql = strings.Replace(sql, "_LENGTH_", config.Get("url.length"), 1)
	return db.Execute(sql)
}
