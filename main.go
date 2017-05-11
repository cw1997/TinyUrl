/*******************************************
 * TinyUrl Ver 1.0.0 短网址服务
 * author:昌维 [github.com/cw1997/TinyUrl]
 * email:867597730@qq.com
 * date:2017-05-09 11:28:59
 *******************************************/

package main

import (
	"flag"
	"fmt"
	"log"

	"cache"
	"config"
	"db"
	"install"
	"util"
	"web"
)

func main() {
	fmt.Println("Welcome to use TinyUrl.")
	fmt.Println("If you have any problems,")
	fmt.Println("you can contact us at any time by issue.")
	fmt.Println("http://github.com/cw1997/TinyUrl/issues")
	fmt.Println("----------------------------------------")
	fmt.Println("start at " + util.GetDatetime())
	fmt.Println("----------------------------------------")

	// configFilePath接收用户从命令行中输入的config参数
	var configFilePath string
	//flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的。
	flag.StringVar(&configFilePath, "config", "", "What's the configuration file's path is?")
	flag.Parse() //解析输入的参数

	// 获取配置文件
	fmt.Println("read the config file.")
	config.GetConfigFile(configFilePath)

	//打开redis缓存连接池
	fmt.Println("create the redis connection pool.")
	cacheErr := cache.ConnectPool()
	if cacheErr != nil {
		log.Fatalf("create redis connection pool failed: %v\n", cacheErr)
	}

	//打开mysql缓存连接池
	fmt.Println("create the mysql connection pool.")
	_, dbErr := db.ConnectPool()
	if dbErr != nil {
		log.Fatalf("create mysql connection pool failed: %v\n", dbErr)
	}

	//执行数据库初始化工作
	fmt.Println("create the database table if it isn't exist.")
	initDatabaseErr := install.InitDatabase()
	if initDatabaseErr != nil {
		log.Fatalf("create the database table failed: %v\n", initDatabaseErr)
	}

	//	启动web服务
	fmt.Println("start the web server.")
	fmt.Println("----------------------------------------")
	web.StartHttpServer()

}
