package servers

import (
	"net/http"
	"time"

	"github.com/brunojoenk/cartesian-api/internal/routers"
)

var TWENTY_SECONDS = 20 * time.Second
var PORT = ":8080"

func NewServer() http.Server {
	router := routers.CreateRouters()
	return http.Server{
		Handler:      router,
		Addr:         PORT,
		WriteTimeout: TWENTY_SECONDS,
		ReadTimeout:  TWENTY_SECONDS,
	}
}
