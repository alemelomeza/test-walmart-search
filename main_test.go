package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("without query it should get status code 303", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := response.Code
		want := 303
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("with query \"something\" it should get status code 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=something", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := response.Code
		want := 200
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("with query \"12\" it should get one result", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=12", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := strings.Count(response.Body.String(), "<div class=\"card\">")
		want := 1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("with query \"121\" it should get one result with discount", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=121", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := strings.Count(response.Body.String(), "<span class=\"card-discount\">")
		want := 1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("with query \"qac\" it should get results with discount", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=qac", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := strings.Count(response.Body.String(), "<span class=\"card-discount\">")
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("with query \"hqh\" it should get results with discount", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/search?q=hqh", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)

		got := strings.Count(response.Body.String(), "<span class=\"card-discount\">")
		want := 2
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
