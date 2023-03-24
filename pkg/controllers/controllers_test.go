package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	storage "github.com/surajn222/url-shortener/pkg/storage"
)

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
		ctx := req.Context()
		ctx = context.WithValue(ctx, "storage", &storage.MemStorage{})
		req = req.WithContext(ctx)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(MuxUrlShorten)
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

	storage.Cleanup()

	req, err := http.NewRequest("GET", "/urlshorten?url="+"yahoo.com", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "storage", &storage.MemStorage{})
	req = req.WithContext(ctx)

	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(MuxUrlShorten)
	handler2.ServeHTTP(rr2, req)

	req, err = http.NewRequest("GET", "/urlshorten?url="+"google.com", nil)
	ctx = req.Context()
	ctx = context.WithValue(ctx, "storage", &storage.MemStorage{})

	req = req.WithContext(ctx)
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(MuxUrlShorten)
	handler3.ServeHTTP(rr3, req)

	req, err = http.NewRequest("GET", "/domaincount", nil)
	ctx = req.Context()
	ctx = context.WithValue(ctx, "storage", &storage.MemStorage{})
	req = req.WithContext(ctx)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MuxDomainCount)
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
