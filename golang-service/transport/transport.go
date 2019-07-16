package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"golang-training/golang-service/datastruct"
	"golang-training/golang-service/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type MeService interface {
	HelloWorldService(context.Context, string) string
	HelloDaerahService(context.Context, string, string, string) string
}

type meService struct{}

var ErrEmpty = errors.New("empty string")

func (meService) HelloWorldService(_ context.Context, name string) string {

	return call_ServiceHelloWorld(name)
}

func (meService) HelloDaerahService(_ context.Context, name, jenisKelamin, asalKota string) string {

	return call_ServiceHelloDaerah(name, jenisKelamin, asalKota)
}

func call_ServiceHelloWorld(name string) string {

	messageResponse := service.HelloWorld(name)

	return messageResponse

}

func call_ServiceHelloDaerah(name, jenisKelamin, asalKota string) string {

	messageResponse := service.HelloDaerah(name, jenisKelamin, asalKota)

	return messageResponse

}

func makeHelloWorldEndpoint(me MeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloWorldRequest)

		v := me.HelloWorldService(ctx, req.NAME)

		return datastruct.HelloWorldResponse{200, v}, nil
	}
}

func makeHelloDaerahEndpoint(me MeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloDaerahRequest)

		v := me.HelloDaerahService(ctx, req.NAME, req.JENIS_KELAMIN, req.ASAL_KOTA)

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
	me := meService{}

	HelloWorldHandler := httptransport.NewServer(
		makeHelloWorldEndpoint(me),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/HelloWorld", HelloWorldHandler)

	HelloDaerahHandler := httptransport.NewServer(
		makeHelloDaerahEndpoint(me),
		decodeHelloDaerahRequest,
		encodeResponse,
	)

	http.Handle("/HelloDaerah", HelloDaerahHandler)
}
