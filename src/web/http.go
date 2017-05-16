// http.go
// author:昌维 [github.com/cw1997]
// date:2017-05-07 14:36:25

package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"config"
	//	"url"
)

var (
	WebUrl string
)

type param struct {
	LongUrl  string
	ShortUrl string
}

func (param *param) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlLength, configErr := strconv.Atoi(config.Get("url.length"))
	if configErr != nil {
		log.Fatalln(configErr)
	}
	shortUrl := getPath(r)
	if len(shortUrl) == urlLength {
		longUrl, err := getLongUrl(shortUrl)
		if err != nil {
			longUrl = ""
		}
		log.Println("redirect:" + longUrl)
		http.Redirect(w, r, longUrl, 302)
	}
	param.LongUrl = getQuery(r, config.Get("parameter.longurl"))
	param.ShortUrl = getQuery(r, config.Get("parameter.shorturl"))
	ip := getIP(r)
	errorno := 0
	errormsg := ""
	returnData := make(map[string]interface{})
	if param.LongUrl != "" {
		shortUrl := storeUrl(param.LongUrl, ip)
		if shortUrl == "" {
			errorno = 1
			log.Println("convert the url failed")
			errormsg = "convert the url failed, or system is error."
		}
		returnData["long_url"] = shortUrl
		json, errJson := jsonReturn(Json{
			Error:    errorno,
			ErrorMsg: errormsg,
			Data:     returnData,
		})
		if errJson != nil {
			log.Println(errJson)
		} else {
			io.WriteString(w, string(json))
		}
		return
	}
	if param.ShortUrl != "" {
		longUrl, err := getLongUrl(param.ShortUrl)
		if err != nil {
			errorno = 1
			log.Println(err)
			errormsg = "shorturl not found, or system is error."
		}
		returnData["long_url"] = longUrl
		json, errJson := jsonReturn(Json{
			Error:    errorno,
			ErrorMsg: errormsg,
			Data:     returnData,
		})
		if errJson != nil {
			log.Println(errJson)
		} else {
			io.WriteString(w, string(json))
		}
		return
	}
}

//type URL struct {
//    Scheme   string
//    Opaque   string    // 编码后的不透明数据
//    User     *Userinfo // 用户名和密码信息
//    Host     string    // host或host:port
//    Path     string
//    RawQuery string // 编码后的查询字符串，没有'?'
//    Fragment string // 引用的片段（文档位置），没有'#'
//}
//type Values map[string][]string

func StartHttpServer() {
	ip := config.Get("webserver.ip")
	port := config.Get("webserver.port")
	uri := ip + ":" + port
	WebUrl = "http://" + uri + "/"
	fmt.Println("now you can visit the url to use the TinyUrl:")
	fmt.Println(WebUrl)
	fmt.Println("----------------------------------------")
	//	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(uri, &param{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
