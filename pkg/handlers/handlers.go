package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {

	fmt.Fprintln(c.Writer, "Pong")
}

func LoggingMiddleware(DefaLogger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		DefaLogger.Info("finished", slog.Group("request",
			slog.String("method", c.Request.Method),
			slog.String("url", c.Request.URL.RawFragment)),
			slog.Int("status", http.StatusOK),
			slog.Duration("duration", time.Second),
		)
		c.Next()
	}

}
