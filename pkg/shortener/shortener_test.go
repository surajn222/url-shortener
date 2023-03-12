package shortener

import "testing"

func Test_getMD5Hash(t *testing.T) {
	type textTest struct {
		name     string
		text     string
		expected string
	}

	arrTextInput := []textTest{
		{name: "google.com", text: "google.com", expected: "1d5920f4b44b27a802bd77c4f0536f5a"},
		{name: "yahoo.com", text: "yahoo.com", expected: "50cd1a9a183758039b0841aa738c3f0b"},
	}

	for _, v := range arrTextInput {
		md5Hash := getMD5Hash(v.text)

		if md5Hash != v.expected {
			t.Errorf("Invalid Md5 Hash: got Error %v, want %v", md5Hash, v.expected)
		}
	}
}

func Test_urlShorten(t *testing.T) {
	type textTest struct {
		name     string
		text     string
		expected string
	}

	arrTextInput := []textTest{
		{name: "google.com", text: "google.com", expected: "1d5920f"},
		{name: "yahoo.com", text: "yahoo.com", expected: "50cd1a9"},
	}

	for _, v := range arrTextInput {
		md5Hash := UrlShorten(v.text)

		if md5Hash != v.expected {
			t.Errorf("Invalid Md5 Hash: got Error %v, want %v", md5Hash, v.expected)
		}
	}
}
