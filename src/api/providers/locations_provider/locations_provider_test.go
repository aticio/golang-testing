package locations_provider

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryRestClientError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/TR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to get country TR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/TR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`,
	})
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/TR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
	})
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country TR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/TR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 123,"name": "Turkey","time_zone": "GMT+03:00"}`,
	})
	country, err := GetCountry("TR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for TR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/TR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"TR","name":"Turkey","locale":"en_US","currency_id":"USD","decimal_separator":".","thousands_separator":",","time_zone":"GMT+03:00","geo_information":{"location":{"latitude":39.0100769,"longitude":30.6887968}},"states":[{"id":"TR-01","name":"Adana"},{"id":"TR-02","name":"Ad??yaman"},{"id":"TR-03","name":"Afyonkarahisar"},{"id":"TR-68","name":"Aksaray"},{"id":"TR-05","name":"Amasya"},{"id":"TR-06","name":"Ankara"},{"id":"TR-07","name":"Antalya"},{"id":"TR-75","name":"Ardahan"},{"id":"TR-08","name":"Artvin"},{"id":"TR-09","name":"Ayd??n"},{"id":"TR-04","name":"A??r??"},{"id":"TR-10","name":"Bal??kesir"},{"id":"TR-74","name":"Bart??n"},{"id":"TR-72","name":"Batman"},{"id":"TR-69","name":"Bayburt"},{"id":"TR-11","name":"Bilecik"},{"id":"TR-12","name":"Bing??l"},{"id":"TR-13","name":"Bitlis"},{"id":"TR-14","name":"Bolu"},{"id":"TR-15","name":"Burdur"},{"id":"TR-16","name":"Bursa"},{"id":"TR-20","name":"Denizli"},{"id":"TR-21","name":"Diyarbak??r"},{"id":"TR-81","name":"D??zce"},{"id":"TR-22","name":"Edirne"},{"id":"TR-23","name":"Elaz????"},{"id":"TR-24","name":"Erzincan"},{"id":"TR-25","name":"Erzurum"},{"id":"TR-26","name":"Eski??ehir"},{"id":"TR-35","name":"Esmirna"},{"id":"TR-34","name":"Estambul"},{"id":"TR-27","name":"Gaziantep"},{"id":"TR-28","name":"Giresun"},{"id":"TR-29","name":"G??m????hane"},{"id":"TR-30","name":"Hakk??ri"},{"id":"TR-31","name":"Hatay"},{"id":"TR-32","name":"Isparta"},{"id":"TR-76","name":"I??d??r"},{"id":"TR-46","name":"Kahramanmara??"},{"id":"TR-78","name":"Karab??k"},{"id":"TR-70","name":"Karaman"},{"id":"TR-36","name":"Kars"},{"id":"TR-37","name":"Kastamonu"},{"id":"TR-38","name":"Kayseri"},{"id":"TR-79","name":"Kilis"},{"id":"TR-39","name":"Kirklareli"},{"id":"TR-41","name":"Kocaeli"},{"id":"TR-42","name":"Konya"},{"id":"TR-43","name":"K??tahya"},{"id":"TR-71","name":"K??r??kkale"},{"id":"TR-40","name":"K??r??ehir"},{"id":"TR-44","name":"Malatya"},{"id":"TR-45","name":"Manisa"},{"id":"TR-47","name":"Mardin"},{"id":"TR-33","name":"Mersin"},{"id":"TR-49","name":"Mus"},{"id":"TR-48","name":"Mu??la"},{"id":"TR-50","name":"Nevsehir"},{"id":"TR-51","name":"Ni??de"},{"id":"TR-52","name":"Ordu"},{"id":"TR-80","name":"Osmaniye"},{"id":"TR-53","name":"Rize"},{"id":"TR-54","name":"Sakarya"},{"id":"TR-55","name":"Samsun"},{"id":"TR-56","name":"Siirt"},{"id":"TR-57","name":"Sinope"},{"id":"TR-58","name":"Sivas"},{"id":"TR-59","name":"Tekirda??"},{"id":"TR-60","name":"Tokat"},{"id":"TR-61","name":"Trebisonda"},{"id":"TR-62","name":"Tunceli"},{"id":"TR-64","name":"U??ak"},{"id":"TR-65","name":"Van"},{"id":"TR-77","name":"Yalova"},{"id":"TR-66","name":"Yozgat"},{"id":"TR-67","name":"Zonguldak"},{"id":"TR-17","name":"??anakkale"},{"id":"TR-18","name":"??ankiri"},{"id":"TR-19","name":"??orum"},{"id":"TR-63","name":"??anl??urfa"},{"id":"TR-73","name":"????rnak"}]}`,
	})
	country, err := GetCountry("TR")

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "TR", country.Id)
	assert.EqualValues(t, "Turkey", country.Name)
	assert.EqualValues(t, "GMT+03:00", country.TimeZone)
	assert.EqualValues(t, 81, len(country.States))
}
