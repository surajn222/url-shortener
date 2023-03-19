package storage

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

// ShortUrls := make(map[string]string)
var MapShortUrls map[string]string
var MapDomainCount map[string]int

type MemStorage struct {
}

func (*MemStorage) InsertShortenedLinks(link string, shortlink string) error {
	logrus.Println("Adding links to Map: ", link)

	if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
		link = "http://" + link
	}

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

func (*MemStorage) GetAllLinks() map[string]string {
	return MapShortUrls
}

func (*MemStorage) UpdateDomainCount(domainName string) error {
	if MapDomainCount == nil {
		MapDomainCount = make(map[string]int)
	}

	MapDomainCount[domainName]++
	return nil
}

func (*MemStorage) GetDomainCount() map[string]int {

	MapDomainCountOrdered := make(map[string]int)
	keys := make([]string, 0, len(MapDomainCount))
	for k := range MapDomainCount {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return MapDomainCount[keys[i]] > MapDomainCount[keys[j]]
	})

	for _, k := range keys {
		MapDomainCountOrdered[k] = MapDomainCount[k]
	}
	return MapDomainCountOrdered
}

func (*MemStorage) GetLink(shortUrl string) (string, error) {
	if val, ok := MapShortUrls[shortUrl]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("Unable to find url")
	}

}

func Cleanup() {
	MapDomainCount = nil
}
