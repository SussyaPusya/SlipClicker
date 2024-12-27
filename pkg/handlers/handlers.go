package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {

	fmt.Fprintln(c.Writer, "Pong")
}

func ClickerHandler(c *gin.Context) {
	c.File("pkg/html/index.html")
}
