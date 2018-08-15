package opsman

import (
	"encoding/json"
	"errors"
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

func NewOpsManager(authFile string, logger *log.Logger) (*OpsManager, error) {

	authConfig := &AuthConfig{}
	authBytes, err := ioutil.ReadFile(authFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(authBytes, authConfig)
	if err != nil {
		return nil, err
	}
	requestTimeout := time.Duration(1800) * time.Second
	connectTimeout := time.Duration(5) * time.Second

	var authedClient httpClient
	authedClient, err = network.NewOAuthClient(authConfig.OpsmanURL, authConfig.Credentials.UserName, authConfig.Credentials.Password, authConfig.Credentials.ClientID, authConfig.Credentials.ClientSecret, authConfig.SkipSSLValidation, false, requestTimeout, connectTimeout)
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

func (o *OpsManager) PASUAAClient(target string, skipSSLValidation, verbose bool) (*uaa.API, error) {
	deployedProducts, err := o.om.ListDeployedProducts()
	if err != nil {
		return nil, err
	}

	var GUID string
	for _, deployedProduct := range deployedProducts {
		if strings.EqualFold(deployedProduct.Type, "cf") {
			GUID = deployedProduct.GUID
			continue
		}
	}

	if GUID == "" {
		return nil, errors.New("Deployed Product CF not found")
	}

	output, err := o.om.Curl(api.RequestServiceCurlInput{
		Path: fmt.Sprintf("/api/v0/deployed/products/%s/credentials/.uaa.admin_client_credentials", GUID),
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
