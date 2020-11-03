package app

import (
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
