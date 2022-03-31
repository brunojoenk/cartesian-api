package servers

import (
	"os"
	"testing"

	"github.com/brunojoenk/cartesian-api/internal/utils"
)

func TestCreateServer(t *testing.T) {
	server := NewServer()

	if server.Addr != ":8080" {
		t.Error("Error on configure address server")
	}
	if server.ReadTimeout.Microseconds() != TWENTY_SECONDS.Microseconds() {
		t.Error("Error on configure read timeout server")
	}
	if server.WriteTimeout.Microseconds() != TWENTY_SECONDS.Microseconds() {
		t.Error("Error on configure write timeout server")
	}
	if server.Handler == nil {
		t.Error("Error on configure handlers of server")
	}
}

func TestGetPort(t *testing.T) {
	expectedPort := "3000"
	os.Setenv("PORT", expectedPort)

	server := NewServer()

	if server.Addr != utils.ConcatStrings(":", expectedPort) {
		t.Error("Error on configure port server")
	}
}
