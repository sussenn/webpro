package controller

import (
	"github.com/gin-gonic/gin"
	"webpro/param"
	"webpro/service"
	"webpro/tool"
)

type MemberController struct{}

//路由匹配
func (mc *MemberController) Router(engine *gin.Engine) {
	//发送验证码
	engine.GET("/api/sendCode", mc.sendSmsCode)
	//手机号+验证码登录
	engine.POST("/api/login_sms", mc.smsLogin)
	//生成图形验证码
	engine.GET("/api/captcha", mc.captcha)
}

//发送短信验证码
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//接收 ?phone=18779816466
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "参数解析失败",
		})
		return
	}
	//注入Service
	ms := service.MemberService{}
	//发送验证码短信
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		return
	}
	tool.Failed(context, "发送失败")
}

//手机号+验证码的登录
func (mc *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	//解析请求体的json参数
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//注入Service
	us := service.MemberService{}
	//调用service方法
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登录失败")
}

//生成图形验证码 并直接返回客户端
func (mc *MemberController) captcha(context *gin.Context) {
	//TODO 暂未完成
	tool.GenerateCaptcha(context)
}
