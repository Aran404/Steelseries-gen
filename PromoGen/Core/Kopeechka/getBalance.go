package Kopeechka

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (in *Instance) GetBalance() (float64, error) {
	resp, err := in.Client.Get(fmt.Sprintf("https://api.kopeechka.store/user-balance?token=%s&type=json&api=2.0", in.ApiKey))

	if err != nil {
		return -1, err
	}

	var jsonContent map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonContent)

	if err != nil {
		return -1, err
	}

	if jsonContent["status"] == "OK" {
		return jsonContent["balance"].(float64), nil
	} else {
		return -1, errors.New(jsonContent["value"].(string))
	}
}
