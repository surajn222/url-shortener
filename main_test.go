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

	urlShorten1 := []urlShorten{
		{name: "google.com", method: "GET", url: "/shortenurl?url=google.com", expected: "google.com:1d5920f"},
		{name: "yahoo.com", method: "GET", url: "/shortenurl?url=yahoo.com", expected: "yahoo.com:50cd1a9"},
	}

	for _, k := range urlShorten1 {
		req, err := http.NewRequest(k.method, k.url, nil)
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
		expected := k.expected
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

func Test_getAllLinks(t *testing.T) {

	type structGetLinks struct {
		name     string
		method   string
		url      string
		expected string
	}

	getLinks := []structGetLinks{
		{name: "getLinks", method: "GET", url: "/getlinks", expected: "abcd"},
	}

	for _, k := range getLinks {
		req, err := http.NewRequest(k.method, k.url, nil)
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
		expected := k.expected
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
