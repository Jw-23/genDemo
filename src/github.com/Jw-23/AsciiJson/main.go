package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AsciiJson 可以转换非ASCII字符
func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
	var err = r.Run(":8000")
	if err != nil {
		fmt.Println("启动服务器失败")
	}
}
