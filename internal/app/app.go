package app

import (
	"log/slog"
	"os"

	"github.com/SussyaPusya/SlipClicker/pkg/handlers"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Path string
}

func ConfigWhitEnv() *Config {
	cfg := new(Config)

	cfg.Path = os.Getenv("PORT")
	if cfg.Path == "" {
		cfg.Path = "8888"

	}
	return cfg
}

type App struct {
	config *Config
}

func New() *App {
	return &App{
		config: ConfigWhitEnv(),
	}
}

func (a *App) Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	logger := slog.Default()

	router.Static("/assets", "./pkg/assets")
	router.StaticFile("/assets", "./pkg/assets")

	router.GET("/ping", handlers.Pong)

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("pkg/html/index.html")

	})
	router.Use(handlers.LoggingMiddleware(logger))

}
