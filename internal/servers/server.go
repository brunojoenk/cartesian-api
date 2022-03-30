package servers

import (
	"net/http"
	"time"

	"github.com/brunojoenk/cartesian-api/internal/routers"
)

var TEN_SECONDS = 10 * time.Second
var ADDRESS = "127.0.0.1:8080"

func NewServer() http.Server {
	router := routers.CreateRouters()
	return http.Server{
		Handler:      router,
		Addr:         ADDRESS,
		WriteTimeout: TEN_SECONDS,
		ReadTimeout:  TEN_SECONDS,
	}
}
