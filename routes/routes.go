package routes

import (
	"golangsensors/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/sensors", controllers.FetchAllSensors)
	e.POST("/sensors", controllers.StoreSensors)
	e.PUT("/sensors", controllers.UpdateSensors)
	e.DELETE("/sensors", controllers.DeleteSensors)

	return e
}
