package steelseries

import (
	http "github.com/bogdanfinn/fhttp"
)

func (in *Instance) SetHeaders(useCommonHeaders bool, headers map[string]string, req *http.Request) *http.Request {
	if useCommonHeaders {
		for k, v := range map[string]string{
			"host":            "steelseries.com",
			"user-agent":      in.SteelSeriesAgent,
			"authorization":   "Basic OThlNTQ1NWE5ZjY4MGM5YTVmZDcwNjo=",
			"accept-encoding": "gzip",
		} {
			req.Header.Set(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req
}

func (in *Instance) EmailHeaders(useCommonHeaders bool, headers map[string]string, req *http.Request) *http.Request {
	if useCommonHeaders {
		for k, v := range map[string]string{
			"authority":                 "emailtemp.org",
			"accept":                    "*/*",
			"accept-language":           "en-CA,en;q=0.9",
			"dnt":                       "1",
			"sec-ch-ua":                 `"Chromium";v="103", "Google Chrome";v="103", "Not;A=Brand";v="99"`,
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        `"Windows"`,
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
			"user-agent":                in.UserAgent,
		} {
			req.Header.Set(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req
}

func Ok(statusCode int) bool {
	for _, v := range []int{200, 201, 204} {
		if v == statusCode {
			return true
		}
	}

	return false
}
