// file.go
// author:昌维 [github.com/cw1997]
// date:2017-05-11 01:29:51
package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadAllFile(filePth string) []byte {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

// 获取到的路径末尾不带斜杠
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
