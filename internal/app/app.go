package app

import (
	"fmt"
	"os"

	"github.com/SussyaPusya/SlipClicker/SlipClicker/pkg/handlers"
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
	router := gin.Default()


	router.Static("/assets", "./assets")
	router.StaticFile("/assets", "./assets")

	router.GET("/ping", handlers.Pong)


	ro

	

	err := router.Run(":" + a.config.Path)
	if err != nil {
		fmt.Println(err)

	}

}
