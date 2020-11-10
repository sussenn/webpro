package tool

import (
	"github.com/gin-gonic/gin"
)

//图形验证码响应体
type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

//生成图形验证码
func GenerateCaptcha(ctx *gin.Context) {
	//TODO ...
}
