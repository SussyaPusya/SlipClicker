package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {

	fmt.Fprintln(c.Writer, "Pong")
}
