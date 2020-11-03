package controllers

import (
	"github.com/Kento75/unittesting-in-golang/golang-testing-echo/src/api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCountry(c echo.Context) {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, country)
}
