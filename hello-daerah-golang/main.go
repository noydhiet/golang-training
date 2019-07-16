package main

import (
	_ "expvar"
	//"hello-daerah/service"
	"hello-daerah/transport"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	//service.HelloDaerah(string, string, string)

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
