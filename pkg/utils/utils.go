package utils

import (
	"log"
	"net/url"
	"strings"
)

func GetDomainFromUrl(urlInput string) string {
	if !strings.HasPrefix(urlInput, "https://") && !strings.HasPrefix(urlInput, "http://") {
		urlInput = "http://" + urlInput
	}

	url, err := url.Parse(urlInput)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Count(url.Hostname(), ".") == 2 {
		arrDomainName := strings.Split(url.Hostname(), ".")
		domainName := arrDomainName[len(arrDomainName)-2] + "." + arrDomainName[len(arrDomainName)-1]
		return domainName
	} else {
		return url.Hostname()
	}

}
