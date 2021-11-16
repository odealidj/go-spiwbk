package spi_sdm

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	g.GET("", h.Get)

}
