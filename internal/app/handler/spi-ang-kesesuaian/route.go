package spi_ang_kesesuaian

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	g.GET("", h.GetBySpiSdmID)
}
