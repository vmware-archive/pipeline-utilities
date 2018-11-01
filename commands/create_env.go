package commands

import (
	"github.com/pivotalservices/pipeline-utilities/common"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

// CreateEnvironmentFile is a env.yml struct that matches `om --env file` format
// https://github.com/pivotal-cf/om/blob/1540039512bdfd848af10d4baaadeb195a003008/main.go#L189-L239
// http://docs-platform-automation.cfapps.io/pcf-automation/task-reference.html#env
type CreateEnvironmentFile struct {
	Target               string `long:"target" env:"OPSMAN_TARGET" description:"OpsManager hostname" required:"true"`
	SkipSSLValidation    bool   `long:"skip-ssl-validation" env:"OPSMAN_SKIP_SSL_VALIDATION" description:"Skip SSL Validation when interacting with OpsManager"`
	Username             string `long:"username" env:"OPSMAN_USERNAME" description:"OpsManager username"`
	Password             string `long:"password" env:"OPSMAN_PASSWORD" description:"OpsManager password"`
	ClientID             string `long:"client-id" env:"OPSMAN_CLIENTID" description:"OpsManager client id"`
	ClientSecret         string `long:"client-secret" env:"OPSMAN_CLIENT_SECRET" description:"OpsManager client secret"`
	ConnectTimeout       int    `long:"connect-timeout" env:"OPSMAN_CONNECT_TIMEOUT" description:"OpsManager Connect timeout" default:"5"`
	RequestTimeout       int    `long:"request-timeout" env:"OPSMAN_REQUEST_TIMEOUT" description:"OpsManager Request timeout" default:"1800"`
	OutputFile           string `long:"output-file" description:"output file for yaml" default:"env.yml"`
	Trace                bool   `long:"trace" env:"OPSMAN_TRACE" description:"Prints HTTP requests and response payloads"`
	DecryptionPassphrase string `long:"decryption-passphrase" env:"OPSMAN_DECRYPTION_PASSPHRASE" description:"OpsManager Decryption passphrase"`
}

// Execute - creates the env.yml file
func (c *CreateEnvironmentFile) Execute([]string) error {
	envConfig := opsman.EnvConfig{
		Target:               c.Target,
		SkipSSLValidation:    c.SkipSSLValidation,
		UserName:             c.Username,
		Password:             c.Password,
		ClientID:             c.ClientID,
		ClientSecret:         c.ClientSecret,
		ConnectTimeout:       c.ConnectTimeout,
		RequestTimeout:       c.RequestTimeout,
		Trace:                c.Trace,
		DecryptionPassphrase: c.DecryptionPassphrase,
	}
	return common.WriteYamlFile(c.OutputFile, &envConfig)
}
