package main

import (

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/server"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/config"
)

const (
  PREFILL_USER="default"
  PREFILL_PASSWORD="1234"
)

func authFunc(username, password string) (bool, error) {
  return username==PREFILL_USER && password==PREFILL_PASSWORD, nil
}

func main() {

  cfg := &config.Config{
    Port: 5000,
    PrefillUser: "default",
    PrefillPassword: "1234",
    HydraAdminHost: "hydra-admin-hostname",
    HydraAdminPort: 9001,
    HydraAdminBasePath: "/",
    SkipSSLCheck: true, // the loginconsent provider will accept any certificate from they Hydra admin host
    AuthFunc: authFunc,
  }

  s := server.New(cfg)
  s.Start()
}

