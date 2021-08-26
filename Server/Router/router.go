package Router

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"userdetection/Configuration"

	"github.com/gorilla/mux"
)

func registerRoutes(clientBuildPath string) (router *mux.Router) {
	router = mux.NewRouter()

	// Get user tweets route
	router.HandleFunc("/tweets", GetTweets).Methods(http.MethodPost)
	// Get top dangerous users route
	router.HandleFunc("/dangerous-users", GetTopDangerousUsers).Methods(http.MethodGet)

	// Static files handler
	spa := spaHandler{staticPath: clientBuildPath, indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	return
}

// InitServer init server router
func InitServer() {
	router := registerRoutes(Configuration.Conf.ClientFolder)
	const address = "127.0.0.1:8080"

	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server is listening on", address)
	log.Fatal(server.ListenAndServe())
}
