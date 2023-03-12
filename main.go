package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/surajn222/url-shortener/pkg/shortener"
	"github.com/surajn222/url-shortener/pkg/storage"
)

func main() {

	logrus.Info("WebServer setup")

	router := mux.NewRouter()

	// router.HandleFunc("/**", redir).Methods("GET")

	// router.HandleFunc("/shortenurl/{url}", muxUrlShorten).Methods("GET")
	// router.HandleFunc("/getlinks", muxGetLinks).Methods("GET")
	// router.PathPrefix("").HandlerFunc(redir).Methods("GET")

	s1 := router.PathPrefix("/getlinks").Subrouter()
	s1.HandleFunc("", muxGetLinks)

	s2 := router.PathPrefix("/shortenurl").Subrouter()
	s2.HandleFunc("", muxUrlShorten)

	s3 := router.PathPrefix("/").Subrouter()
	s3.HandleFunc("/{*}", redir)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		logrus.Fatalln("There's an error with the server ", err)
	}

}

func muxGetLinks(response http.ResponseWriter, request *http.Request) {
	allLinks := storage.GetAllLinks()
	logrus.Println(allLinks)
	mapJson, err := json.Marshal(allLinks)
	logrus.Println(err)
	fmt.Fprintf(response, string(mapJson))
}

func muxUrlShorten(response http.ResponseWriter, request *http.Request) {
	logrus.Printf("%+v", request)
	// path := mux.Vars(request)["URL"]
	// path := request.URL.String()
	path := request.URL.Query().Get("url")
	logrus.Println(path)
	shortenedUrl := shortener.UrlShorten(path)
	resString := path + ":" + shortenedUrl
	storage.StoreShortenedLinks(resString, shortenedUrl)
	fmt.Fprintf(response, resString)
}

func redir(response http.ResponseWriter, request *http.Request) {
	logrus.Info("HERE-----")
	logrus.Info(request.URL.Path)
	mapShortUrls := storage.GetAllLinks()
	logrus.Infof("%+v", mapShortUrls)
	_, ok := mapShortUrls[request.URL.Path[1:]]
	// If the key exists
	if ok {
		logrus.Info("Not found-----")
		http.Redirect(response, request, "https://google.com", http.StatusMovedPermanently)
	} else {
		logrus.Info("Not found-----")
		fmt.Fprintf(response, "404 Not Found")
	}

}
