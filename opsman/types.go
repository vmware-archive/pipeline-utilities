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

type ENVConfig struct {
	Target               string `yaml:"target"`
	SkipSSLValidation    bool   `yaml:"skip-ssl-validation"`
	UserName             string `yaml:"username,omitempty"`
	Password             string `yaml:"password,omitempty"`
	ConnectTimeout       int    `yaml:"connect-timeout,omitempty"`
	RequestTimeout       int    `yaml:"request-timeout,omitempty"`
	ClientID             string `yaml:"client-id,omitempty"`
	ClientSecret         string `yaml:"client-secret,omitempty"`
	DecryptionPassphrase string `yaml:"decryption-passphrase,omitempty"`
	SAMLConfiguration    struct {
		IDPMetadata         string `yaml:"idp-metadata,omitempty"`
		BOSHIDPMetadata     string `yaml:"bosh-idp-metadata,omitempty"`
		RBACAdminGroup      string `yaml:"rbac-admin-group,omitempty"`
		RBACGroupsAttribute string `yaml:"rbac-groups-attribute,omitempty"`
	} `yaml:"saml-configuration,omitempty"`
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}
