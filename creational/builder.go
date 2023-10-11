package creational

import "fmt"

type ApiRequest struct {
	Url         string
	Method      string
	Body        string
	Headers     map[string]string
	QueryParams map[string]string
}

func (a *ApiRequest) Print() {
	fmt.Printf("Request : %+v\n", a)
}

type ApiRequestBuilder interface {
	SetUrl(url string) ApiRequestBuilder
	SetMethod(method string) ApiRequestBuilder
	SetBody(body string) ApiRequestBuilder
	SetHeaders(headers map[string]string) ApiRequestBuilder
	SetQueryParams(queryParams map[string]string) ApiRequestBuilder
	Build() ApiRequest
}

type apiRequestBuilder struct {
	request ApiRequest
}

// Build implements ApiRequestBuilder.
func (a *apiRequestBuilder) Build() ApiRequest {
	return a.request
}

// SetBody implements ApiRequestBuilder.
func (a *apiRequestBuilder) SetBody(body string) ApiRequestBuilder {
	a.request.Body = body
	return a
}

// SetHeaders implements ApiRequestBuilder.
func (a *apiRequestBuilder) SetHeaders(headers map[string]string) ApiRequestBuilder {
	a.request.Headers = headers
	return a
}

// SetMethod implements ApiRequestBuilder.
func (a *apiRequestBuilder) SetMethod(method string) ApiRequestBuilder {
	a.request.Method = method
	return a
}

// SetQueryParams implements ApiRequestBuilder.
func (a *apiRequestBuilder) SetQueryParams(queryParams map[string]string) ApiRequestBuilder {
	a.request.QueryParams = queryParams
	return a
}

// SetUrl implements ApiRequestBuilder.
func (a *apiRequestBuilder) SetUrl(url string) ApiRequestBuilder {
	a.request.Url = url
	return a
}

func NewApiRequestBuilder() ApiRequestBuilder {
	return &apiRequestBuilder{}
}
