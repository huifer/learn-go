package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" form:"username"`
	Age      int    `json:"age" form:"age"`
}

//Query `/user/find?username=张三&age=18`
func Query(c *gin.Context) {
	username := c.DefaultQuery("username", "默认值")
	age := c.Query("age")
	m := make(map[string]string)
	m["username"] = username
	m["age"] = age
	user := User{}
	c.ShouldBind(&user)
	fmt.Println("参数绑定到对象", user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    m,
	})
}

//Path Url获取姓名和年龄 /user/find/张三
func Path(c *gin.Context) {
	name := c.Param("name")
	byName := c.Params.ByName("name")
	get, b := c.Params.Get("name")
	if b {
		fmt.Println("获取到 name ", get)
	}
	m := make(map[string]string)
	m["c.Param"] = name
	m["c.Params.ByName"] = byName
	m["c.Params.Get"] = get

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    m,
	})

}

func BodyForm(c *gin.Context) {
	user := User{}
	err := c.ShouldBind(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    user,
		})
	}
}

func BodyJson(c *gin.Context) {
	user := User{}
	err := c.ShouldBind(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    user,
		})
	}
}
