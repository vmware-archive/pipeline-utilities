package commands

type UAAConfig struct {
	Target                     string `yaml:"target"`
	SkipSSLValidation          bool   `yaml:"skip_ssl_validation"`
	Verbose                    bool   `yaml:"verbose"`
	Deployment                 string `yaml:"deployment"`
	UAAAdminCredentialProperty string `yaml:"uaa_admin_credential_property"`
	Clients                    map[string]struct {
		Secret              string   `yaml:"secret"`
		GrantTypes          []string `yaml:"grant_types"`
		Scope               []string `yaml:"scope"`
		Authorities         []string `yaml:"authorities"`
		AccessTokenValidity int64    `yaml:"access_token_validity"`
	} `yaml:"clients"`
}
