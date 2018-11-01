package opsman

import "net/http"

// AuthConfig used by the auth-file command
type AuthConfig struct {
	UserName             string `yaml:"username,omitempty"`
	Password             string `yaml:"password,omitempty"`
	DecryptionPassphrase string `yaml:"decryption-passphrase"`
	IDPMetadata          string `yaml:"saml-idp-metadata,omitempty"`
	BOSHIDPMetadata      string `yaml:"saml-bosh-idp-metadata,omitempty"`
	RBACAdminGroup       string `yaml:"saml-rbac-admin-group,omitempty"`
	RBACGroupsAttribute  string `yaml:"saml-rbac-groups-attribute,omitempty"`
	HTTPProxyURL         string `yaml:"http-proxy-url,omitempty"`
	HTTPSProxyURL        string `yaml:"https-proxy-url,omitempty"`
	NoProxy              string `yaml:"no-proxy,omitempty"`
}

// EnvConfig used by the env-file command
type EnvConfig struct {
	Target               string `yaml:"target"`
	SkipSSLValidation    bool   `yaml:"skip-ssl-validation"`
	UserName             string `yaml:"username,omitempty"`
	Password             string `yaml:"password,omitempty"`
	ConnectTimeout       int    `yaml:"connect-timeout,omitempty"`
	RequestTimeout       int    `yaml:"request-timeout,omitempty"`
	ClientID             string `yaml:"client-id,omitempty"`
	ClientSecret         string `yaml:"client-secret,omitempty"`
	Trace                bool   `yaml:"trace"`
	DecryptionPassphrase string `yaml:"decryption-passphrase"`
}

// OmAuthConfig used by the deprecated opsman-auth command
type OmAuthConfig struct {
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

// OmEnvConfig used by the deprecated om-env-file command
type OmEnvConfig struct {
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
