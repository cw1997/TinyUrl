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
	for i := 0; i < retries; i++ {
		shortUrl = url.GeneralShortgUrl(longUrl)
		if l, err := getLongUrl(shortUrl); l == "" && err != nil {
			break
		}
	}
	datetime := util.GetTimestamp()
	if db.SetShortUrlByLongUrl(longUrl, shortUrl, datetime, ip) {
		cache.SetShortUrlCache(shortUrl, longUrl)
		return shortUrl
	}
	return ""
}

func getLongUrl(shortUrl string) (string, error) {
	if v, err := cache.Get(shortUrl); err == nil {
		return v, nil
	} else {
		//		log.Println(err)
	}
	if v, err := url.GetLongUrl(shortUrl); err == nil {
		if ok, errCache := cache.SetShortUrlCache(shortUrl, v); ok == true {
			return v, nil
		} else {
			return "", errCache
		}
	} else {
		return "", err
	}
}
