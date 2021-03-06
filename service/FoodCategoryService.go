package service

import (
	"webpro/dao"
	"webpro/model"
)

type FoodCategoryService struct{}

func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	foodCategoryDao := dao.NewFoodCategoryDao()
	return foodCategoryDao.QueryCategories()
}
