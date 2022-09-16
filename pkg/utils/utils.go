package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, w interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	json.Unmarshal([]byte(body), w)
}
