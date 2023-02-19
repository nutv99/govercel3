package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/tbxark/g4vercel"
)

func check2(c *gin.Context) {

	server := New()
	server.GET("/users", func(context *Context) {
		context.JSON(http.StatusFound, H{
			"message": "hello go from vercel from  folder USERS !!!!",
		})
	})

}

func ReadAll(c *gin.Context) {

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
		"mesage": "Read All user Data  Success",
	})
}
