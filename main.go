package main

import (
	"gin-demo/common"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	common.InitConfig("config/config.yaml")
	common.InitDB()

	//创建DB实例进行数据库操作
	db := common.GetDB()
	//延迟关闭数据库
	defer db.Close()
	/*
		r.GET("/ping",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"code":200,
				"msg":"pong",
			})
		})
		r.POST("/user/register",controller.Register)
		r.POST("/user/login",controller.Login)*/
	r = router.InitRouter(r)
	r.Run(":8888")
}
