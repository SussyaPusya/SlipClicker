package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var JwtAuth = func(ctx *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		notAuth := []string{"/login", "/pong"} //Тут урлы где нету автр
		reqPath := c.Request.RequestURI        // путь запроса

		for _, value := range notAuth { // проверяем нужна ли авторизация

			if value == reqPath {
				c.Next()
				return
			}
		}
	}

	// response := make(map[string]interface{})
	tokenHeader := ctx.Request.Header.Get("Authorization")

	if tokenHeader == "" {
		ctx.JSON(http.StatusUnauthorize, "Missing Auth")

	}

}
