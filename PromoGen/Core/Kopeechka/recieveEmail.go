package Kopeechka

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (in *Instance) GetVerificationEmail() (string, error) {
	resp, err := in.Client.Get(fmt.Sprintf("https://api.kopeechka.store/mailbox-get-message?full=1&spa=1&id=%s&token=%s&api=2.0", in.ID, in.ApiKey))

	if err != nil {
		return "", err
	}

	var jsonContent map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonContent)

	if err != nil {
		return "", err
	}

	if jsonContent["status"] == "OK" {
		return jsonContent["value"].(string), nil
	}
	
	return "", errors.New(jsonContent["value"].(string))
}
