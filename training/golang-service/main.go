package main

import (
	_ "expvar"
	"net/http"
	"os"
	"training/golang-service/transport"

	"github.com/go-kit/kit/log"
)

func main() {

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
