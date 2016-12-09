package logger

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpTransportWithFakeServer(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"ok": true}`)
	}))
	defer testServer.Close()

	jsonFormat := NewJSONFormat()
	ht := NewHttpTransport(testServer.URL, jsonFormat)

	log := New(make(map[string]string))
	log.AddTransport(ht)

	err := log.Info("SomeEvent", "Message...")

	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestHttpTransportWithNoServer(t *testing.T) {
	jsonFormat := NewJSONFormat()
	ht := NewHttpTransport("http://test-go.travix.com/post", jsonFormat)

	log := New(make(map[string]string))
	log.AddTransport(ht)

	err := log.Info("SomeEvent", "Message...")

	if err == nil {
		t.Errorf("expected request to fail")
	}
}
