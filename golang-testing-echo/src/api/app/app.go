package app

import (
	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func StartApp() {
	mapUrls()
	if err := router.Start("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
