package main

import (
	db "code-boiler/database"
	"code-boiler/database/migration"
	"code-boiler/internal/factory"
	"code-boiler/internal/http"
	"code-boiler/internal/middleware"
	"code-boiler/pkg/util/env"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title code-boiler
// @version 0.0.1
// @description This is a doc for code-boiler.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv("PORT")

	db.Init()
	migration.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
