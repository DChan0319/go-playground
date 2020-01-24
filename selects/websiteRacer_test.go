package selects

import (
	"net/http/httptest"
	"net/http"
	"time"
	"testing"
)

func TestRacer(t *testing.T) {
	t.Run("returns fast url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
	
		expected := fastUrl
		actual, err := Racer(slowUrl, fastUrl)

		if (err != nil) {
			t.Error("Received an error when expecting none.")
		}

		if (actual != expected) {
			t.Errorf("expected: %q, actual: %q", expected, actual)
		}
	})

	t.Run("returns an error if server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(24 * time.Millisecond)
		
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 1 * time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}
