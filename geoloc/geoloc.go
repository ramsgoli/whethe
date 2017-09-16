package geoloc

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

const (
    API_URL= "https://www.googleapis.com/geolocation/v1/geolocate"
)

// Hold the response data from google
type Location struct {
    Accuracy float64 `json:"accuracy"`
    Location struct {
        Lat float64 `json:"lat"`
        Lng float64 `json:"lng"`
    } `json:"location"`
}

// Use the google API key to fetch the current latitude and longitude
func Locate(API_KEY string) (float64, float64, error) {
    endpoint := fmt.Sprintf("%s?key=%s", API_URL, API_KEY)

    resp, postErr := http.Post(endpoint, "application/json", nil)
    if postErr != nil {
        return nil, nil, postErr
    }

    // Make sure we close the response body
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, nil, readErr
    }

    location := Location{}

    // Unmarshal the byte stream stored in body into a Go data type
    jsonErr := json.Unmarshal(body, &location)
    if jsonErr != nil {
        return nil, nil, jsonErr
    }

    return location.Location.Lat, location.Location.Lng, nil
}


