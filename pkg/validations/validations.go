package validations

import (
	"net/url"
	"strings"

	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

// You can use as well instead of *url.URL a string with your URL, in this case you might drop `u.String()`
// - transformation to string in method call.
func validateURL(u string) error {
	valid := govalidator.IsRequestURL(u)
	if valid == false {
		return fmt.Errorf("%v is a invalid url", u)
	}
	return nil
}

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
