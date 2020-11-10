package dao

import (
	"log"
	"webpro/model"
	"webpro/tool"
)

type MemberDao struct {
	*tool.Orm
}

//验证码信息入库
func (md *MemberDao) AddCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Println("AddCode() err: ", err)
	}
	return result
}

//判断手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		log.Println("ValidateSmsCode() err: ", err)
	}
	return &sms
}

//根据手机号查询用户信息
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		log.Println("QueryByPhone() err: ", err)
	}
	return &member
}

//用户数据入库
func (md *MemberDao) InsertMember(member model.Member) int64 {
	//返回自增id
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Println("InsertMember() err: ", err)
		return 0
	}
	return result
}
