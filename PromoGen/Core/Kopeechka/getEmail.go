package Kopeechka

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (in *Instance) GetEmail(domain string) (string, error) {
	if strings.ToLower(domain) == "hotmail" {
		domain = "OUTLOOK"
	}

	resp, err := in.Client.Get(fmt.Sprintf("https://api.kopeechka.store/mailbox-get-email?site=steelseries.com&mail_type=%s&token=%s&api=2.0", domain, in.ApiKey))

	if err != nil {
		return "", err
	}

	var jsonContent map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonContent)

	if err != nil {
		return "", err
	}

	if jsonContent["status"] == "OK" {
		id := jsonContent["id"].(string)
		email := jsonContent["mail"].(string)

		if err != nil {
			return "", err
		}

		in.Email = email
		in.ID = id
		return fmt.Sprintf("%s|%s", email, id), nil

	}
	return "", fmt.Errorf(fmt.Sprintf("could not get email, error: %s", jsonContent["value"].(string)))
}
