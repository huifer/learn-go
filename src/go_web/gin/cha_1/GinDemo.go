package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/src/go_web/gin/cha_1/controller"
	"learn-go/src/go_web/gin/cha_1/utils"
)

func main() {
	err := utils.LoadConfig("C:\\Users\\admin\\GolandProjects\\learn-go\\src\\go_web\\gin\\cha_1\\config.json")
	if err != nil {
		fmt.Println("读取配置失败")
		panic(err)

	}

	port := ":" + utils.GlobalConfig.ServerConfig.Port

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/go", controller.Function)

	r.Run(port)
}
