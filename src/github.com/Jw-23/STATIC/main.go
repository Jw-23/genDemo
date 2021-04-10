package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("github.com/Jw-23/HTTP2Server/dist/index.html")
	r.StaticFS("/js", http.Dir("github.com/Jw-23/HTTP2Server/dist/js"))
	r.StaticFS("/css", http.Dir("github.com/Jw-23/HTTP2Server/dist/css"))
	r.StaticFS("/img", http.Dir("github.com/Jw-23/HTTP2Server/dist/img"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":8000")
}
