package commands

import (
	"github.com/pivotalservices/pipeline-utilities/common"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

// CreateAuthFile is a auth.yml struct as defined by p-automator
// http://docs-platform-automation.cfapps.io/pcf-automation/task-reference.html#auth
type CreateAuthFile struct {
	Username             string             `long:"username" env:"OPSMAN_USERNAME" description:"OpsManager username"`
	Password             string             `long:"password" env:"OPSMAN_PASSWORD" description:"OpsManager password"`
	DecryptionPassphrase string             `long:"decryption-passphrase" env:"OPSMAN_DECRYPTION_PASSPHRASE" description:"OpsManager Decryption Passphrase"`
	SAMLConfiguration    *SAMLConfiguration `group:"SAML"`
	HTTPProxyURL         string             `long:"http-proxy-url" env:"OPSMAN_HTTP_PROXY_URL" description:"proxy for outbound HTTP network traffic"`
	HTTPSProxyURL        string             `long:"https-proxy-url" env:"OPSMAN_HTTPS_PROXY_URL" description:"proxy for outbound HTTPS network traffic"`
	NoProxy              string             `long:"no-proxy"    env:"OPSMAN_NO_PROXY" description:"comma-separated list of hosts that do not go through the proxy"`
	OutputFile           string             `long:"output-file" description:"output file for yaml" default:"auth.yml"`
}

// Execute - creates om env
func (c *CreateAuthFile) Execute([]string) error {
	authFileConfig := opsman.AuthConfig{
		DecryptionPassphrase: c.DecryptionPassphrase,
		UserName:             c.Username,
		Password:             c.Password,
		HTTPProxyURL:         c.HTTPProxyURL,
		HTTPSProxyURL:        c.HTTPSProxyURL,
		NoProxy:              c.NoProxy,
	}

	if c.SAMLConfiguration != nil {
		authFileConfig.IDPMetadata = c.SAMLConfiguration.IDPMetadata
		authFileConfig.BOSHIDPMetadata = c.SAMLConfiguration.BOSHIDPMetadata
		authFileConfig.RBACAdminGroup = c.SAMLConfiguration.RBACAdminGroup
		authFileConfig.RBACGroupsAttribute = c.SAMLConfiguration.RBACGroupsAttribute
	}

	return common.WriteYamlFile(c.OutputFile, &authFileConfig)

}
