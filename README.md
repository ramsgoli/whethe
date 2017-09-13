# Whether

a command line tool to fetch the weather, built with love and go

### Setup
This tool depends on Google Maps Api to fetch your current geo-coordinates, and OpenWeatherMap to fetch the weather. 
To use this tool, you need to sign up for a Google Maps API key [here](https://developers.google.com/maps/documentation/javascript/get-api-key) and an OpenWeatherMap app [here](https://openweathermap.org/price). If you're my friend, you can ask me for mine. 

* create an environment file `.env` to hold your Google API key and OWM APP id
```
export GOOGLE_MAPS_API_KEY={your API key}
export OWM_APP_ID={your App id}
```
* get the source code (via go)
```
$ go get github.com/ramsgoli/whether
```

### Usage
```
# get the weather based off of your current location
$ whether

# get the weather in a certain city (for a city with multiple words, wrap the name in quotes)
$ whether -city petaluma
>>> The current temperature in Petaluma is 65.43 degrees
$ whether -city "los angeles"
>>> The current temperature in Los Angeles is 70.32 degrees
```

