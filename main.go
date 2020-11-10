package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"webpro/controller"
	"webpro/tool"
)

//go get -u github.com/aliyun/alibaba-cloud-sdk-go		阿里云...
//go get -u github.com/go-xorm/xorm						持久层框架X-ORM
//go get -u github.com/mojocn/base64Captcha 			图形验证码
//go get -u github.com/go-redis/redis					redis
func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		//log.Fatal(err)
		panic(err.Error())
	}
	//加载数据库orm引擎
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Println("tool.OrmEngine() err: ", err)
		return
	}
	//初始化redis
	tool.InitRedisStore()

	app := gin.Default()
	//设置允许跨域
	app.Use(Cors())
	//注入Controller
	registerRouter(app)
	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置 加载controller
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}

//设置允许跨域访问
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		//获取前端请求头 跨域访问标识
		origin := context.Request.Header.Get("Origin")
		//获取所有的请求头参数
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if "" != headerStr {
			//拼接允许跨域的请求头
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		//https://blog.csdn.net/LDY1016/article/details/106691752?utm_medium=distribute.pc_relevant.none-task-blog-OPENSEARCH-2.compare&depth_1-utm_source=distribute.pc_relevant.none-task-blog-OPENSEARCH-2.compare
		if "" != origin {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") //允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma, token, openid, opentoken")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false") //设置为true,允许ajax异步请求带cookie信息
			context.Set("content-type", "application/json")             //返回json类型
		}
		if "OPTIONS" == method {
			context.JSON(http.StatusOK, "Options 请求!")
		}
		//继续执行后续代码,处理业务
		context.Next()
	}
}
