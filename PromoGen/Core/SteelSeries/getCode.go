package steelseries

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (in *Instance) GetCode() error {
	Payload := strings.NewReader("{}")

	req, err := http.NewRequest("POST", "https://steelseries.com/api/v1/promos/discordnitrodec2022", Payload)

	if err != nil {
		return err
	}

	req = in.SetHeaders(true, map[string]string{
		"content-type":  "application/json",
		"authorization": "Bearer " + in.AccessToken,
	}, req)

	fmt.Println(req)

	resp, err := in.Client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodt, _ := io.ReadAll(resp.Body)

	fmt.Println(string(bodt))

	if !Ok(resp.StatusCode) {
		return fmt.Errorf("could not get nitro code, status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var Response GetCodeResponse

	json.Unmarshal(bodyText, &Response)

	if len(Response.PromoCode) <= 0 {
		return fmt.Errorf("promo was not returned")
	}

	in.PromoCode = Response.PromoCode

	return nil
}
