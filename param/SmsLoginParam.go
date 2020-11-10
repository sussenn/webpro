package param

//手机号+验证码 登录的参数传递
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
