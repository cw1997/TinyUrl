// response.go
// author:昌维 [github.com/cw1997]
// date:2017-05-09 09:15:50

package web

import (
	"encoding/json"
)

type Json struct {
	Error    int         `json:"error"`
	ErrorMsg string      `json:"errormsg"`
	Data     interface{} `json:"data"`
}

func jsonReturn(raw Json) ([]byte, error) {
	return json.Marshal(raw)
}
