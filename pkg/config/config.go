package config

import (
  "fmt"
)

const (
  DefaultPort = 8888
  DefaultPrefillUser = "test"
  DefaultPrefillPassword = "1234"
  DefaultHydraAdminHost = "localhost"
  DefaultHydraAdminPort = 9001
  DefaultHydraAdminBasePath = "/"
)

type Config struct {
  Port uint16
  PrefillUser string
  PrefillPassword string 
  HydraAdminHost string
  HydraAdminPort uint16
  HydraAdminBasePath string
  SkipSSLCheck bool
  AuthFunc func(string,string) (bool, error)
}

func DefaultConfig() *Config {

  return &Config{
    Port: DefaultPort,
    PrefillUser: DefaultPrefillUser,
    PrefillPassword: DefaultPrefillPassword,
    HydraAdminHost: DefaultHydraAdminHost,
    HydraAdminPort: DefaultHydraAdminPort,
    HydraAdminBasePath: DefaultHydraAdminBasePath,
    AuthFunc: AuthNever,
  }
}

func (c *Config) SetSkipSSLCheck() {
  c.SkipSSLCheck = true
}

func (c *Config) GetHydraAdminHostname() string {
  return fmt.Sprintf("%s:%d", c.HydraAdminHost, c.HydraAdminPort)
}

func AuthNever(username, password string) (bool, error) {
  return false, nil
}
