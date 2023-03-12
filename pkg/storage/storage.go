package storage

import (
	"github.com/sirupsen/logrus"
)

// ShortUrls := make(map[string]string)
var ShortUrls map[string]string

func StoreShortenedLinks(link string, shortlink string) map[string]string {
	logrus.Println("Shortened Links Storage")
	logrus.Println("Adding links to Map: ", link)
	if ShortUrls == nil {
		ShortUrls = make(map[string]string)
	}
	ShortUrls[shortlink] = link
	logrus.Printf("%v\n\n", ShortUrls)
	return ShortUrls
}

func GetAllLinks() map[string]string {
	return ShortUrls
}
