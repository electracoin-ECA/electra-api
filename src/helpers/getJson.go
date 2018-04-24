package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// GetJSON fetches a JSON response from the given <url>.
func GetJSON(url string, target interface{}) bool {
	r, err := http.Get(url)
	if err != nil {
		log.Println("helpers/GetJSON(): Error feching url.", err)

		return true
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Println("helpers/GetJSON(): The response status is not 200 (" + strconv.Itoa(r.StatusCode) + ").")

		return true
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("helpers/GetJSON(): Error reading body.", err)

		return true
	}

	if err := json.Unmarshal(body, &target); err != nil {
		log.Println("helpers/GetJSON(): Error decoding JSON.", err)

		return true
	}

	return false
}
