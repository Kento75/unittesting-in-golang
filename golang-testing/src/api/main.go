package main

import (
	"fmt"
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/providers/locations_provider"
)

// 動作確認
func main() {
	country, err := locations_provider.GetCountry("AR")
	fmt.Println(err)
	fmt.Println(country)
}
