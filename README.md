Hydray login/consent provider in GO
See: https://github.com/ory/hydra-login-consent-node

This is the initial version and has lots of raw edges.

Usage:

package main

import (

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/server"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/config"
)

func main() {

  cfg := &config.Config{
    Port: 5000,
    LoginPageHTMLTitle: "LoginPageNae",
    PrefillUser: "default",
    PrefillPassword: "1234",
    HydraAdminHost: "hydra-admin-hostname",
    HydraAdminPort: 9001,
    HydraAdminBasePath: "/",
    SkipAuthLoginResponseSSLCheck: true, // the loginconsent provider will accept any certificate from they Hydra admin host
  }

  s := server.New(cfg)
  s.Start()
}
