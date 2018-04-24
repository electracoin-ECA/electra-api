package helpers

import (
	"log"
	"strconv"
)

// StringToFloat converts a string to a float64.
func StringToFloat(s string) float64 {
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		log.Println("helpers/StringToFloat(): Error converting string to float64.", e)

		return 0
	}

	return r
}
