package servers

import (
	"testing"
)

func TestCreateServer(t *testing.T) {
	server := NewServer()

	if server.Addr != "127.0.0.1:8080" {
		t.Error("Error on configure address server")
	}
	if server.ReadTimeout.Microseconds() != TEN_SECONDS.Microseconds() {
		t.Error("Error on configure read timeout server")
	}
	if server.WriteTimeout.Microseconds() != TEN_SECONDS.Microseconds() {
		t.Error("Error on configure write timeout server")
	}
	if server.Handler == nil {
		t.Error("Error on configure handlers of server")
	}
}
