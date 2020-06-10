package goIris

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// first demo
func RunWithIris() {
	application := iris.New()

	application.Get("/hello", func(context context.Context) {
		context.HTML("<h1>hello iris</h1>")
	})

	// 模板引擎设置
	html := iris.HTML("./src/goIris/", ".html")
	application.RegisterView(html)

	application.Get("/index", func(c context.Context) {
		c.ViewData("title", "标题")
		c.ViewData("name", "huifer")

		c.View("index.html")

	})

	application.Run(iris.Addr(":9090"))
}
