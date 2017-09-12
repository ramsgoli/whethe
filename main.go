package main

import(
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/subosito/gotenv"
    "time"
)

const (
    API_URL string = "api.openweathermap.org"
)

type Main struct {
    Temp float64 `json:"temp"`
}

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
    if len(os.Args) > 2 {
        fmt.Fprintf(os.Stderr, "command called with %d arguments, should be called with only 1\n", len(os.Args) - 1)
        os.Exit(1)
    }

    city := os.Args[1]
    url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, os.Getenv("APP_ID"))

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

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    weather := Weather{}

    jsonErr := json.Unmarshal(body, &weather)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    fmt.Printf("%s: %f\n", weather.Name, weather.Main.Temp)
}
