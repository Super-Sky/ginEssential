package middleware

import (
	"github.com/gin-gonic/gin"
	"mwx563796/ginessential/common"
	"mwx563796/ginessential/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//验证
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		//验证通过  获取userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		//用户
		if user.ID == 0{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		//存在 将user的信息写入上下文
		ctx.Set("user",user)
		ctx.Next()

	}

}
