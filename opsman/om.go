package opsman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	uaa "github.com/cloudfoundry-community/go-uaa"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/network"
	yaml "gopkg.in/yaml.v2"
)

type OpsManager struct {
	om api.Api
}

type uaaCreds struct {
	Credential struct {
		Value struct {
			Identity string `yaml:"identity"`
			Password string `yaml:"password"`
		} `json:"value"`
	} `json:"credential"`
}

func NewOpsManager(envFile string, logger *log.Logger) (*OpsManager, error) {

	envConfig := &ENVConfig{}
	envBytes, err := ioutil.ReadFile(envFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(envBytes, envConfig)
	if err != nil {
		return nil, err
	}
	requestTimeout := time.Duration(1800) * time.Second
	connectTimeout := time.Duration(5) * time.Second

	var authedClient httpClient
	authedClient, err = network.NewOAuthClient(envConfig.Target, envConfig.UserName, envConfig.Password, envConfig.ClientID, envConfig.ClientSecret, envConfig.SkipSSLValidation, false, requestTimeout, connectTimeout)
	if err != nil {
		return nil, err
	}
	om := api.New(api.ApiInput{
		Client:                 authedClient,
		UnauthedClient:         nil,
		ProgressClient:         nil,
		UnauthedProgressClient: nil,
		Logger:                 logger,
	})

	return &OpsManager{
		om: om,
	}, nil
}

func (o *OpsManager) UAAClient(deployment, credentialPath, target string, skipSSLValidation, verbose bool) (*uaa.API, error) {
	deployedProducts, err := o.om.ListDeployedProducts()
	if err != nil {
		return nil, err
	}

	var GUID string
	for _, deployedProduct := range deployedProducts {
		if strings.EqualFold(deployedProduct.Type, deployment) {
			GUID = deployedProduct.GUID
			continue
		}
	}

	if GUID == "" {
		return nil, fmt.Errorf("Deployed Product %s not found", deployment)
	}

	output, err := o.om.Curl(api.RequestServiceCurlInput{
		Path: fmt.Sprintf("/api/v0/deployed/products/%s/credentials/%s", GUID, credentialPath),
	})

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read api response body: %s", err)
	}
	creds := &uaaCreds{}
	err = json.Unmarshal(body, creds)
	if err != nil {
		return nil, err
	}

	uaaClient, err := uaa.NewWithClientCredentials(target, "", creds.Credential.Value.Identity, creds.Credential.Value.Password, uaa.OpaqueToken, skipSSLValidation)
	uaaClient.Verbose = verbose
	if err != nil {
		return nil, err
	}
	return uaaClient, nil
}
