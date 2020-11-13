package service

import (
	"log"
	"strconv"
	"webpro/dao"
	"webpro/model"
)

type ShopService struct{}

//根据经纬度查询商铺列表
func (ss *ShopService) ShopList(lon, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(lon, 10)
	if err != nil {
		log.Println("strconv.ParseFloat() err: ", err)
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		log.Println("strconv.ParseFloat() err: ", err)
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude, "")
}

//商铺名关键字搜索商铺列表
func (ss *ShopService) SearchShops(lon string, lat string, keyword string) []model.Shop {
	longitude, err := strconv.ParseFloat(lon, 10)
	if err != nil {
		log.Println("strconv.ParseFloat() err: ", err)
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		log.Println("strconv.ParseFloat() err: ", err)
		return nil
	}
	sd := dao.NewShopDao()
	return sd.QueryShops(longitude, latitude, keyword)
}
