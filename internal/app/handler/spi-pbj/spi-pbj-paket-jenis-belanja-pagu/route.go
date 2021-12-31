package spi_pbj_paket_jenis_belanja_pagu

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	//g.PUT("/:id", h.Upsert)
	g.GET("", h.GetSpiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID)
}
