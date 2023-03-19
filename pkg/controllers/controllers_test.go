package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/surajn222/url-shortener/pkg/controllers"
	storage "github.com/surajn222/url-shortener/pkg/storage/mem_storage"
)

// func Test_muxUrlShorten(t *testing.T) {

// 	type urlShorten struct {
// 		name     string
// 		method   string
// 		url      string
// 		expected string
// 	}

// 	urlShorten1 := []urlShorten{
// 		{name: "google.com", method: "GET", url: "/shortenurl?url=google.com", expected: "google.com:1d5920f"},
// 		{name: "yahoo.com", method: "GET", url: "/shortenurl?url=yahoo.com", expected: "yahoo.com:50cd1a9"},
// 	}

// 	for _, k := range urlShorten1 {
// 		req, err := http.NewRequest(k.method, k.url, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(controllers.MuxUrlShorten)
// 		handler.ServeHTTP(rr, req)
// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v",
// 				status, http.StatusOK)
// 		}

// 		// Check the response body is what we expect.
// 		expected := k.expected
// 		if rr.Body.String() != expected {
// 			t.Errorf("handler returned unexpected body: got %v want %v",
// 				rr.Body.String(), expected)
// 		}
// 	}
// }

// func Test_getAllLinks(t *testing.T) {

// 	type structGetLinks struct {
// 		name     string
// 		method   string
// 		url      string
// 		expected string
// 	}

// 	getLinks := []structGetLinks{
// 		{name: "getLinks", method: "GET", url: "/getlinks", expected: "abcd"},
// 	}

// 	for _, k := range getLinks {
// 		req, err := http.NewRequest(k.method, k.url, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(controllers.MuxUrlShorten)
// 		handler.ServeHTTP(rr, req)
// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v",
// 				status, http.StatusOK)
// 		}

// 		// Check the response body is what we expect.
// 		expected := k.expected
// 		if rr.Body.String() != expected {
// 			t.Errorf("handler returned unexpected body: got %v want %v",
// 				rr.Body.String(), expected)
// 		}
// 	}
// }

// func TestMyHandler(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/path", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.MuxPath)

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := "Path"
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

func Test_MuxUrlShorten(t *testing.T) {

	type testMuxUrl struct {
		name     string
		url      string
		shorturl string
	}

	testCases := []testMuxUrl{
		{name: "google.com", url: "google.com", shorturl: "1d5920f"},
		{name: "yahoo.com", url: "yahoo.com", shorturl: "50cd1a9"},
		{name: "chat.openai.com", url: "chat.openai.com", shorturl: "234455f"},
	}

	for _, v := range testCases {

		req, err := http.NewRequest("GET", "/urlshorten?url="+v.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(controllers.MuxUrlShorten)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "{" + v.url + ":" + "http://localhost:8081/" + v.shorturl + "}"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

func Test_MuxDomainCount(t *testing.T) {

	// type testMuxUrl struct {
	// 	name     string
	// 	url      string
	// 	shorturl string
	// }

	// testCases := []testMuxUrl{
	// 	{name: "google.com", url: "google.com", shorturl: "1d5920f"},
	// 	{name: "yahoo.com", url: "yahoo.com", shorturl: "50cd1a9"},
	// 	{name: "chat.openai.com", url: "chat.openai.com", shorturl: "234455f"},
	// }

	// for _, v := range testCases {

	storage.Cleanup()

	req, err := http.NewRequest("GET", "/urlshorten?url="+"yahoo.com", nil)
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(controllers.MuxUrlShorten)
	handler2.ServeHTTP(rr2, req)

	req, err = http.NewRequest("GET", "/urlshorten?url="+"google.com", nil)
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(controllers.MuxUrlShorten)
	handler3.ServeHTTP(rr3, req)

	req, err = http.NewRequest("GET", "/domaincount", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.MuxDomainCount)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"google.com\":1,\"yahoo.com\":1}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	// }
}
