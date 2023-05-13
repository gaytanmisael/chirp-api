package main

import (
	"chirp-api/app/middleware"
	"chirp-api/app/router"
	"chirp-api/internal/bootstrap"
	"chirp-api/internal/bootstrap/database"
	"chirp-api/utils/config"

	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(router.NewRouter),

		fx.Invoke(bootstrap.Start),

		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
