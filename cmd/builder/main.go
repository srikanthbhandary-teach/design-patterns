package main

import (
	"design-patterns/creational"
	"fmt"
)

func main() {
	fmt.Println("-----------------------------")
	fmt.Println("Builder Pattern Demo")
	// Create API request builder
	apiRequestBuilder := creational.NewApiRequestBuilder()
	request := apiRequestBuilder.SetUrl("https://example.com").
		SetBody("body").
		SetMethod("GET").
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetQueryParams(map[string]string{"key": "value"}).Build()
	request.Print()
	fmt.Println("-----------------------------")
}
