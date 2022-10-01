package handler

import (
	"encoding/json"
)

func JsonError(msg string) [] byte {
	error := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}

	json, err := json.Marshal(error)

	if err != nil {
		return []byte(err.Error())
	}

	return json
}