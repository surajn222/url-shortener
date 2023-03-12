package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/surajn222/url-shortener/pkg/shortener"
	storage "github.com/surajn222/url-shortener/pkg/storage/redis_storage"
	"github.com/surajn222/url-shortener/pkg/utils"
	"github.com/surajn222/url-shortener/pkg/validations"
)

func MuxDomainCount(response http.ResponseWriter, request *http.Request) {
	jsonDomainCount := storage.GetDomainCount()
	mapDomainCount, err := json.Marshal(jsonDomainCount)
	if err != err {
		fmt.Fprintf(response, "unable to get domain count from database")
	} else {
		fmt.Fprintf(response, string(mapDomainCount))
	}

}

func MuxGetLinks(response http.ResponseWriter, request *http.Request) {
	jsonAllLinks := storage.GetAllLinks()
	mapAllLinks, err := json.Marshal(jsonAllLinks)
	if err != err {
		fmt.Fprintf(response, "unable to get mux links")
	} else {
		fmt.Fprintf(response, string(mapAllLinks))
	}

}

func MuxUrlShorten(response http.ResponseWriter, request *http.Request) {
	url := request.URL.Query().Get("url")

	// Validate if input is URL
	err := validations.UrlValidation(url)
	if err == nil {
		// Shorten URL with suitable algorithm
		shortenedUrl := shortener.UrlShorten(url)
		responseString := "{" + url + ":" + "http://localhost:8081/" + shortenedUrl + "}"

		// Store shortened links to database
		err1 := storage.StoreShortenedLinks(url, shortenedUrl)
		domainName := utils.GetDomainFromUrl(url)
		err2 := storage.UpdateDomainCount(domainName)
		if err2 != nil && err1 != nil {
			panic(err2)
			fmt.Fprintf(response, "unable to shorten URL")
		}

		fmt.Fprintf(response, responseString)
	} else {
		fmt.Fprintf(response, "invalid URL")
	}
}

func MuxRedirect(response http.ResponseWriter, request *http.Request) {
	// Get link from database and redirect
	strShortenedUrl, err := storage.GetLink(request.URL.Path[1:])
	fmt.Println(strShortenedUrl)
	if err != nil {
		fmt.Fprintf(response, "unable to find URL")
	} else {
		http.Redirect(response, request, strShortenedUrl, http.StatusMovedPermanently)
	}
}

func MuxIndex(response http.ResponseWriter, request *http.Request) {
	responseString := `
a. To shorten url:
	http://localhost:8081/shortenurl?url=google.com
	http://localhost:8081/shortenurl?url=chat.openai.com

b. To get domain count of which URLs were shortened:
    http://localhost:8081/domaincount

c. Redirect to :
	http://localhost:8081/1d5920f
	`
	fmt.Fprintf(response, responseString)
}
