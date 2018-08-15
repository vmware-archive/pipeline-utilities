package commands

import (
	"fmt"
	"io/ioutil"

	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"
	"gopkg.in/yaml.v2"
)

type manager struct {
	EnvironmentToYAML EnvironmentToYAML `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
	CreateOpsmanAuth  CreateOpsmanAuth  `command:"opsman-auth" description:"create yaml file used for opsman authorization"`
	PASUAA            PASUAA            `command:"pas-uaa" description:"creates/updates uaa client(s) within PAS uaa"`
	Version           Version           `command:"version" description:"prints version"`
}

var Manager manager

func writeYamlFile(targetFile string, dataType interface{}) error {

	data, err := yaml.Marshal(dataType)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetFile, data, 0755)

}

func interpolate(templateFile string, varsFiles []string) ([]byte, error) {
	contents, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	tpl := boshtpl.NewTemplate(contents)
	vars := []boshtpl.Variables{}

	for i := len(varsFiles) - 1; i >= 0; i -= 1 {
		path := varsFiles[i]
		payload, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("could not read template variables file (%s): %s", path, err.Error())
		}
		var staticVars boshtpl.StaticVariables
		err = yaml.Unmarshal(payload, &staticVars)
		if err != nil {
			return nil, fmt.Errorf("could not unmarhsal template variables file (%s): %s", path, err.Error())
		}
		vars = append(vars, staticVars)
	}
	evalOpts := boshtpl.EvaluateOpts{
		UnescapedMultiline: true,
		ExpectAllKeys:      true,
	}

	bytes, err := tpl.Evaluate(boshtpl.NewMultiVars(vars), nil, evalOpts)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
