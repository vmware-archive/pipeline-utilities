package commands

import (
	"fmt"

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

// SAMLConfiguration used to configure SAML auth on OpsMan
type SAMLConfiguration struct {
	IDPMetadata         string `long:"idp-metadata" env:"OPSMAN_SAML_IDP_METADATA" description:"OpsManager SAML IDP Metadata"`
	BOSHIDPMetadata     string `long:"bosh-idp-metadata" env:"OPSMAN_SAML_BOSH_IDP_METADATA" description:"OpsManager SAML BOSH IDP Metadata"`
	RBACAdminGroup      string `long:"rbac-admin-group" env:"OPSMAN_RBAC_ADMIN_GROUP" description:"OpsManager RBAC admin group"`
	RBACGroupsAttribute string `long:"rbac-groups-attribute" env:"OPSMAN_RBAC_GROUPS_ATTRIBUTE" description:"OpsManager RBAC groups attribute"`
}

//Execute - creates om env
func (c *CreateOMEnvironmentFile) Execute([]string) error {
	fmt.Println("******* WARNING DEPRECATED COMMAND - use env-file and auth-file ************")
	omEnvConfig := opsman.OmEnvConfig{
		Target:               c.Target,
		SkipSSLValidation:    c.SkipSSLValidation,
		DecryptionPassphrase: c.DecryptionPassphrase,
	}

	omEnvConfig.UserName = c.Username
	omEnvConfig.Password = c.Password
	omEnvConfig.ClientID = c.ClientID
	omEnvConfig.ClientSecret = c.ClientSecret
	omEnvConfig.ConnectTimeout = c.ConnectTimeout
	omEnvConfig.RequestTimeout = c.RequestTimeout

	if c.SAMLConfiguration != nil {
		omEnvConfig.SAMLConfiguration.IDPMetadata = c.SAMLConfiguration.IDPMetadata
		omEnvConfig.SAMLConfiguration.BOSHIDPMetadata = c.SAMLConfiguration.BOSHIDPMetadata
		omEnvConfig.SAMLConfiguration.RBACAdminGroup = c.SAMLConfiguration.RBACAdminGroup
		omEnvConfig.SAMLConfiguration.RBACGroupsAttribute = c.SAMLConfiguration.RBACGroupsAttribute
	}

	return common.WriteYamlFile(c.OutputFile, &omEnvConfig)

}
