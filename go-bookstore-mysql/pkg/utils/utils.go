package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { //ReadAll reads from r until an error or EOF and returns the data it read
		if err := json.Unmarshal([]byte(body), X); err != nil { //Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by X.
			return
		}
	}
}
