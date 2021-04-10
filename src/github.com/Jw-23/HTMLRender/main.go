package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("github.com/Jw-23/HTMLRender/templates/*")
	// 上面是导入这个文件夹，也可以导入一个文件, router.LoadHTMLFiles("......", "......")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "王尼玛妈死了",
		})
	})
	router.Run(":8000")
}
