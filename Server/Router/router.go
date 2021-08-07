package Router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func registerRoutes(clientBuildPath string) (router *mux.Router) {
	router = mux.NewRouter()

	// an example API handler
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
	// static files handler
	spa := spaHandler{staticPath: clientBuildPath, indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	return
}

func InitServer() {
	clientBuildPath := `C:\Users\home\Desktop\UserDetection\Client\client\build`
	router := registerRoutes(clientBuildPath)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8999",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
