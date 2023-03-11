package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/surajn222/url-shortener/pkg/shortener"
	"github.com/surajn222/url-shortener/pkg/storage"
)

func main() {

	fmt.Println("Webserver setup")

	router := mux.NewRouter()

	router.HandleFunc("/shortenurl/{url}", muxUrlShorten).Methods("GET")
	router.HandleFunc("/getlinks", muxGetLinks).Methods("GET")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}

func muxGetLinks(response http.ResponseWriter, request *http.Request) {
	allLinks := storage.GetAllLinks()
	fmt.Println(allLinks)
	mapJson, err := json.Marshal(allLinks)
	fmt.Println(err)
	fmt.Fprintf(response, string(mapJson))
}

func muxUrlShorten(response http.ResponseWriter, request *http.Request) {
	code := mux.Vars(request)["url"]
	shortenedUrl := shortener.UrlShorten(code)
	resString := code + ":" + shortenedUrl
	storage.StoreShortenedLinks(resString, shortenedUrl)
	fmt.Fprintf(response, resString)
}
