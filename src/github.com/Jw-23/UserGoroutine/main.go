package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//当在中间件或者handler中启动新的Goroutine时，不能使用原始的上下文，必须使用只读副本
func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {

			time.Sleep(5 * time.Second)

			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {

		time.Sleep(5 * time.Second)

		log.Println("Done ! in path" + c.Request.URL.Path)
	})
	if err := r.Run(":8000"); err != nil {
		log.Println("server run failed,", err)
	}
}
