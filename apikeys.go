package main

import "os"

func apiKeysPresent() bool {
    if (os.Getenv("OWM_APP_ID") == "" || os.Getenv("GOOGLE_MAPS_API_KEY") == "") {
		return false
    }
    return true
}
