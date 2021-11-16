package pegawai

import (
	"codeid-boiler/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save, middleware.Authentication)
	g.GET("", h.Get, middleware.Authentication)

}
