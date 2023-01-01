package steelseries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (in *Instance) Login() error {
	data := url.Values{}

	data.Add("grant_type", "password")
	data.Add("username", in.Email)
	data.Add("password", in.Password)

	req, err := http.NewRequest("POST", "https://steelseries.com/oauth2/token", strings.NewReader(data.Encode()))

	if err != nil {
		return err
	}

	req = in.SetHeaders(true, map[string]string{
		"content-type": "application/x-www-form-urlencoded",
	}, req)

	resp, err := in.Client.Do(req)

	if err != nil {
		return err
	}

	if !Ok(resp.StatusCode) {
		return fmt.Errorf("could not register account, status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var Response LoginResponse

	json.Unmarshal(bodyText, &Response)

	if len(Response.AccessToken) <= 0 {
		return fmt.Errorf("access token was not returned")
	}

	in.AccessToken = Response.AccessToken

	return nil
}
