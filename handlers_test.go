package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	t.Run("it should return status code 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		IndexHandler(response, request)

		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSearchHandler(t *testing.T) {
	t.Run("it should return status code 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=something", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("it should return status code 303", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := response.Code
		want := http.StatusSeeOther
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
