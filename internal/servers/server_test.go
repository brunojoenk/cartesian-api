package servers

import (
	"testing"
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
