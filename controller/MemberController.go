package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
	"webpro/model"
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
	//TODO 验证码校验
	//用户名+密码+验证码登录
	engine.POST("/api/login_pwd", mc.nameLogin)
	//用户头像上传
	engine.POST("/api/upload/avatar", mc.uploadAvatar)

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
		//用户信息保存到session
		ses, _ := json.Marshal(member)
		err = tool.SetSession(context, "user_"+string(member.Id), ses)
		if err != nil {
			log.Println("nameLogin() session保存失败. err: ", err)
			tool.Failed(context, "登录失败")
			return
		}
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

//用户名+密码+验证码登录
func (mc *MemberController) nameLogin(context *gin.Context) {
	//解析用户登录传参
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//TODO 验证码校验
	//登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		//用户信息保存到session
		ses, _ := json.Marshal(member)
		err = tool.SetSession(context, "user_"+string(member.Id), ses)
		if err != nil {
			log.Println("nameLogin() session保存失败. err: ", err)
			tool.Failed(context, "登录失败")
			return
		}
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "登录失败")
}

//用户头像上传
func (mc *MemberController) uploadAvatar(context *gin.Context) {
	//解析传参: file,user_id
	userId := context.PostForm("user_id")
	file, err := context.FormFile("avatar")
	if err != nil || "" == userId {
		tool.Failed(context, "参数解析失败")
		return
	}

	//先判断用户是否登录
	ses := tool.GetSession(context, "user_"+userId)
	if nil == ses {
		tool.Failed(context, "非法参数")
		return
	}
	var member model.Member
	//类型强转
	json.Unmarshal(ses.([]byte), &member)

	//file保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	//将文件上传到FastDFS
	fileId := tool.UploadFile(fileName)
	if fileId != "" {
		//删除本地uploadfile文件夹下文件
		os.Remove(fileName)
		//将文件路径入库
		membersService := service.MemberService{}
		path := membersService.UploadAvatar(member.Id, fileName[1:])
		if path != "" {
			tool.Success(context, tool.FileServerAddr()+"/"+path)
			return
		}
	}
}
