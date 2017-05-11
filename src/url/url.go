// url.go
// author:昌维 [github.com/cw1997]
// date:2017-05-07 20:39:39
package url

import (
	"log"
	"math/rand"
	"strconv"
	//	"strings"

	"config"
	"db"
)

func GeneralShortgUrl(longUrl string) string {
	var shortUrl string
	shortUrlLen, err := strconv.Atoi(config.Get("url.length"))
	if err != nil {
		log.Println(err)
		return ""
	}
	urlCharsSlice := config.Get("url.chars")
	urlCharsLen := len(urlCharsSlice)

	for i := 0; i < shortUrlLen; i++ {
		randNumber := rand.Intn(urlCharsLen)
		shortUrl += urlCharsSlice[randNumber : randNumber+1]
	}
	return shortUrl
}

func GetLongUrl(shortUrl string) (string, error) {
	return db.GetLongUrlByShortUrl(shortUrl)
}
