package commands

type manager struct {
	EnvironmentToYAML EnvironmentToYAML `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
	CreateOpsmanAuth  CreateOpsmanAuth  `command:"opsman-auth" description:"create yaml file used for opsman authorization"`
	UAAConfiguration  UAAConfiguration  `command:"uaa-configuration" description:"creates/updates uaa client(s) within target uaa"`
	Version           Version           `command:"version" description:"prints version"`
}

var Manager manager
