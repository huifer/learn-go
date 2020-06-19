package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"learn-go/src/go_web/gin/cha_1/controller"
	"learn-go/src/go_web/gin/cha_1/utils"
	"learn-go/src/go_web/gin/cha_1/validators"
)

func valid() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sizeValid", validators.SizeValid)
	}

}

func main() {
	err := utils.LoadConfig("C:\\Users\\admin\\GolandProjects\\learn-go\\src\\go_web\\gin\\cha_1\\config.json")
	if err != nil {
		fmt.Println("读取配置失败")
		panic(err)

	}

	port := ":" + utils.GlobalConfig.ServerConfig.Port

	r := gin.Default()
	valid()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/go", controller.Function)

	r.GET("/user/find", controller.Query)
	r.GET("/user/find/:name", controller.Path)

	r.POST("/user/form", controller.BodyForm)
	r.POST("/user/json", controller.BodyJson)
	r.POST("/param/valid", controller.ParamValid)

	r.Run(port)
}
