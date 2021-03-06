package spi_sdm

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	g.POST("/file", h.SaveWithFile)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("", h.Get)
	g.POST("/id", h.GetByID)

}
