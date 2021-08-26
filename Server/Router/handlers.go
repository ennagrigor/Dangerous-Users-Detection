package Router

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"userdetection/Database"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, path)
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	filter := new(Database.Filter)
	err := json.NewDecoder(r.Body).Decode(filter)
	if err != nil {
		httpError := GetHTTPError(BadRequest)
		http.Error(w, httpError.Message, httpError.ErrorCode)
		return
	}
	defer func() { _ = r.Body.Close() }()

	twits, err := Database.GetTweets(filter)
	if err != nil {
		httpError := GetHTTPError(ErrorTwitsNotFound)
		http.Error(w, httpError.Message, httpError.ErrorCode)
		return
	}

	err = json.NewEncoder(w).Encode(twits)
	if err != nil {
		log.Println("Failed encoding twits: ", err)
	}
	return
}

func GetTopDangerousUsers(w http.ResponseWriter, _ *http.Request) {
	dangerousUsers, err := Database.GetTopDangerousUsers()
	if err != nil {
		httpError := GetHTTPError(ErrorDangerousUsersNotFound)
		http.Error(w, httpError.Message, httpError.ErrorCode)
		return
	}

	err = json.NewEncoder(w).Encode(dangerousUsers)
	if err != nil {
		log.Println("Failed encoding dangerous users: ", err)
	}
	return
}
