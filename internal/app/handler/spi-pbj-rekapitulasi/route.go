package spi_pbj_rekapitulasi

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Upsert)
	g.PUT("/:id", h.Upsert)
	g.GET("", h.GetSpiPbjRekapitulasiByID)
}
