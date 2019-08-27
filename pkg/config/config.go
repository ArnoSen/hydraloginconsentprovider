package config

import (
  "fmt"

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/htmltemplates"
)

const (
  DefaultPort = 8888
  DefaultPrefillUser = "test"
  DefaultPrefillPassword = "1234"
  DefaultHydraAdminHost = "localhost"
  DefaultHydraAdminPort = 9001
  DefaultHydraAdminBasePath = "/"
  DefaultLoginPageTemplate = htmltemplates.DefaultLoginTemplate
  DefaultConsentPageTemplate = htmltemplates.DefaultConsentTemplate
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
  LoginPageTemplate string
  ConsentPageTemplate string
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
    LoginPageTemplate: DefaultLoginPageTemplate,
    ConsentPageTemplate: DefaultConsentPageTemplate,
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
