package spi_ang_item

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
}
