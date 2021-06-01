package route

import (
	_ "code-boiler/docs"
	"code-boiler/internal/app/auth"
	"code-boiler/internal/app/sample"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(g *echo.Group) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
	)

	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	g.GET("/swagger/*", echoSwagger.WrapHandler)

	//routes
	auth.NewHandler().Route(g.Group("/auth"))
	sample.NewHandler().Route(g.Group("/samples"))
}
