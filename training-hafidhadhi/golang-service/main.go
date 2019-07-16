package main

import (
	_ "expvar"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/hafidhadhi/training/golang-service/transport"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	//port := os.Getenv("PORT")
	//if port == "" {
	port := "8080"
	//}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
