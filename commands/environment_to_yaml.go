package commands

import (
	"fmt"
	"os"
	"strings"
)

type EnvironmentToYAML struct {
	EnvPrefix  string `long:"environment-prefix" description:"prefix for environment variables" default:"YAML"`
	OutputFile string `long:"output-file" description:"output file for yaml" required:"true"`
}

//Execute - creates yaml file based on enviroment variable prefix
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
	if len(dataType) > 0 {
		return writeYamlFile(c.OutputFile, dataType)
	}
	return fmt.Errorf("No environment variables with prefix: %s", c.EnvPrefix)
}
