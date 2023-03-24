package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/surajn222/url-shortener/pkg/shortener"

	storage "github.com/surajn222/url-shortener/pkg/storage"
	"github.com/surajn222/url-shortener/pkg/utils"
	"github.com/surajn222/url-shortener/pkg/validations"
	"github.com/surajn222/url-shortener/templates/static"
)

func MuxDomainCount(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	storageObject := ctx.Value("storage").(storage.InterfaceStorage)
	jsonDomainCount := storageObject.GetDomainCount()
	mapDomainCount, err := json.Marshal(jsonDomainCount)

	if err != err {
		response.Write([]byte("Unable to get domain count from database"))
		response.WriteHeader(404)
	} else {
		response.Write([]byte(string(mapDomainCount)))
		response.WriteHeader(200)
	}

}

func MuxPath(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	storageObject := ctx.Value("storage").(storage.InterfaceStorage)
	jsonAllLinks := storageObject.GetAllLinks()
	mapAllLinks, err := json.Marshal(jsonAllLinks)
	if err != err {
		response.Write([]byte("Unable to get all links"))
		response.WriteHeader(404)
	} else {
		response.Write([]byte(string(mapAllLinks)))
		response.WriteHeader(200)
	}
}

func MuxGetLinks(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	storageObject := ctx.Value("storage").(storage.InterfaceStorage)
	jsonAllLinks := storageObject.GetAllLinks()
	mapAllLinks, err := json.Marshal(jsonAllLinks)
	if err != err {
		response.Write([]byte("Unable to get all links"))
		response.WriteHeader(404)
	} else {
		response.Write([]byte(string(mapAllLinks)))
		response.WriteHeader(200)
	}

}

func MuxUrlShorten(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	storageObject := ctx.Value("storage").(storage.InterfaceStorage)
	url := request.URL.Query().Get("url")

	// Validate if input is URL
	err := validations.UrlValidation(url)
	if err == nil {
		// Shorten URL with suitable algorithm
		shortenedUrl := shortener.UrlShorten(url)
		responseString := "{" + url + ":" + "http://localhost:8081/" + shortenedUrl + "}"

		// Store shortened links to database
		err1 := storageObject.InsertShortenedLinks(url, shortenedUrl)
		domainName := utils.GetDomainFromUrl(url)
		err2 := storageObject.UpdateDomainCount(domainName)
		if err2 != nil && err1 != nil {
			panic(err2)
			response.Write([]byte("Unable to shorten URL"))
			response.WriteHeader(404)
		}
		response.Write([]byte(string(responseString)))
		response.WriteHeader(200)
	} else {
		response.Write([]byte("Unable to shorten URL"))
		response.WriteHeader(404)
	}
}

func MuxRedirect(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	storageObject := ctx.Value("storage").(storage.InterfaceStorage)
	// Get link from database and redirect
	strShortenedUrl, err := storageObject.GetLink(request.URL.Path[1:])
	if err != nil {
		http.Redirect(response, request, "http://localhost:8081/index.html", http.StatusMovedPermanently)
	} else {
		http.Redirect(response, request, strShortenedUrl, http.StatusMovedPermanently)
	}
}

// func MuxPath(response http.ResponseWriter, request *http.Request) {
// 	// Get link from database and redirect
// 	fmt.Fprintf(response, "Path")

// }

func MuxIndex(response http.ResponseWriter, request *http.Request) {
	responseString := static.IndexResponse
	response.Write([]byte(string(responseString)))
	response.WriteHeader(200)
}
