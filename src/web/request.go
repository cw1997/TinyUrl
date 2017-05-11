// request.go
// author:昌维 [github.com/cw1997]
// date:2017-05-07 19:13:06
package web

import (
	//	"fmt"
	"net/http"
	"net/url"
	//	"strconv"
	"log"
	"strings"

	"util"
)

func getQuery(r *http.Request, parameter string) (query string) {
	path := r.URL.String() //获得访问的路径
	//	fmt.Println(path)
	urlStruct, err := url.Parse(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	queryMap := urlStruct.Query()
	//	fmt.Println(queryMap)
	if err != nil {
		log.Println(err)
		return ""
	}
	//	fmt.Println(queryMap)
	if len(queryMap) > 0 {
		param := queryMap[parameter]
		if param != nil {
			return util.Slice2string(param[0:1])
		}
		return ""
	}
	return ""
}

func getIP(r *http.Request) (ip string) {
	return r.RemoteAddr
}

func getPath(r *http.Request) (path string) {
	url := r.URL.Path[1:]
	queryIndex := strings.Index(url, "?")
	//	fmt.Println(url, queryIndex)
	if queryIndex == -1 {
		return url
	} else {
		return url[0 : queryIndex+1]
	}
}
