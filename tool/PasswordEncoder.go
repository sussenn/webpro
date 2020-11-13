package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

//密码加密 工具
func EncoderSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	//获得16进制数值,需转换为string
	s := hex.EncodeToString(sum)
	return string(s)
}

func Md5(data string) string {
	w := md5.New()
	io.WriteString(w, data)
	bydate := w.Sum(nil)
	result := fmt.Sprintf("%x", bydate)
	return result
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
