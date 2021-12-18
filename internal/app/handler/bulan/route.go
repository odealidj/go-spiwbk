package bulan

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	//g.POST("", h.Save)
	//g.PUT("/:id", h.Update)
	//g.DELETE("/:id", h.Delete)
	//g.POST("/spisdmid", h.GetSpiSdmItemBySpiSdmID)
	g.GET("", h.Get)

}
