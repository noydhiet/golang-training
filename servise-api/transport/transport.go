package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"servise-api/datastruct"
	service "servise-api/servise"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type WaService interface {
	HelloDaerahService(context.Context, string, string, string) string
}

type waService struct{}

var ErrEmpty = errors.New("empty string")

func (waService) HelloDaerahService(_ context.Context, name, jenisKelamin, asalKota string) string {

	return call_ServiceHelloDaerah(name, jenisKelamin, asalKota)
}
func call_ServiceHelloDaerah(name, jenisKelamin, asalKota string) string {

	messageResponse := service.HelloDaerah(name, jenisKelamin, asalKota)

	return messageResponse

}

func makeHelloDaerahEndpoint(me WaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloDaerahRequest)

		v := me.HelloDaerahService(ctx, req.NAME, req.JENIS_KELAMIN, req.ASAL_KOTA)

		return datastruct.HelloDaerahResponse{200, v}, nil
	}
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
	wa := waService{}

	HelloDaerahHandler := httptransport.NewServer(
		makeHelloDaerahEndpoint(wa),
		decodeHelloDaerahRequest,
		encodeResponse,
	)

	http.Handle("/HelloDaerah", HelloDaerahHandler)
}
