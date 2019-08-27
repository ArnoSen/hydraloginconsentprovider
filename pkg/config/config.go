package config

import (
  "fmt"
)

const (
  DefaultPort = 8888
  DefaultTitle = "Login page"
  DefaultPrefillUser = "test"
  DefaultPrefillPassword = "1234"
  DefaultHydraAdminHost = "localhost"
  DefaultHydraAdminPort = 9001
  DefaultHydraAdminBasePath = "/"
)

type Config struct {
  Port uint16
  LoginPageHTMLTitle string
  PrefillUser string
  PrefillPassword string 
  HydraAdminHost string
  HydraAdminPort uint16
  HydraAdminBasePath string
  SkipAuthLoginResponseSSLCheck bool
}

func DefaultConfig() *Config {

  return &Config{
    Port: DefaultPort,
    LoginPageHTMLTitle: DefaultTitle,
    PrefillUser: DefaultPrefillUser,
    PrefillPassword: DefaultPrefillPassword,
    HydraAdminHost: DefaultHydraAdminHost,
    HydraAdminPort: DefaultHydraAdminPort,
    HydraAdminBasePath: DefaultHydraAdminBasePath,
  }
}

func (c *Config) SetSkipAuthLoginResponseSSLCheck() {
  c.SkipAuthLoginResponseSSLCheck = true
}

func (c *Config) GetHydraAdminHostname() string {
  return fmt.Sprintf("%s:%d", c.HydraAdminHost, c.HydraAdminPort)
}
