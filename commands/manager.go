package commands

type manager struct {
	EnvironmentToYAML       EnvironmentToYAML       `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
	CreateAuthFile          CreateAuthFile          `command:"auth-file" description:"create yaml file used for opsman auth configuration. http://docs-platform-automation.cfapps.io/pcf-automation/task-reference.html#auth"`
	CreateEnvironmentFile   CreateEnvironmentFile   `command:"env-file" description:"create om env file used for om --env flag. http://docs-platform-automation.cfapps.io/pcf-automation/task-reference.html#env"`
	CreateOpsmanAuth        CreateOpsmanAuth        `command:"opsman-auth" description:"DEPRECATED - create yaml file used for opsman authorization"`
	CreateOMEnvironmentFile CreateOMEnvironmentFile `command:"om-env-file" description:"DEPRECATED - create om env file used for om --env flag"`
	UAAConfiguration        UAAConfiguration        `command:"uaa-configuration" description:"creates/updates uaa client(s) within target uaa"`
	Version                 Version                 `command:"version" description:"prints version"`
}

var Manager manager
