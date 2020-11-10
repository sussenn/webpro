package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//自定义最终响应体
const (
	SUCCESS = 0
	FAILED  = 1
)

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"smg":  v,
	})
}
