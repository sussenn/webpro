package dao

import (
	"log"
	"webpro/model"
	"webpro/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

//Dao实例化
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

//查询所有食品分类列表
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	if err := fcd.Engine.Find(&categories); err != nil {
		log.Println("fcd.Engine.Find() err: ", err)
		return nil, err
	}
	return categories, nil
}
