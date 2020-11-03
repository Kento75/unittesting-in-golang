package app

import (
	"github.com/Kento75/unittesting-in-golang/golang-testing-echo/src/api/controllers"
)

func mapUrls() {
	// Routes
	router.GET("/", "/locations/countries/:country_id", controllers.GetCountry)
}
