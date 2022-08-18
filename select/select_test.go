package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
    t.Run("should return the fastest URL", func(t *testing.T) {
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(0 * time.Millisecond)

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got, _ := Racer(slowURL, fastURL)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }

        slowServer.Close()
        fastServer.Close()
    })

    t.Run("returns an error if a server doesn't respond within 10 secons", func(t *testing.T) {
        server := makeDelayedServer(2 * time.Microsecond)

        defer server.Close()

        _, err := ConfigurableRacer(server.URL, server.URL, 1 * time.Microsecond)

        if err == nil {
            t.Error("expected an error but didn't got")
        }
    })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
