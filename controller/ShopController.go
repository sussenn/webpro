package controller

import (
	"github.com/gin-gonic/gin"
	"webpro/service"
	"webpro/tool"
)

type ShopController struct{}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.getShopList)
	engine.GET("/api/search_shop", sc.SearchShops)
}

//获取商铺列表
func (sc *ShopController) getShopList(context *gin.Context) {
	//接收传参经,纬度
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	if "" == longitude || "undefined" == longitude || "" == latitude || "undefined" == latitude {
		//先写死
		longitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) > 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "暂无商铺信息")
}

//商铺名关键字搜索商铺列表
func (sc *ShopController) SearchShops(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.Query("keyword")
	if "" == keyword {
		tool.Failed(context, "请输入商铺名称")
		return
	}
	shopService := service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keyword)
	if len(shops) > 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "暂无商铺信息")
}
