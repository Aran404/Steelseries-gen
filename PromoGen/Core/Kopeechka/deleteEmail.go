package Kopeechka

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (in *Instance) DeleteEmail() (bool, error) {
	resp, err := in.Client.Get(fmt.Sprintf("https://api.kopeechka.store/user-balance?token=%s&id=%s&type=json&api=2.0", in.ApiKey, in.ID))

	if err != nil {
		return false, err
	}

	var jsonContent map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonContent)
	
	if err != nil {
		return false, err
	}

	if jsonContent["status"] == "OK" {
		return true, nil
	}

	return false, errors.New(jsonContent["value"].(string))
}
