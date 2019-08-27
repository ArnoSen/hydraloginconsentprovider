package main

import (
  "os"
  "strconv"
  "log"

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/server"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/config"
)

const (
  ENV_LOGINPROVIDER_PORT = "LOGINPROVIDER_PORT"
)

func main() {

  cfg := config.DefaultConfig()
  cfg.SetSkipAuthLoginResponseSSLCheck()

  if os.Getenv(ENV_LOGINPROVIDER_PORT) != "" {
     u, err := strconv.ParseUint(os.Getenv("LOGINPROVIDER_PORT"), 10, 16)
     if err != nil {
       log.Fatalf("Invalid value for '%s': %s", ENV_LOGINPROVIDER_PORT, os.Getenv(ENV_LOGINPROVIDER_PORT) )
     }
     cfg.Port = uint16(u)

  }

  s := server.New(cfg)
  s.Start()
}
