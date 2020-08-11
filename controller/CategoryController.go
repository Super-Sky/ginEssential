package controller

import (
	"github.com/gin-gonic/gin"
	"mwx563796/ginessential/model"
	"mwx563796/ginessential/repository"
	"mwx563796/ginessential/response"
	"mwx563796/ginessential/vo"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {

	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory);err != nil {
		response.Fail(ctx,nil,"数据验证错误，分类名称必须填写")
	}

	updateCategory, err := c.Repository.SelectByName(requestCategory.Name)
	if updateCategory != nil {
		response.Fail(ctx,nil,"名称已存在")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "创建失败")
		return
	}
	response.Success(ctx,gin.H{"category":category},"添加数据成功")

}

func (c CategoryController) Update(ctx *gin.Context) {
	//获取body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory);err != nil {
		response.Fail(ctx,nil,"数据验证错误，分类名称必须填写")
	}
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)

	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}

	//更新分类
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx,nil,"更新失败,请重试")
		return
	}

	response.Success(ctx, gin.H{"category":category},"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)

	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category":category},"")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))
	_, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}
	err = c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"删除失败，请重试")
		return
	}
	response.Success(ctx, nil,"删除成功")
}

//func isNameExist(db *gorm.DB,name string) bool {
//	var category model.Category
//	db.Where("name = ?",name).First(&category)
//	if category.ID != 0{
//		return true
//	}else {
//		return false
//	}
//}

