package steelseries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	http "github.com/bogdanfinn/fhttp"
)

func (in *Instance) Register() error {
	Payload := &RegisterPayload{
		Email:                 in.Email,
		Password1:             in.Password,
		Password2:             in.Password,
		AcceptedPrivacyPolicy: true,
		SubscribeToNewsletter: false,
	}

	buffer := new(bytes.Buffer)

	json.NewEncoder(buffer).Encode(Payload)

	req, err := http.NewRequest("POST", "https://steelseries.com/api/v2/users", buffer)

	if err != nil {
		return err
	}

	req = in.SetHeaders(true, map[string]string{
		"content-type": "application/json",
	}, req)

	resp, err := in.Client.Do(req)

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(bodyText))

	if err != nil {
		return err
	}

	if !Ok(resp.StatusCode) {
		return fmt.Errorf("could not register account, status code: %v", resp.StatusCode)
	}

	return nil
}
