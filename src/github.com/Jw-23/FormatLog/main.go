package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	gin.DisableBindValidation()
	gin.SetMode(gin.ReleaseMode)
	f, err := os.Create("gin.log")
	if err != nil {
		log.Println("创建文件失败")
	}
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.String(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.String(200, "bar")
	})
	err = r.Run(":8000")
	if err != nil {
		return
	}
}
