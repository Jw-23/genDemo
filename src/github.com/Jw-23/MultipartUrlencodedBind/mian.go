package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	PassWord string `form:"password" binding:"required"`
	Message  string `form:"message"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {

		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "wnm" && form.PassWord == "sb2223" {
				c.JSONP(200, gin.H{"status": "王凤贺妈死了，全家都死了"})
			} else {
				c.JSONP(201, gin.H{"status": "unauthorized"})
			}
			fmt.Println("收到数据")
			fmt.Println(form.Message)
		}
	})
	router.Run(":8000")
}
