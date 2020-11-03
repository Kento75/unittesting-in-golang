package controllers

import (
	"encoding/json"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/domain/locations"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/services"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/gin-gonic/gin"
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

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	// Act
	GetCountry(c)

	// Assert
	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiError errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
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

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	// Act
	GetCountry(c)

	// Assert
	assert.EqualValues(t, http.StatusOK, response.Code)

	var country locations.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
