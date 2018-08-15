package opsman

import "net/http"

type AuthConfig struct {
	OpsmanURL         string `yaml:"opsman_url"`
	SkipSSLValidation bool   `yaml:"skip_ssl_validation"`
	Credentials       struct {
		UserName     string `yaml:"username,omitempty"`
		Password     string `yaml:"password,omitempty"`
		ClientID     string `yaml:"client-id,omitempty"`
		ClientSecret string `yaml:"client-secret,omitempty"`
	} `yaml:"credentials,omitempty"`
	DecryptionPassphrase string `yaml:"decryption_passphrase"`
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}
