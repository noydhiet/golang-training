package main

import (
	_ "expvar"
	"net/http"
	"os"
	"servise-api/transport"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("Listen.error", err)
	}
}
