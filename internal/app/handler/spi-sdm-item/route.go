package spi_sdm_item

import (
	"codeid-boiler/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
	g.POST("/spisdmid", h.GetSpiSdmItemBySpiSdmID, middleware.Authentication)

}
