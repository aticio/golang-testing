package locations_provider

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountryRestClientError(t *testing.T) {

	// execution
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client error when getting country TR", err.Message)
	// validation
}

func TestGetCountryNotFound(t *testing.T) {
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country TR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for TR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	country, err := GetCountry("TR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "TR", country.Id)
	assert.EqualValues(t, "Turkey", country.Name)
	assert.EqualValues(t, "GMT+03:00", country.TimeZone)
	assert.EqualValues(t, 81, len(country.States))
}
