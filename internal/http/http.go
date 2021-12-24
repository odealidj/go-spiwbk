package http

import (
	docs "codeid-boiler/docs"
	"codeid-boiler/internal/app/handler/auth"
	"codeid-boiler/internal/app/handler/bulan"
	"codeid-boiler/internal/app/handler/group-package-value"
	jenis_belanja_pagu "codeid-boiler/internal/app/handler/jenis-belanja-pagu"
	"codeid-boiler/internal/app/handler/jenis-certificate"
	jenis_kesesuaian "codeid-boiler/internal/app/handler/jenis-kesesuaian"
	jenis_pengendali "codeid-boiler/internal/app/handler/jenis-pengendali"
	jenis_rekapitulasi "codeid-boiler/internal/app/handler/jenis-rekapitulasi"
	"codeid-boiler/internal/app/handler/jenis-sdm"
	method_apbj "codeid-boiler/internal/app/handler/method-apbj"
	"codeid-boiler/internal/app/handler/pegawai"
	"codeid-boiler/internal/app/handler/rkakl"
	"codeid-boiler/internal/app/handler/satker"
	spi_ang "codeid-boiler/internal/app/handler/spi-ang"
	spi_ang_item "codeid-boiler/internal/app/handler/spi-ang-item"
	spi_ang_kesesuaian "codeid-boiler/internal/app/handler/spi-ang-kesesuaian"
	spi_bmn "codeid-boiler/internal/app/handler/spi-bmn"
	spi_pbj_paket_jenis_belanja_pagu "codeid-boiler/internal/app/handler/spi-pbj-paket-jenis-belanja-pagu"
	spi_pbj_rekapitulasi "codeid-boiler/internal/app/handler/spi-pbj-rekapitulasi"
	"codeid-boiler/internal/app/handler/spi-sdm"
	"codeid-boiler/internal/app/handler/spi-sdm-item"
	"codeid-boiler/internal/app/handler/thnang"
	wbk_program_ranker "codeid-boiler/internal/app/handler/wbk-program-ranker"
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

	//static
	e.Static("/upload", "./upload")

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
	rkakl.NewHandler(f).Route(g.Group("/rkakl"))
	spi_ang.NewHandler(f).Route(g.Group("/spi-ang"))
	spi_ang_item.NewHandler(f).Route(g.Group("/spi-ang-item"))
	spi_ang_kesesuaian.NewHandler(f).Route(g.Group("/spi-ang-kesesuaian"))
	jenis_kesesuaian.NewHandler(f).Route(g.Group("/jenis-kesesuaian"))
	jenis_pengendali.NewHandler(f).Route(g.Group("/jenis-pengendali"))
	jenis_rekapitulasi.NewHandler(f).Route(g.Group("/jenis-rekapitulasi"))
	spi_pbj_rekapitulasi.NewHandler(f).Route(g.Group("/spi-pbj-rekapitulasi"))
	bulan.NewHandler(f).Route(g.Group("/bulan"))
	group_package_value.NewHandler(f).Route(g.Group("/group-package-value"))
	jenis_belanja_pagu.NewHandler(f).Route(g.Group("/jenis-belanja-pagu"))
	method_apbj.NewHandler(f).Route(g.Group("/method-apbj"))
	spi_pbj_paket_jenis_belanja_pagu.NewHandler(f).Route(g.Group("/spi-pbj-paket-jenis-belanja-pagu"))
	spi_bmn.NewHandler(f).Route(g.Group("/spi-bmn"))
	wbk_program_ranker.NewHandler(f).Route(g.Group("/wbk-program-ranker"))
}
