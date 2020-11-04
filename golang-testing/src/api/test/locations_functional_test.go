package test

import (
	"encoding/json"
	"fmt"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCountriesNotFound(t *testing.T) {
	fmt.Println("about to functional test get countries")

	// Arrange
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "no country with id AR","error":"not_found","status":404,"cause":[]}`,
	})

	// Act
	response, err := http.Get("http://127.0.0.1:8080/locations/countries/AR")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id AR", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T) {
	//TODO:
}
