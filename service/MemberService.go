package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
	"webpro/dao"
	"webpro/model"
	"webpro/param"
	"webpro/tool"
)

type MemberService struct{}

//注入Dao	[不行,错误的写法]
//var memberDao = dao.MemberDao{Orm: tool.DbEngine}

//调用阿里云发送验证码
func (ms *MemberService) SendCode(phone string) bool {
	//生成随机验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//调用阿里云短信
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Println("dysmsapi.NewClientWithAccessKey() err: ", err)
		return false
	}
	//短信模板参数配置
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	//设置验证码 {"code":"1234"}
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)
	//发送短信
	response, err := client.SendSms(request)
	if err != nil {
		log.Println("client.SendSms() err: ", err)
		return false
	}
	//接收发送结果
	if response.Code == "OK" {
		//将验证码入库
		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      response.BizId,
			CreateTime: time.Now().Unix(),
		}
		//注入Dao
		memberDao := dao.MemberDao{Orm: tool.DbEngine}
		res := memberDao.AddCode(smsCode)
		return res > 0
	}
	return false
}

//手机号+验证码的登录
func (ms *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {
	//注入Dao
	memberDao := dao.MemberDao{Orm: tool.DbEngine}
	//验证手机号和验证码
	sms := memberDao.ValidateSmsCode(loginParam.Phone, loginParam.Code)
	if sms.Id == 0 {
		return nil
	}
	//根据手机号查询用户信息
	member := memberDao.QueryByPhone(loginParam.Phone)
	if member.Id != 0 {
		return member
	}
	//如果查询不到,则新增用户数据
	user := model.Member{}
	user.UserName = loginParam.Phone
	user.Mobile = loginParam.Phone
	user.RegisterTime = time.Now().Unix()
	//insert会返回自增id
	user.Id = memberDao.InsertMember(user)
	return &user
}

//用户名+密码 登录
func (ms *MemberService) Login(name string, password string) *model.Member {
	//根据用户名+密码查询用户信息
	md := dao.MemberDao{Orm: tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}
	//查不到则新增用户
	user := model.Member{}
	user.UserName = name
	//密码加密
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	//插入数据库,返回自增id
	id := md.InsertMember(user)
	user.Id = id
	log.Println("未查询到用户,进行用户注册,用户name:", user.UserName)
	return &user
}

//头像上传 更新用户头像信息
func (ms *MemberService) UploadAvatar(userId int64, fileName string) string {
	memberDao := dao.MemberDao{Orm: tool.DbEngine}
	result := memberDao.UploadAvatar(userId, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}
