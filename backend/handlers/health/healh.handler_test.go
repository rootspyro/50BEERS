package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler HealthHandler = *NewHealthHandler()

func TestServerStatus(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handler.ServerStatus))

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
}
