package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Webserver setup")

	router := mux.NewRouter()

	router.HandleFunc("/shortenurl", urlShorten).Methods("POST")
	router.HandleFunc("/getlinks", getLinks).Methods("GET")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}

func getLinks(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "List of all Links")
}

func urlShorten(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "urlShorten")
}
