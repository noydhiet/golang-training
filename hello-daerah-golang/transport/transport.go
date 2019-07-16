package transport

import (
	"aph-go-service/logging"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"hello-daerah/datastruct"
	"hello-daerah/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type DaerahService interface {
	HelloDaerahService(context.Context, string, string, string) string
}

type daerahService struct{}

var ErrEmpty = errors.New("empty string")

func (daerahService) HelloDaerahService(_ context.Context, name, sex, origin string) string {
	return call_ServiceHelloDaerah(name, sex, origin)
}

func call_ServiceHelloDaerah(name, sex, origin string) string {

	messageResponse := service.HelloDaerah(name, sex, origin)

	return messageResponse
}

func makeHelloDaerahEndpoint(daerah DaerahService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.HelloDaerahRequest)
		logging.Log(fmt.Sprintf("Name Request %s", req.NAME))
		logging.Log(fmt.Sprintf("Sex Request %s", req.SEX))
		logging.Log(fmt.Sprintf("Origin Request %s", req.ORIGIN))
		v := daerah.HelloDaerahService(ctx, req.NAME, req.SEX, req.ORIGIN)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.HelloDaerahResponse{v}, nil
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
	daerah := daerahService{}

	HelloDaerahHandler := httptransport.NewServer(
		makeHelloDaerahEndpoint(daerah),
		decodeHelloDaerahRequest,
		encodeResponse,
	)

	http.Handle("/HelloDaerah", HelloDaerahHandler)
}
