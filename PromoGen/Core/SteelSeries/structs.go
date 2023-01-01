package steelseries

import (
	tls_client "github.com/bogdanfinn/tls-client"
)

type Config struct {
	Threads        int  `json:"threads"`
	Iterations     int  `json:"iterations"`
	Proxyless      bool `json:"proxyless"`
	Fingerprinting struct {
		UserAgent     string `json:"userAgent"`
		ChromeVersion string `json:"chromeVersion"`
		ClientTimeout int    `json:"clientTimeout"`
	} `json:"fingerprinting"`
}

type Instance struct {
	AccessToken      string
	UserAgent        string
	Email            string
	SteelSeriesAgent string
	PromoCode        string
	Password         string
	Csrf             string
	VerificationLink string
	Client           tls_client.HttpClient
}

type RegisterPayload struct {
	Email                 string `json:"email"`
	Password1             string `json:"password1"`
	Password2             string `json:"password2"`
	AcceptedPrivacyPolicy bool   `json:"accepted_privacy_policy"`
	SubscribeToNewsletter bool   `json:"subscribe_to_newsletter"`
}

type RegisterResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type GetCodeResponse struct {
	PromoCode string `json:"promo_code_url"`
}

type VerifyEmailPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
