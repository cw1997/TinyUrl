// config.go
// author:昌维 [github.com/cw1997]
// date:2017-05-07 19:13:06
package config

import (
	//	"fmt"
	//	"log"
	"strings"

	"github.com/widuu/gojson"

	"util"
)

var (
	configJson []byte
)

func GetConfigFile(configFilePath string) {
	if configFilePath == "" {
		configJson = util.ReadAllFile(util.GetCurrentDirectory() + "/config.json")
	} else {
		configJson = util.ReadAllFile(configFilePath)
	}
}

func json2Struct(json []byte) *gojson.Js {
	// http://www.jb51.net/article/73996.htm
	return gojson.Json(string(json))
}

func Get(key string) string {
	keySlice := strings.Split(key, ".")
	//	fmt.Println(util.Slice2string(keySlice[0:1]))
	//	fmt.Println(configJson)
	value := json2Struct(configJson)
	for _, v := range keySlice {
		//		fmt.Println(v)
		value = value.Get(v)
	}
	return value.Tostring()
	// 下面为配置文件获取模块未写之前的代理测试代码
	//	switch key {
	//	case "website.domain":
	//		return "http://127.0.0.1/"
	//	case "parameter.shorturl":
	//		return "s"
	//	case "parameter.longurl":
	//		return "l"
	//	case "rand.retires":
	//		return "1997"
	//	case "url.length":
	//		return "4"
	//	case "url.chars":
	//		return "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	//	case "database.ip":
	//		return "127.0.0.1"
	//	case "database.port":
	//		return "3306"
	//	case "database.username":
	//		return "root"
	//	case "database.password":
	//		return "root"
	//	case "database.database":
	//		return "tinyurl"
	//	case "database.charset":
	//		return "utf8"
	//	case "database.prefix":
	//		return "tinyurl_"
	//	case "database.maxIdle":
	//		return "2"
	//	case "database.maxOpen":
	//		return "500"
	//	case "database.maxLifetime":
	//		return "60"
	//	case "cache.ip":
	//		return "127.0.0.1"
	//	case "cache.port":
	//		return "6379"
	//	case "cache.username":
	//		return ""
	//	case "cache.password":
	//		return ""
	//	case "cache.prefix":
	//		return "tinyurl:"
	//	case "cache.expire":
	//		return "7200"
	//	case "cache.db":
	//		return "0"
	//	case "cache.maxIdle":
	//		return "1"
	//	case "cache.maxActive":
	//		return "200"
	//	case "cache.idleTimeout":
	//		return "180"
	//	default:
	//		return ""
	//	}
}
