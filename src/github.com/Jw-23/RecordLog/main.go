package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableBindValidation()

	// 记录到文件

	f, _ := os.Create("RecordLog/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	if err := r.Run(":8000"); err != nil {
		log.Println("server run failed,", err)
	}
}
