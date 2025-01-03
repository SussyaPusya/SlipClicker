package app

import "github.com/gin-gonic/gin"

var JwtAuth = func() gin.HandlerFunc {
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

}
