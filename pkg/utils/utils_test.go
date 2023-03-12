package utils

import "testing"

func Test_getDomainFromUrl(t *testing.T) {
	type textTest struct {
		name     string
		text     string
		expected string
	}

	arrTextInput := []textTest{
		{name: "google.com", text: "www.google.com", expected: "google.com"},
		{name: "yahoo.com", text: "yahoo.com", expected: "yahoo.com"},
		{name: "github.com", text: "https://github.com/pvillard31/url-shortener-go", expected: "github.com"},
	}

	for _, v := range arrTextInput {
		domainName := GetDomainFromUrl(v.text)

		if domainName != v.expected {
			t.Errorf("Invalid domain: got Error %v, want %v", domainName, v.expected)
		}
	}
}
