package main

import (
	"github.com/gin-gonic/gin"
	"mwx563796/ginessential/common"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	router := gin.Default()
	router = CollectRoute(router)
	panic(router.Run())
}