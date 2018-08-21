package commands

import (
	"github.com/pivotalservices/pipeline-utilities/common"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

type CreateOMEnvironmentFile struct {
	Target               string             `long:"target" env:"OPSMAN_TARGET" description:"OpsManager hostname" required:"true"`
	SkipSSLValidation    bool               `long:"skip-ssl-validation" env:"OPSMAN_SKIP_SSL_VALIDATION" description:"Skip SSL Validation when interacting with OpsManager"`
	Username             string             `long:"username" env:"OPSMAN_USERNAME" description:"OpsManager username"`
	Password             string             `long:"password" env:"OPSMAN_PASSWORD" description:"OpsManager password"`
	ClientID             string             `long:"client-id" env:"OPSMAN_CLIENTID" description:"OpsManager client id"`
	ClientSecret         string             `long:"client-secret" env:"OPSMAN_CLIENT_SECRET" description:"OpsManager client secret"`
	DecryptionPassphrase string             `long:"decryption-passphrase" env:"OPSMAN_DECRYPTION_PASSPHRASE" description:"OpsManager Decryption Passphrase"`
	ConnectTimeout       int                `long:"connect-timeout" env:"OPSMAN_CONNECT_TIMEOUT" description:"OpsManager Connect timeout" default:"5"`
	RequestTimeout       int                `long:"request-timeout" env:"OPSMAN_REQUEST_TIMEOUT" description:"OpsManager Request timeout" default:"1800"`
	SAMLConfiguration    *SAMLConfiguration `group:"SAML"`
	OutputFile           string             `long:"output-file" description:"output file for yaml" required:"true"`
}

type SAMLConfiguration struct {
	IDPMetadata         string `long:"idp-metadata" env:"OPSMAN_SAML_IDP_METADATA" description:"OpsManager SAML IDP Metadata"`
	BOSHIDPMetadata     string `long:"bosh-idp-metadata" env:"OPSMAN_SAML_BOSH_IDP_METADATA" description:"OpsManager SAML BOSH IDP Metadata"`
	RBACAdminGroup      string `long:"rbac-admin-group" env:"OPSMAN_RBAC_ADMIN_GROUP" description:"OpsManager RBAC admin group"`
	RBACGroupsAttribute string `long:"rbac-groups-attribute" env:"OPSMAN_RBAC_GROUPS_ATTRIBUTE" description:"OpsManager RBAC groups attribute"`
}

//Execute - creates om env
func (c *CreateOMEnvironmentFile) Execute([]string) error {
	envConfig := opsman.ENVConfig{
		Target:               c.Target,
		SkipSSLValidation:    c.SkipSSLValidation,
		DecryptionPassphrase: c.DecryptionPassphrase,
	}

	envConfig.UserName = c.Username
	envConfig.Password = c.Password
	envConfig.ClientID = c.ClientID
	envConfig.ClientSecret = c.ClientSecret
	envConfig.ConnectTimeout = c.ConnectTimeout
	envConfig.RequestTimeout = c.RequestTimeout

	if c.SAMLConfiguration != nil {
		envConfig.SAMLConfiguration.IDPMetadata = c.SAMLConfiguration.IDPMetadata
		envConfig.SAMLConfiguration.BOSHIDPMetadata = c.SAMLConfiguration.BOSHIDPMetadata
		envConfig.SAMLConfiguration.RBACAdminGroup = c.SAMLConfiguration.RBACAdminGroup
		envConfig.SAMLConfiguration.RBACGroupsAttribute = c.SAMLConfiguration.RBACGroupsAttribute
	}

	return common.WriteYamlFile(c.OutputFile, &envConfig)

}
