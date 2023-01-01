package client

import (
	"fmt"

	"github.com/bogdanfinn/fhttp/cookiejar"
	tls_client "github.com/bogdanfinn/tls-client"
)

func NewTlsClient(proxy, chromeProfile string, followRedirects bool, timeout int) (tls_client.HttpClient, error) {
	jar, err := cookiejar.New(nil)

	if err != nil {
		return nil, fmt.Errorf("creating cookiejar: %v", err)
	}

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(timeout),
		tls_client.WithCookieJar(jar),
		func() tls_client.HttpClientOption {
			switch chromeProfile {
			case "103":
				return tls_client.WithClientProfile(tls_client.Chrome_103)
			case "104":
				return tls_client.WithClientProfile(tls_client.Chrome_104)
			case "105":
				return tls_client.WithClientProfile(tls_client.Chrome_105)
			case "106":
				return tls_client.WithClientProfile(tls_client.Chrome_106)
			case "107":
				return tls_client.WithClientProfile(tls_client.Chrome_107)
			default:
				return tls_client.WithClientProfile(tls_client.DefaultClientProfile)
			}
		}(),
		tls_client.WithInsecureSkipVerify(),
	}

	if proxy != "" {
		options = append(options, tls_client.WithProxyUrl("http://"+proxy))
	} else if !followRedirects {
		options = append(options, tls_client.WithNotFollowRedirects())
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)

	if err != nil {
		return nil, fmt.Errorf("creating client: %v", err)
	}

	return client, nil
}
