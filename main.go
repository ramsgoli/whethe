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
    "flag"
    "time"
)

const (
    API_URL string = "api.openweathermap.org"
)

var (
    city string
)

type Weather struct {
    Name string `json:"name"`
    Id int `json:"id"`
    Dt int `json:"dt"`
    Clouds struct {
	All int `json:"all"`
    } `json:"clouds"`
    Main struct {
	Temp float64 `json:"temp"`
    } `json:"main"`
}

func apiKeysPresent() bool {
    if (os.Getenv("OWM_APP_ID") == "" || os.Getenv("GOOGLE_MAPS_API_KEY") == "") {
	return false
    }
    return true
}

func init() {
    // Load environment variables
    gotenv.Load()

    flag.StringVar(&city, "city", "", "City to get the weather from")
    flag.StringVar(&city, "c", "", "(alias) City to get the weather from")

    flag.Usage = func() {
	flag.PrintDefaults()
    }

    flag.Parse()
}

func main() {

    // check if api keys are present before attempting to make request
    if !apiKeysPresent() {
		fmt.Println("\nNo api keys present\n")
		os.Exit(1)
    }

    var url string
    if city == "" {
		// get latitude and longitude of client
		lat, long, err := geoloc.Locate(os.Getenv("GOOGLE_MAPS_API_KEY"))
		if err != nil {
			log.Fatal(err)
		}

		url = fmt.Sprintf("http://%s/data/2.5/weather?lat=%f&lon=%f&units=imperial&APPID=%s", API_URL, lat, long, os.Getenv("OWM_APP_ID"))
    } else {
		// use the city
		url = fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, os.Getenv("OWM_APP_ID"))
    }

    // Build an http client so we can have control over timeout
    client := &http.Client{
		Timeout: time.Second * 2,
    }

    res, getErr := client.Get(url)
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

    fmt.Printf("The current temperature in %s is %.2f degrees\n", weather.Name, weather.Main.Temp)
}
