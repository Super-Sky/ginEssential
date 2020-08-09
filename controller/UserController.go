package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mwx563796/ginessential/common"
	"mwx563796/ginessential/dto"
	"mwx563796/ginessential/model"
	"mwx563796/ginessential/response"
	"mwx563796/ginessential/util"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	var requestUser = model.User{}
	//json.NewDecoder(c.Request.Body).Decode(&requestUser)
	c.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		//fmt.Println("telephone",telephone)
		//fmt.Println(len(telephone))
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	if len(password) <6{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不少于6位")
		return
	}
	if len(name) ==0 {
		name = util.RandomString(10)
	}
	log.Println(name,telephone,password)
	//判断手机号是否存在
	if isTelephoneExist(db,telephone){
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已存在")
		return
	}
	//创建用户
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}
	newUser := model.User{
		Name:       name,
		Telephone:  telephone,
		Password:   string(hasedPassword),
	}
	db.Create(&newUser)
	//发放token
	token,err := common.ReleaseToken(newUser)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error: %v",err)
		return
	}
	//返回结果
	response.Success(c,gin.H{"token":token},"注册成功")
}

func Loginer(c *gin.Context)  {
	db := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		fmt.Println("telephone",telephone)
		fmt.Println(len(telephone))
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号必须为11位"})
		return
	}
	if len(password) <6{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不少于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		return
	}
	//发放token
	token,err := common.ReleaseToken(user)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error: %v",err)
		return
	}

	//返回结果
	response.Success(c,gin.H{"token":token},"登陆成功")
}

func Info(ctx *gin.Context)  {
	user,_ := ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0{
		return true
	}else {
		return false
	}
}