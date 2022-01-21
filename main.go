package main

import (
	db "codeid-boiler/database"
	"codeid-boiler/internal/factory"
	"codeid-boiler/internal/http"
	"codeid-boiler/internal/middleware"
	"context"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	netHTTP "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
func init() {
	//err := godotenv.Load(".env.development")
	//err := godotenv.Load(".env.local")
	err := godotenv.Load(".env.prod")
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
	//ENV := "LOCAL"
	logrus.Info("Choosen environment " + ENV)

	var PORT = os.Getenv("PORT")

	db.Init()
	defer db.Close()
	//migration.Init()
	//elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	debugpprof := strings.ToLower(strings.TrimSpace(os.Getenv("DEBUG_PPROF")))
	if debugpprof == "active" {
		// automatically add routers for net/http/pprof
		// e.g. /debug/pprof, /debug/pprof/heap, etc.
		echopprof.Wrap(e)
		logrus.Info("ECHOPPROF registered")
	}
	logrus.Info("DEBUG_PPROF ", debugpprof == "active", " "+debugpprof)

	// Start server
	go func() {
		if err := e.Start(":" + PORT); err != nil && err != netHTTP.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	/*
		db.Init()
		//migration.Init()

		e := echo.New()
		middleware.Init(e)

		f := factory.NewFactory()
		http.Init(e, f)

		e.Logger.Fatal(e.Start(":" + PORT))

	*/
}
