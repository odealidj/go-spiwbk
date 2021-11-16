package thnang

import (
	"codeid-boiler/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save, middleware.Authentication)
	g.POST("/form", h.SaveFrom, middleware.Authentication)
	g.POST("/batch", h.SaveBatch, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
	g.GET("", h.GetAll, middleware.Authentication)

}
