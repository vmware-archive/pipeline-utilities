package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type EnvironmentToYAML struct {
	EnvPrefix  string `long:"environment-prefix" description:"prefix for environment variables" default:"YAML"`
	OutputFile string `long:"output-file" description:"output file for yaml" required:"true"`
}

//Execute - generates structs
func (c *EnvironmentToYAML) Execute([]string) error {
	dataType := make(map[string]interface{})
	thePrefix := fmt.Sprintf("%s_", c.EnvPrefix)
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, thePrefix) {
			keyValue := strings.SplitN(env, "=", 2)
			key := strings.ToLower(strings.Replace(keyValue[0], thePrefix, "", -1))
			dataType[key] = keyValue[1]
		}
	}

	return writeYamlFile(c.OutputFile, dataType)
}

func writeYamlFile(targetFile string, dataType map[string]interface{}) error {
	if len(dataType) > 0 {
		data, err := yaml.Marshal(dataType)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(targetFile, data, 0755)
	}
	return ioutil.WriteFile(targetFile, []byte(""), 0755)
}
