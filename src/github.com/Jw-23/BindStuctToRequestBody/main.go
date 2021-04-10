package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {

}

/*func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// ShouldBind 使用了c.Request.Body,不可重用
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(200, `the body should be formA`)
		// 此时的c.Request.Body已经是EOF(end of file),这里会报错
		// 要想多次绑定，需要使用ShouldBindBodyWith
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(200, `the body should be formB`)
	} else {
		c.String(200, `Bind nothing!`)
	}

}*/

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// ShouldBind 使用了c.Request.Body,不可重用
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(200, `the body should be formA JSON`)
		// 此时的c.Request.Body已经是EOF(end of file),这里会报错
		// 要想多次绑定，需要使用ShouldBindBodyWith
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(200, `the body should be formB JSON`)
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK,`the body should be formB XML`)
	} else {
		c.String(200, `Bind nothing!`)
	}
}
