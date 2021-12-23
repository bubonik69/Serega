package API

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)
type APOD struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}
var recordings []APOD

func getRecord (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recordings)
}


func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record APOD
	_ = json.NewDecoder(r.Body).Decode(&record)
	record.Date = strconv.Itoa(rand.Intn(1000000))
	recordings = append(recordings, record)
	json.NewEncoder(w).Encode(record)
}
