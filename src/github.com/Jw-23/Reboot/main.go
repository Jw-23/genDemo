package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "残暴的欢愉必将以残暴的杀戮结束!")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Listen failed, error:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// 表明quit只能用来接受os.Signal类型的数据
	<-quit
	log.Println("Server ShutDown...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server ShutDown:", err)
	}
	log.Println("Server Has Stopped")

}
