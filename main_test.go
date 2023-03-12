package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_muxUrlShorten(t *testing.T) {

	type urlShorten struct {
		name     string
		method   string
		url      string
		expected string
	}

	urlShorten1 := urlShorten{name: "google.com", method: "GET", url: "/?shortenurl=google.com", expected: "google.com:1d5920f"}

	req, err := http.NewRequest(urlShorten1.method, urlShorten1.url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(muxUrlShorten)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := urlShorten1.expected
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
