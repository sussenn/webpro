package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello/:id", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 20001,
		"msg":  "success",
		"data": id,
	})
}
