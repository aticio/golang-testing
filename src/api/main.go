package main

import (
	"fmt"

	"github.com/aticio/golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("TR")

	fmt.Println(err)
	fmt.Println(country)
}
