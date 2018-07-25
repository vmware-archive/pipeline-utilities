# pipeline-utilities
Utilities that make working with concourse easier

## Getting Started

## Maintainer

* [Caleb Washburn](https://github.com/calebwashburn)

## Support

pipeline-utilities is a community supported cli.  Opening issues for questions, feature requests and/or bugs is the best path to getting "support".  We strive to be active in keeping this tool working and meeting your needs in a timely fashion.

## Build from the source

`pipeline-utilities` is written in [Go](https://golang.org/).
To build the binary yourself, follow these steps:

* Install `Go`.
* Install [dep](https://github.com/golang/dep), a dependency management tool for Go.
* Clone the repo:
  - `mkdir -p $(go env GOPATH)/src/github.com/pivotalservices`
  - `cd $(go env GOPATH)/src/github.com/pivotalservices`
  - `git clone git@github.com:pivotalservices/pipeline-utilities.git`
* Install dependencies:
  - `cd pipeline-utilities`
  - `dep ensure`
  - `go build -o pipeline-utilities cmd/pipeline-utilities/main.go`

To cross compile, set the `$GOOS` and `$GOARCH` environment variables.
For example: `GOOS=linux GOARCH=amd64 go build`.


## Contributing

PRs are always welcome or open issues if you are experiencing an issue and will do my best to address issues in timely fashion.

## Documentation

## `env-to-yaml`

`env-to-yaml` takes any environment variable with the prefix specified in `--environment-prefix` and will generate a yaml vars file with those values.  This is useful to taking values managed in a concourse credential management solution like credhub and converting into a vars file that can be passed to `om` or `bosh` interpolate commands.

### Command Usage

```
Usage:
  pipeline-utilities [OPTIONS] env-to-yaml [env-to-yaml-OPTIONS]

Help Options:
  -h, --help                    Show this help message

[env-to-yaml command options]
  --environment-prefix= prefix for environment variables (default: YAML)
  --output-file=        output file for yaml
```


## `opsman-auth`

`opsman-auth` uses flags/environment variables to build a auth.yml that can be used with PCF Automation tooling

### Command Usage

```
Usage:
  pipeline-utilities [OPTIONS] opsman-auth [opsman-auth-OPTIONS]

Help Options:
  -h, --help                       Show this help message

[opsman-auth command options]
  --url=                   OpsManager URL [$OPSMAN_URL]
  --skip-ssl-validation    Skip SSL Validation when interacting with OpsManager [$OPSMAN_SKIP_SSL_VALIDATION]
  --username=              OpsManager username [$OPSMAN_USERNAME]
  --password=              OpsManager password [$OPSMAN_PASSWORD]
  --client-id=             OpsManager client id [$OPSMAN_CLIENTID]
  --client-secret=         OpsManager client secret [$OPSMAN_CLIENT_SECRET]
  --decryption-passphrase= OpsManager Decryption Passphrase [$OPSMAN_DECRYPTION_PASSPHRASE]
  --output-file=           output file for yaml
```
