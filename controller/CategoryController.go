package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mwx563796/ginessential/common"
	"mwx563796/ginessential/model"
	"mwx563796/ginessential/response"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB:db}
}

func (c CategoryController) Create(ctx *gin.Context) {

	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx,nil,"数据验证错误，分类名称必须填写")
		return
	}
	if isNameExist(c.DB,requestCategory.Name) {
		response.Fail(ctx,nil,"名称已存在")
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx,gin.H{"category":requestCategory},"添加数据成功")

}

func (c CategoryController) Update(ctx *gin.Context) {
	//获取body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx,nil,"数据验证错误，分类名称必须填写")
		return
	}
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	var updeCategory model.Category
	if c.DB.First(&updeCategory,categoryId).RecordNotFound() {
		response.Fail(ctx,nil,"分类不存在")
		return
	}

	//更新分类
	c.DB.Model(&updeCategory).Update("name",requestCategory.Name)

	response.Success(ctx, gin.H{"category":updeCategory},"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category,categoryId).RecordNotFound() {
		response.Fail(ctx,nil,"分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category":category},"")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	if err:= c.DB.Delete(model.Category{},categoryId).Error;err != nil {
		response.Fail(ctx,nil,"删除失败，请重试")
		return
	}
	response.Success(ctx, nil,"删除成功")
}

func isNameExist(db *gorm.DB,name string) bool {
	var category model.Category
	db.Where("name = ?",name).First(&category)
	if category.ID != 0{
		return true
	}else {
		return false
	}
}

