package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, x)
}
