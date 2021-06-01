package main

import (
	"code-boiler/database"
	"code-boiler/internal/route"
	"code-boiler/pkg/util"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	if ENV == "" {
		ENV = "DEV"
	}
	env := util.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environments " + ENV)
}

// @title code-boiler
// @version 0.0.1
// @description This is a doc for code-boiler.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var (
		APP  = os.Getenv("APP")
		ENV  = os.Getenv("ENV")
		PORT = os.Getenv("PORT")
		NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

	connections, err := database.Connect()
	if err != nil {
		logrus.Error(err)
		panic("Failed to connect database")
	}
	for _, v := range connections {
		db, err := v.DB()
		if err != nil {
			logrus.Error(err)
		}
		defer db.Close()
	}
	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | &{latency_human} | ${remote_ip} | ${method} | ${uri}", NAME),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	route.Init(e.Group(""))

	e.Logger.Fatal(e.Start(":" + PORT))
}
