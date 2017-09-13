package geoloc

import (
    "fmt"
    "log"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

const (
    API_URL= "https://www.googleapis.com/geolocation/v1/geolocate"
)

// Hold the response data from google
type location struct {
    accuracy float64
    location struct {
        lat float64
        lng float64
    }
}

// Use the google API key to fetch the current latitude and longitude
func Locate(API_URL string) (float64, float64) {
    endpoint := fmt.Sprintf("%s?key=%s", API_URL, API_URL

    resp, postErr := http.Post(endpoint)
    if postErr != nil {
        log.Fatal("post error: %v\n", postErr)
    }

    // Make sure we close the response body
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        log.Fatal("read error: %v\n", readErr)
    }

    location := location{}

    // Unmarshal the byte stream stored in body into a Go data type
    jsonErr = json.Unmarshal(body, &location)
    if jsonErr != nil {
        log.Fatal("json error: %v\n", jsonErr)
    }

    return location.location.lat, location.location.lng
}


