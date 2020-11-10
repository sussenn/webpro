package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct{}

//json参数 解析工具
func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
