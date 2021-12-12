package main

import (
	db "codeid-boiler/database"
	"codeid-boiler/internal/factory"
	"codeid-boiler/internal/http"
	"codeid-boiler/internal/middleware"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/labstack/echo/v4"
)

/*
func init() {
	//env := env.NewEnv()
	//env.Load(ENV)

	err := godotenv.Load(".env.development")
	if err != nil {
		panic("Failed to load .env file, Make sure .env is exists")

	}
}
*/

// @title codeid-boiler
// @version 0.0.1
// @description This is a doc for codeid-boiler.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {

	ENV := os.Getenv("ENV")
	logrus.Info("Choosen environment " + ENV)

	var PORT = os.Getenv("PORT")

	db.Init()
	//migration.Init()
	//elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
