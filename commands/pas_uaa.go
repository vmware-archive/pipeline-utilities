package commands

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

type PASUAA struct {
	AuthFile   string   `long:"auth-file" description:"path to auth file" default:"auth/auth.yml"`
	ConfigFile string   `long:"config" short:"c" required:"true" description:"path to config file"`
	VarsFile   []string `long:"vars-file" short:"l" description:"path to vars file"`
}

//Execute - creates/updates specified uaa client within PAS uaa
func (c *PASUAA) Execute([]string) error {

	config := &UAAConfig{}
	configBytes, err := interpolate(c.ConfigFile, c.VarsFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return err
	}

	logger := log.New(os.Stderr, "", 0)
	om, err := opsman.NewOpsManager(c.AuthFile, logger)
	if err != nil {
		return err
	}
	uaaClient, err := om.PASUAAClient(config.Target, config.SkipSSLValidation, config.Verbose)
	if err != nil {
		return err
	}

	for clientID, clientConfig := range config.Clients {
		client, err := uaaClient.GetClient(clientID)
		if err == nil {
			client.ClientSecret = clientConfig.Secret
			client.AuthorizedGrantTypes = clientConfig.GrantTypes
			client.Scope = clientConfig.Scope
			client.Authorities = clientConfig.Authorities
			client.AccessTokenValidity = clientConfig.AccessTokenValidity
			logger.Println(fmt.Sprintf("Updating uaa client %s", clientID))
			_, err = uaaClient.UpdateClient(*client)
			if err != nil {
				return err
			}

		} else {
			logger.Println(fmt.Sprintf("Creating uaa client %s", clientID))
			_, err = uaaClient.CreateClient(uaa.Client{
				ClientID:             clientID,
				ClientSecret:         clientConfig.Secret,
				AuthorizedGrantTypes: clientConfig.GrantTypes,
				Scope:                clientConfig.Scope,
				Authorities:          clientConfig.Authorities,
				AccessTokenValidity:  clientConfig.AccessTokenValidity,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
