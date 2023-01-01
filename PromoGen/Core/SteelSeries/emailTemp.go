package steelseries

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
)

func (in *Instance) GetPage() error {
	req, err := http.NewRequest("GET", "https://emailtemp.org/en", nil)

	if err != nil {
		return err
	}

	req = in.EmailHeaders(true, nil, req)

	resp, err := in.Client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	in.Csrf = strings.Split(strings.Split(string(bodyText), `"csrf-token" content="`)[1], `">`)[0]

	return nil
}

func (in *Instance) GetNewEmail() (string, error) {
	var data = strings.NewReader(fmt.Sprintf(`_token=%s&captcha=`, in.Csrf))
	req, err := http.NewRequest("POST", "https://emailtemp.org/messages", data)

	if err != nil {
		return "", err
	}

	req = in.EmailHeaders(true, map[string]string{
		"content-type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"origin":           "https://emailtemp.org",
		"referer":          "https://emailtemp.org/en",
		"sec-fetch-dest":   "empty",
		"sec-fetch-mode":   "cors",
		"sec-fetch-site":   "same-origin",
		"x-requested-with": "XMLHttpRequest",
	}, req)

	resp, err := in.Client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var TempMailResponse struct {
		Mailbox  string        `json:"mailbox,omitempty"`
		Messages []interface{} `json:"messages,omitempty"`
	}

	err = json.Unmarshal(bodyText, &TempMailResponse)

	if err != nil {
		return "", err
	}

	in.Email = TempMailResponse.Mailbox
	return TempMailResponse.Mailbox, nil
}

// Timeout is iterations, timeout=60, tries for 60 times (60 seconds)
func (in *Instance) WaitForMail(timeout int, parseMailVerification bool) (interface{}, error) {
	for i := 0; i < timeout; i++ {
		var data = strings.NewReader(fmt.Sprintf(`_token=%s&captcha=`, in.Csrf))
		req, err := http.NewRequest("POST", "https://emailtemp.org/messages", data)

		if err != nil {
			return "", err
		}

		req = in.EmailHeaders(true, map[string]string{
			"content-type":     "application/x-www-form-urlencoded; charset=UTF-8",
			"origin":           "https://emailtemp.org",
			"referer":          "https://emailtemp.org/en",
			"sec-fetch-dest":   "empty",
			"sec-fetch-mode":   "cors",
			"sec-fetch-site":   "same-origin",
			"x-requested-with": "XMLHttpRequest",
		}, req)

		resp, err := in.Client.Do(req)

		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		var TempMailResponse struct {
			Mailbox  string `json:"mailbox"`
			Messages []struct {
				Subject     string        `json:"subject"`
				IsSeen      bool          `json:"is_seen"`
				From        string        `json:"from"`
				FromEmail   string        `json:"from_email"`
				ReceivedAt  string        `json:"receivedAt"`
				ID          string        `json:"id"`
				Attachments []interface{} `json:"attachments"`
				Content     string        `json:"content"`
			} `json:"messages"`
		}

		err = json.Unmarshal(bodyText, &TempMailResponse)

		if err != nil {
			return "", err
		}

		if !parseMailVerification && len(TempMailResponse.Messages) > 0 {
			return TempMailResponse.Messages[0], nil
		} else if parseMailVerification && len(TempMailResponse.Messages) > 0 {
			re := regexp.MustCompile(`http:\/\/link\.steelseries\.com\/ls\/click\?upn=([a-zA-z0-9-]+)`)
			return re.FindAllString(fmt.Sprintf("%v", TempMailResponse.Messages[0]), -1)[1], nil
		} else {
			time.Sleep(1 * time.Second)
			continue
		}
	}

	return "", nil
}
