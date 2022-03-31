package servers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brunojoenk/cartesian-api/internal/routers"
)

var TWENTY_SECONDS = 20 * time.Second
var DEFAULT_PORT = "8080"

func NewServer() http.Server {
	router := routers.CreateRouters()
	return http.Server{
		Handler:      router,
		Addr:         ":" + getPort(),
		WriteTimeout: TWENTY_SECONDS,
		ReadTimeout:  TWENTY_SECONDS,
	}
}

func getPort() string {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		return DEFAULT_PORT
	}

	return port
}
