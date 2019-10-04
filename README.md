# Hydra login/consent provider in GO

NB: this application is new and has some raw edges

This is an Hydra Login Consent node written in GO (https://github.com/ory/hydra-login-consent-node).
It uses the Hydra SDK for now.

Supported authentication sources are:
- builtin authentication (intended for quick test cycles)
- active directory authentication

# Docker

Docker: arnosen/loginprovider-0.0.2

# Running the application

To run the prebuild application:
```
$ bin/loginprovider
```

To compile yourself:
```
$ cd cmd/loginprovider
$ go build
```

# Configuration

For configuration, both config files can be used or you can use environment variables.
Config files are looked for either in `/etc` or `./etc` and should be named `loginprovider.yml`

A sample configuration file can be found `cmd/loginprovider/etc`.

Sensible defaults have been pre-defined.

To verfify what paramters `loginprovider` has read, run:
```
$ bin/loginprovider show-config
```

The following environment variables are supported:
- LP_PRIVATEKEYLOCATION: filepath for the private key
- LP_CERTLOCATION: filepath for the cert
- LP_PORT: listing port for the service
- LP_AUTHMODE: ad|builtin
- LP_BUILTINUSER: set the builtin username
- LP_BUILTINPASSWORD: set the builtin password
- LP_PREFILLBUILTINCREDENTIALS: if true, the username and password are prefilled in the login form
- LP_HYDRAADMINHOST: hydra admin host (so not the public host)
- LP_HYDRAADMINPORT: hydra admin port
- LP_HYDRAADMINBASEPATH: base path
- LP_SKIPSSLCHECK: if true, the Hydra certificate is not checked for validity
- LP_ADDOMAINCONTROLLERS: list of domain controllers seperated by spaces
- LP_ADDOMAIN: name of the domain, e.g. mydomain.company.com
- LP_ADPORT: active directory port to use. Only LDAPs ports are supported
