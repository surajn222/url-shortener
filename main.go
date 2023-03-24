package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/surajn222/url-shortener/pkg/config"
	"github.com/surajn222/url-shortener/pkg/controllers"
	"github.com/surajn222/url-shortener/pkg/storage"
)

func getConfig() config.Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration

}

func main() {
	// Main function
	logrus.Info("WebServer Init")

	config := getConfig()
	storageObject := storage.GetStorageObject(config)

	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the database connection in the request context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "storage", storageObject)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/getlinks", controllers.MuxGetLinks).Methods("GET")

	router.HandleFunc("/shortenurl", controllers.MuxUrlShorten).Methods("GET")

	router.HandleFunc("/domaincount", controllers.MuxDomainCount).Methods("GET")

	router.HandleFunc("/{*}", controllers.MuxRedirect).Methods("GET")

	router.HandleFunc("/", controllers.MuxIndex).Methods("GET")

	// Start server
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalln("There's an error with the server ", err)
	}

}
