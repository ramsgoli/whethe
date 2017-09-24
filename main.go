package main

import(
    "fmt"
    "os"
    "log"
    "github.com/subosito/gotenv"
	"github.com/ramsgoli/openweathermap"
    "github.com/ramsgoli/whether/geoloc"
    "flag"
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

    owm := openweathermap.OpenWeatherMap{API_KEY: os.Getenv("OWM_APP_ID")}

    var currentWeather *openweathermap.CurrentWeatherResponse
	var weatherErr error

    if city == "" {
		// get latitude and longitude of client
		lat, long, geoErr := geoloc.Locate(os.Getenv("GOOGLE_MAPS_API_KEY"))
		if (geoErr != nil) {
			log.Fatal("Error with GeoLocating: ", geoErr)
		}

		currentWeather, weatherErr = owm.CurrentWeatherFromCoordinates(lat, long)
		if (weatherErr != nil) {
			log.Fatal("Error with CurrentWeatherFromCoordinates: ", weatherErr)
		}

    } else {
        // use the city
        currentWeather, weatherErr = owm.CurrentWeatherFromCity(city)
		if (weatherErr != nil) {
			log.Fatal("Error with CurrentWeatherFrom City: ", weatherErr)
        }
    }

    fmt.Printf("The current temperature in %s is %.2f degrees\n", currentWeather.Name, currentWeather.Main.Temp)
}
