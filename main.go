package main

import(
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/subosito/gotenv"
)

const (
    API_URL string = "api.openweathermap.org"
)

type Weather struct {
    name string
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

    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
    }

    b, err := ioutil.ReadAll(resp.Body) // b is a byte array
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error reading resp data: %v\n", err)
        os.Exit(1)
    }

    // Unmarshall byte array to a go data type
    var w Weather
    err = json.Unmarshal(b, &w)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error reading resp data: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(w)
}
