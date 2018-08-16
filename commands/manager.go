package commands

type manager struct {
	EnvironmentToYAML       EnvironmentToYAML       `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
	CreateOpsmanAuth        CreateOpsmanAuth        `command:"opsman-auth" description:"DEPRECATED - create yaml file used for opsman authorization"`
	CreateOMEnvironmentFile CreateOMEnvironmentFile `command:"om-env-file" description:"create om env file used for om --env flag"`
	UAAConfiguration        UAAConfiguration        `command:"uaa-configuration" description:"creates/updates uaa client(s) within target uaa"`
	Version                 Version                 `command:"version" description:"prints version"`
}

var Manager manager
