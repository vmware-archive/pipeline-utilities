package commands

import (
	"github.com/pivotalservices/pipeline-utilities/common"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

type CreateOpsmanAuth struct {
	URL                  string `long:"url" env:"OPSMAN_URL" description:"OpsManager URL" required:"true"`
	SkipSSLValidation    bool   `long:"skip-ssl-validation" env:"OPSMAN_SKIP_SSL_VALIDATION" description:"Skip SSL Validation when interacting with OpsManager"`
	Username             string `long:"username" env:"OPSMAN_USERNAME" description:"OpsManager username"`
	Password             string `long:"password" env:"OPSMAN_PASSWORD" description:"OpsManager password"`
	ClientID             string `long:"client-id" env:"OPSMAN_CLIENTID" description:"OpsManager client id"`
	ClientSecret         string `long:"client-secret" env:"OPSMAN_CLIENT_SECRET" description:"OpsManager client secret"`
	DecryptionPassphrase string `long:"decryption-passphrase" env:"OPSMAN_DECRYPTION_PASSPHRASE" description:"OpsManager Decryption Passphrase"  required:"true"`
	OutputFile           string `long:"output-file" description:"output file for yaml" required:"true"`
}

//Execute - generates structs
func (c *CreateOpsmanAuth) Execute([]string) error {
	authConfig := opsman.AuthConfig{
		OpsmanURL:            c.URL,
		SkipSSLValidation:    c.SkipSSLValidation,
		DecryptionPassphrase: c.DecryptionPassphrase,
	}

	authConfig.Credentials.UserName = c.Username
	authConfig.Credentials.Password = c.Password
	authConfig.Credentials.ClientID = c.ClientID
	authConfig.Credentials.ClientSecret = c.ClientSecret

	return common.WriteYamlFile(c.OutputFile, &authConfig)
}
