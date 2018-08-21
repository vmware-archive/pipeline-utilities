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

## `om-env-file`

`om-env-file` uses flags/environment variables to build a auth.yml that can be used with `om`

### Command Usage

```
Usage:
  pipeline-utilities [OPTIONS] om-env-file [om-env-file-OPTIONS]

Help Options:
  -h, --help                       Show this help message

[om-env-file command options]
    --target=                OpsManager hostname [$OPSMAN_TARGET]
    --skip-ssl-validation    Skip SSL Validation when interacting with OpsManager [$OPSMAN_SKIP_SSL_VALIDATION]
    --username=              OpsManager username [$OPSMAN_USERNAME]
    --password=              OpsManager password [$OPSMAN_PASSWORD]
    --client-id=             OpsManager client id [$OPSMAN_CLIENTID]
    --client-secret=         OpsManager client secret [$OPSMAN_CLIENT_SECRET]
    --decryption-passphrase= OpsManager Decryption Passphrase [$OPSMAN_DECRYPTION_PASSPHRASE]
    --connect-timeout=       OpsManager Connect timeout (default: 5) [$OPSMAN_CONNECT_TIMEOUT]
    --request-timeout=       OpsManager Request timeout (default: 1800) [$OPSMAN_REQUEST_TIMEOUT]
    --output-file=           output file for yaml

    SAML:
    --idp-metadata=          OpsManager SAML IDP Metadata [$OPSMAN_SAML_IDP_METADATA]
    --bosh-idp-metadata=     OpsManager SAML BOSH IDP Metadata [$OPSMAN_SAML_BOSH_IDP_METADATA]
    --rbac-admin-group=      OpsManager RBAC admin group [$OPSMAN_RBAC_ADMIN_GROUP]
    --rbac-groups-attribute= OpsManager RBAC groups attribute [$OPSMAN_RBAC_GROUPS_ATTRIBUTE]
```

## `opsman-auth *** deprecated`

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

## `pas-uaa`

`pas-uaa` creates or updates uaa clients in target uaa based on configuration file

### Command Usage

```
Usage:
  pipeline-utilities [OPTIONS] uaa-configuration [uaa-configuration-OPTIONS]

Help Options:
  -h, --help           Show this help message

[uaa-configuration command options]
          --auth-file= path to auth file (default: auth/auth.yml)
      -c, --config=    path to config file
      -l, --vars-file= path to vars file
```

### Sample Config

```
target: uaa.((system-domain))
skip_ssl_validation: true
verbose: false
deployment: cf
uaa_admin_credential_property: .uaa.admin_client_credentials
clients:
  foo:
    secret: ((foo-secret))
    grant_types: ["client_credentials","refresh_token"]
    scope: ["doppler.firehose","cloud_controller.admin_read_only","oauth.login"]
    authorities: ["doppler.firehose","cloud_controller.admin_read_only","openid,oauth.approvals"]
    access_token_validity: 1209600
  hello:
    secret: ((hello-secret))
    grant_types: ["client_credentials","refresh_token"]
    scope: ["doppler.firehose","cloud_controller.admin_read_only","oauth.login"]
    authorities: ["doppler.firehose","cloud_controller.admin_read_only","openid,oauth.approvals"]
    access_token_validity: 1209600
```
