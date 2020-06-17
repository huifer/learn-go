package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/src/go_web/gin/cha_1/utils"
	"net/http"
)

func Function(c *gin.Context) {
	config := utils.GlobalConfig
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取数据成功",
		"data": config,
	})
}
