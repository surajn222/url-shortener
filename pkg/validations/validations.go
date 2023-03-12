package validations

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

func UrlValidation(urlInput string) error {

	if !strings.HasPrefix(urlInput, "https://") || !strings.HasPrefix(urlInput, "http://") {
		urlInput = "http://" + urlInput
	}

	_, err := url.ParseRequestURI(urlInput)
	if err != nil {
		logrus.Printf("URL is not valid: %+v", err)
		return err
	}

	return nil
}
