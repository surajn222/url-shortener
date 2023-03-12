package storage

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
)

// ShortUrls := make(map[string]string)
var MapShortUrls map[string]string
var MapDomainCount map[string]int

func StoreShortenedLinks(link string, shortlink string) error {
	logrus.Println("Shortened Links Storage")
	logrus.Println("Adding links to Map: ", link)
	if MapShortUrls == nil {
		MapShortUrls = make(map[string]string)
	}
	MapShortUrls[shortlink] = link
	logrus.Printf("%v\n\n", MapShortUrls)

	if _, ok := MapShortUrls[shortlink]; ok {
		return nil
	}

	return fmt.Errorf("unable to shorten url")
}

func GetAllLinks() map[string]string {
	return MapShortUrls
}

func UpdateDomainCount(domainName string) error {
	if MapDomainCount == nil {
		MapDomainCount = make(map[string]int)
	}
	// counter := map[string]int{}
	fmt.Println(MapDomainCount)
	MapDomainCount[domainName]++
	fmt.Println(MapDomainCount)
	return nil
}

func GetDomainCount() map[string]int {

	MapDomainCountOrdered := make(map[string]int)
	keys := make([]string, 0, len(MapDomainCount))
	for k := range MapDomainCount {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return MapDomainCount[keys[i]] > MapDomainCount[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, MapDomainCount[k])
		MapDomainCountOrdered[k] = MapDomainCount[k]
	}
	return MapDomainCountOrdered
}

func GetLink(shortUrl string) (string, error) {

	fmt.Println("shortURL: ", shortUrl)
	if val, ok := MapShortUrls[shortUrl]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("Unable to find url")
	}

}
