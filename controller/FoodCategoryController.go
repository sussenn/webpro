package controller

import (
	"github.com/gin-gonic/gin"
	"webpro/service"
	"webpro/tool"
)

type FoodCategoryController struct{}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)
}

//获取食品分类列表
func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {
	//这里可以不用 "&"?
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(context, "食品分类查询失败")
		return
	}
	for _, category := range categories {
		if category.ImageUrl != "" {
			//将fastFDS路径拼接
			//值传递 当前类下,值被改变了,能传递下去?
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
	}
	tool.Success(context, categories)
}
