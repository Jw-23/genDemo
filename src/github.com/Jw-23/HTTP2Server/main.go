package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

var html = template.Must(template.New("https").Parse(`
<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="icon" href="/favicon.ico"><title>blog</title><link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900"><link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@mdi/font@latest/css/materialdesignicons.min.css"><link href="/css/chunk-5204670c.d690010b.css" rel="prefetch"><link href="/css/chunk-cbb6f52a.dad7a9d3.css" rel="prefetch"><link href="/css/chunk-db9c1e70.292aa765.css" rel="prefetch"><link href="/js/about.64b836fb.js" rel="prefetch"><link href="/js/chunk-2d225f0b.ba31fe82.js" rel="prefetch"><link href="/js/chunk-5204670c.807f6212.js" rel="prefetch"><link href="/js/chunk-72f072e4.d76ea550.js" rel="prefetch"><link href="/js/chunk-cbb6f52a.d0ec9adf.js" rel="prefetch"><link href="/js/chunk-db9c1e70.70094fb4.js" rel="prefetch"><link href="/css/chunk-vendors.988cb658.css" rel="preload" as="style"><link href="/js/app.0de51653.js" rel="preload" as="script"><link href="/js/chunk-vendors.7debc800.js" rel="preload" as="script"><link href="/css/chunk-vendors.988cb658.css" rel="stylesheet"></head><body><noscript><strong>We're sorry but blog doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id="app"></div><script src="/js/chunk-vendors.7debc800.js"></script><script src="/js/app.0de51653.js"></script></body></html>
`))

func main() {
	r := gin.Default()
	r.Static("/css", "./dist/css")
	r.Static("/img", "./dist/img")
	r.Static("/js", "./dist/js")
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/js/app.0de51653.js", nil); err != nil {
				fmt.Println("push failed,err:", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	r.Run(":8000")

}
