package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logrus.Info("Beginning database transaction when its detected")
		c.Response().Before(func() {
			if _, ok := c.Get("db_transaction").(*gorm.DB); ok {
				tx, ok := c.Get("db_transaction").(*gorm.DB)
				if !ok {
					logrus.Error("db transaction not found in context")
				}

				statusInList := func(status int, statusList []int) bool {
					for _, i := range statusList {
						if i == status {
							return true
						}
					}
					return false
				}

				if statusInList(c.Response().Status, []int{http.StatusOK, http.StatusAccepted, http.StatusAlreadyReported, http.StatusCreated}) {
					logrus.Info("Commiting database transaction")
					if err := tx.Commit().Error; err != nil {
						logrus.Error("trx commit error: ", err)
					}
				} else {
					logrus.Info("rolling back transaction due to status code: ", c.Response().Status)
					tx.Rollback()
				}
			} else {
				logrus.Error("db transaction not found in context")
			}

		})

		return next(c)
	}
}
