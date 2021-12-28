package wbk_program_ranker

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Save)
	//g.PUT("/:id", h.Upsert)
	g.GET("", h.Get)
	g.GET("/satker-nilai", h.GetSatkerNilaiByThnAngID)
	g.GET("/program-ranker-nilai", h.GetByThnAngIDAndSatkerID)
}
