package function

import (
	"fmt"
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	message := fmt.Sprintf("Hello from GKE Cluster!")

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
