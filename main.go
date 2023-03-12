package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/surajn222/url-shortener/pkg/controllers"
)

func main() {
	// Main function
	logrus.Info("WebServer setup")

	router := mux.NewRouter()

	// Define Subrouters
	s1 := router.PathPrefix("/getlinks").Subrouter()
	s1.HandleFunc("", controllers.MuxGetLinks)

	s2 := router.PathPrefix("/shortenurl").Subrouter()
	s2.HandleFunc("", controllers.MuxUrlShorten)

	s3 := router.PathPrefix("/domaincount").Subrouter()
	s3.HandleFunc("", controllers.MuxDomainCount)

	s4 := router.PathPrefix("/").Subrouter()
	s4.HandleFunc("/index.html", controllers.MuxIndex)

	s5 := router.PathPrefix("/").Subrouter()
	s5.HandleFunc("/{*}", controllers.MuxRedirect)

	// Start server
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		logrus.Fatalln("There's an error with the server ", err)
	}

}
