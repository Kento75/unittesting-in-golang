package controllers

import (
	"encoding/json"
	"github.com/Kento75/unittesting-in-golang/golang-testing-echo/src/api/domain/locations"
	"github.com/Kento75/unittesting-in-golang/golang-testing-echo/src/api/services"
	"github.com/Kento75/unittesting-in-golang/golang-testing-echo/src/api/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Mock
var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	// Arrange
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("country_id")
	c.SetParamValues("AR")

	// Act
	GetCountry(c)

	// Assert
	assert.EqualValues(t, http.StatusNotFound, rec.Code)

	var apiError errors.ApiError
	err := json.Unmarshal(rec.Body.Bytes(), &apiError)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "Country not found", apiError.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// Arrange
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("country_id")
	c.SetParamValues("AR")
	c.SetParamNames("name")
	c.SetParamValues("Argentina")

	// Act
	GetCountry(c)

	// Assert
	assert.EqualValues(t, http.StatusOK, rec.Code)

	var country locations.Country
	err := json.Unmarshal(rec.Body.Bytes(), &country)
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
