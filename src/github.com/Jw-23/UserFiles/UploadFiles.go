package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("github.com/Jw-23/UserFiles/index.html")

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		name := c.PostForm("name")
		email := c.PostForm("email")
		if err != nil {
			log.Println("create multiform failed,error:", err)
			return
		}
		files := form.File["userFiles"]
		for _, file := range files {
			err := c.SaveUploadedFile(file, filepath.Join("github.com/Jw-23/UserFiles", file.Filename))
			if err != nil {
				log.Println("save filed failed:", err)
				c.String(http.StatusOK, "save failed ,please try again!")
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded successfully with fields name=%s and email=%s.", len(files), name, email))
	})
	r.GET("/ker", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	if err := r.Run(":8000"); err != nil {
		log.Println("serve run failed", err)
	}
}
