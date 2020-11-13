package dao

import (
	"log"
	"webpro/model"
	"webpro/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{Orm: tool.DbEngine}
}

const DEFAULT_RANGE = 5

//条件查询商铺列表
func (sd *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	if "" == keyword {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1",
			longitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).
			Find(&shops)
		if err != nil {
			log.Println("shopDao.Engine.Where() err: ", err)
			return nil
		}
	} else {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, keyword).
			Find(&shops)
		if err != nil {
			log.Println("shopDao.Engine.Where() err: ", err)
			return nil
		}
	}
	return shops
}
