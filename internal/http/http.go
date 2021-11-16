package http

import (
	docs "codeid-boiler/docs"
	"codeid-boiler/internal/app/handler/auth"
	"codeid-boiler/internal/app/handler/jenis-certificate"
	"codeid-boiler/internal/app/handler/jenis-sdm"
	"codeid-boiler/internal/app/handler/pegawai"
	"codeid-boiler/internal/app/handler/satker"
	"codeid-boiler/internal/app/handler/spi-sdm"
	"codeid-boiler/internal/app/handler/spi-sdm-item"
	"codeid-boiler/internal/app/handler/thnang"
	"codeid-boiler/internal/factory"
	"codeid-boiler/pkg/constant"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
		HOST    = os.Getenv("HOST")
		SCHEME  = os.Getenv("SCHEME")
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	g := e.Group(constant.ROUTE_GROUP_V1)
	auth.NewHandler(f).Route(g.Group("/auth"))
	thnang.NewHandler(f).Route(g.Group("/thnang"))
	satker.NewHandler(f).Route(g.Group("/satker"))
	spi_sdm.NewHandler(f).Route(g.Group("/spi-sdm"))
	jenis_sdm.NewHandler(f).Route(g.Group("/jenis-sdm"))
	jenis_certificate.NewHandler(f).Route(g.Group("/jenis-certificate"))
	pegawai.NewHandler(f).Route(g.Group("/pegawai"))
	spi_sdm_item.NewHandler(f).Route(g.Group("/spi-sdm_item"))
}
