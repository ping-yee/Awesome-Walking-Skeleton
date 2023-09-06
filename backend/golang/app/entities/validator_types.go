package entities

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DecodeAndValidate(r *http.Request, v InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(v); err != nil {
		return err
	}

	defer r.Body.Close()
	// peform validation on the InputValidation implementation
	return v.Validate()
}
