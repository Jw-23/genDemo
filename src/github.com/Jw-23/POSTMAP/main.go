package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		log.Printf("ids : %v; names: %v\n", ids, names)
	})
	err := r.Run(":8000")
	if err != nil {
		log.Println("server run failed", err)
		return
	}
}
