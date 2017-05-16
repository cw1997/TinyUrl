// action.go
// author:昌维 [github.com/cw1997]
// date:2017-05-09 09:00:57

package web

import (
	"log"
	"strconv"

	"cache"
	"config"
	"db"
	"url"
	"util"
)

func storeUrl(longUrl string, ip string) string {
	var shortUrl string
	// 遇到随机数碰撞情况重试
	retries, err := strconv.Atoi(config.Get("rand.retires"))
	if err != nil {
		log.Println(err)
		return ""
	}
	//	为了性能，允许插入时间戳有少量误差
	datetime := util.GetTimestamp()
	for i := 0; i < retries; i++ {
		shortUrl = url.GeneralShortgUrl(longUrl)
		//		此处逻辑有误：当事务隔离级别为“Repeatable read”，按照老方法判断会出现幻读现象
		//		if l, err := getLongUrl(shortUrl); l == "" && err != nil {
		//			break
		//		}
		if db.SetShortUrlByLongUrl(longUrl, shortUrl, datetime, ip) {
			cache.SetShortUrlCache(shortUrl, longUrl)
			return shortUrl
		}
	}
	return ""
}

func getLongUrl(shortUrl string) (string, error) {
	if v, err := cache.Get(config.Get("cache.prefix") + shortUrl); err == nil {
		return v, nil
	} else {
		//		log.Println(err)
	}
	if v, err := url.GetLongUrl(shortUrl); err == nil {
		_, errCache := cache.SetShortUrlCache(shortUrl, v)
		return v, errCache
	} else {
		return "", err
	}
}
