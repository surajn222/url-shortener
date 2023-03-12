package validations

import (
	"fmt"
	"testing"
)

func Test_Validations(t *testing.T) {

	type urlTest struct {
		name string
		url  string
	}

	arrUrlInput := []urlTest{
		{name: "www.google.com", url: "google.com"},
		{name: "www.google.com", url: "google.com"},
	}

	for _, v := range arrUrlInput {
		err := UrlValidation(v.url)
		fmt.Println(err)

		if err != nil {
			t.Errorf("Invalid URL: got Error %v, want %v",
				err, nil)
		}
	}
}
