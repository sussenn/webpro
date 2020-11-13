package model

type FoodCategory struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Title       string `xorm:"varchar(20)" json:"title"`       //食品类别标题
	Description string `xorm:"varchar(30)" json:"description"` //食品描述
	ImageUrl    string `xorm:"varchar(255)" json:"image_url"`  //图片地址
	LinkUrl     string `xorm:"varchar(255)" json:"link_url"`   //食品类别链接
	IsInServing bool   `json:"is_in_serving"`                  //该类别是否在服务状态
}
