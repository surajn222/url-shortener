package storage

import "fmt"

// ShortUrls := make(map[string]string)
var ShortUrls map[string]string

func StoreShortenedLinks(link string, shortlink string) map[string]string {
	fmt.Println("Shortened Links Storage")
	fmt.Println("Adding links to Map: ", link)
	if ShortUrls == nil {
		ShortUrls = make(map[string]string)
	}
	ShortUrls[link] = shortlink
	fmt.Printf("%v\n\n", ShortUrls)
	return ShortUrls
}

func GetAllLinks() map[string]string {
	return ShortUrls
}
