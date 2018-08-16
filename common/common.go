package common

import (
	"fmt"
	"io/ioutil"

	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"
	yaml "gopkg.in/yaml.v2"
)

func WriteYamlFile(targetFile string, dataType interface{}) error {

	data, err := yaml.Marshal(dataType)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetFile, data, 0755)

}

func Interpolate(templateFile string, varsFiles []string) ([]byte, error) {
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
