package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/budi/ゴランぐ/src/golang-service/datastruct"
	"github.com/budi/ゴランぐ/src/golang-service/logging"
	"github.com/budi/ゴランぐ/src/golang-service/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	HelloWorldService(context.Context, string) string
	HelloDaerahService(context.Context, string, string, string) string
}

type aphService struct{}

var ErrEmpty = errors.New("empty string")

func (aphService) HelloWorldService(_ context.Context, name string) string {

	return call_ServiceHelloWorld(name)
}

func (aphService) HelloDaerahService(_ context.Context, name string, gender string, city string) string {

	return call_ServiceHelloDaerah(name, gender, city)
}

func call_ServiceHelloWorld(name string) string {

	messageResponse := service.HelloWorld(name)

	return messageResponse

}

func call_ServiceHelloDaerah(name string, gender string, city string) string {

	messageResponse := service.HelloDaerah(name, gender, city)

	return messageResponse

}

func makeHelloWorldEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloWorldRequest)
		logging.Log(fmt.Sprintf("Name Request %s", req.NAME))
		v := aph.HelloWorldService(ctx, req.NAME)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.HelloWorldResponse{v}, nil
	}
}

func makeHelloDaerahEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloDaerahRequest)
		logging.Log(fmt.Sprintf("Name Request %s", req.NAME, "Gender Request %s", req.GENDER, "City Request %s", req.CITY))
		v := aph.HelloDaerahService(ctx, req.NAME, req.GENDER, req.CITY)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.HelloDaerahResponse{200, v}, nil
	}
}

func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeHelloDaerahRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.HelloDaerahRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	HelloWorldHandler := httptransport.NewServer(
		makeHelloWorldEndpoint(aph),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	HelloDaerahHandler := httptransport.NewServer(
		makeHelloDaerahEndpoint(aph),
		decodeHelloDaerahRequest,
		encodeResponse,
	)

	http.Handle("/HelloWorld", HelloWorldHandler)
	http.Handle("/HelloDaerah", HelloDaerahHandler)
}
