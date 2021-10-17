package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("returns geeting", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Greeting(response, request)

		got := response.Body.String()
		want := "Hello World!"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
