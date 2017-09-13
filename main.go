package main

import(
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/subosito/gotenv"
    "github.com/ramsgoli/whether/geoloc"
    "time"
)

const (
    API_URL string = "api.openweathermap.org"
)

type Weather struct {
    Name string `json:"name"`
    Main struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
}

func init() {
    gotenv.Load()
}

func main() {

    lat, long := geoloc.Locate(os.Getenv("GOOGLE_MAPS_API_KEY")
    url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, os.Getenv("OWM_APP_ID"))

    // We use http.Client to have more control over headers, redirect policy, etc
    client := http.Client{
        Timeout: time.Second * 2, // set a timeout of two seconds for the api call
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Do sends an http request and returns an http response
    res, getErr := client.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    // defer the closing of the res body
    defer res.Body.Close()

    // read the http response body into a byte stream
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    weather := Weather{}

    // unmarshal the byte stream into a Go data type
    jsonErr := json.Unmarshal(body, &weather)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    fmt.Printf("%s: %f\n", weather.Name, weather.Main.Temp)
}
