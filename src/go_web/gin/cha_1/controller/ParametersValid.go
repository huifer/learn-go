package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
)

type Page struct {
	Size int    `json:"size" binding:"required"`
	Name string `json:"name" validate:"regexp=^[a-zA-Z]*$"`
}

func ParamValid(c *gin.Context) {
	page := Page{}
	if err := c.ShouldBind(&page); err != nil {

		if err := validator.Validate(page); err != nil {
			fmt.Println("erererer")
		}

		c.JSON(http.StatusOK, gin.H{
			"data": page,
		})

	}

}
