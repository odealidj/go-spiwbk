package sample

import (
	"code-boiler/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middlewares.Authentication)
	g.GET("/:id", h.GetByID, middlewares.Authentication)
	g.POST("", h.Store, middlewares.Authentication, middlewares.Transaction)
	g.PATCH("/:id", h.Update, middlewares.Authentication, middlewares.Transaction)
	g.DELETE("/:id", h.Delete, middlewares.Authentication, middlewares.Transaction)
}
