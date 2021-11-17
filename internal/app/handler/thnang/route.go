package thnang

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	g.POST("/id", h.Get)
	g.POST("/form", h.SaveFrom)
	g.POST("/batch", h.SaveBatch)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("", h.GetAll)

}
