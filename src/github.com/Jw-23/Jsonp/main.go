package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(200, data)
	})
	if err := r.Run(":8000"); err != nil {
		fmt.Println("run failed: ", err)
	}
}
