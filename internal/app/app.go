package app

import "os"

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

func (a *App) Run() error {
	//логика сервера
}
