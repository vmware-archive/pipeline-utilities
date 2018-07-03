package commands

type manager struct {
	EnvironmentToYAML EnvironmentToYAML `command:"env-to-yaml" description:"create yaml variable file from environment variables"`
}

var Manager manager
