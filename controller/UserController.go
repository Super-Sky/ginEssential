package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mwx563796/ginessential/common"
	"mwx563796/ginessential/model"
	"mwx563796/ginessential/util"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := c.PostForm("name")
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
	if len(name) ==0 {
		name = util.RandomString(10)
	}
	log.Println(name,telephone,password)
	//判断手机号是否存在
	if isTelephoneExist(db,telephone){
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已存在"})
		return
	}
	//创建用户
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":500,"msg":"加密错误"})
		return
	}
	newUser := model.User{
		Name:       name,
		Telephone:  telephone,
		Password:   string(hasedPassword),
	}
	db.Create(&newUser)
	//返回结果
	c.JSON(200,gin.H{
		"code":200,
		"msg":"注册成功",
	})
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
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":400,"msg":"密码错误"})
		return
	}
	//发放token
	token := "111"
	//返回结果
	c.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"注册成功",
	})
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