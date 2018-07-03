package commands

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type manager struct {
	EnvironmentToYAML EnvironmentToYAML `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
	CreateOpsmanAuth  CreateOpsmanAuth  `command:"opsman-auth" description:"create yaml file used for opsman authorization"`
}

var Manager manager

func writeYamlFile(targetFile string, dataType interface{}) error {

	data, err := yaml.Marshal(dataType)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetFile, data, 0755)

}
