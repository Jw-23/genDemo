package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.Set("ContentType", "application/json")
		c.JSONP(200, gin.H{
			"status": "posted",
			message:  "message",
			nick:     "nick",
		})
	})
	r.Run(":8000")
}
