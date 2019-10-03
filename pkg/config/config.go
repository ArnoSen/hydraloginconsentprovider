package config

import (
  "fmt"

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/htmltemplates"
)

const (
  DefaultPort = 8888
  DefaultPrefillUser = "test"
  DefaultPrefillPassword = "1234"
  DefaultBuiltinUser = "test"
  DefaultBuiltinPassword = "test"
  DefaultHydraAdminHost = "localhost"
  DefaultHydraAdminPort = 9001
  DefaultHydraAdminBasePath = "/"
  DefaultLoginPageTemplate = htmltemplates.DefaultLoginTemplate
  DefaultConsentPageTemplate = htmltemplates.DefaultConsentTemplate
  DefaultADPort = 6686

  AUTHMODE_BUILTIN = "builtin"
  AUTHMODE_AD = "ad"
)

type Config struct {
  Port uint16
  PrefillUser string
  PrefillPassword string 
  BuiltinUser string
  BuiltinPassword string
  HydraAdminHost string
  HydraAdminPort uint16
  HydraAdminBasePath string
  SkipSSLCheck bool
  AuthFunc func(string,string) (bool, error)
  LoginPageTemplate string
  ConsentPageTemplate string
  AuthMode string
  ADDomainControllers []string
  ADDomain string
  ADPort uint16
  ADUserIdentifierProperty string
}

func DefaultConfig() *Config {

  return &Config{
    Port: DefaultPort,
    AuthMode: AUTHMODE_BUILTIN,
    PrefillUser: DefaultPrefillUser,
    PrefillPassword: DefaultPrefillPassword,
    BuiltinUser: DefaultBuiltinUser,
    BuiltinPassword: DefaultBuiltinPassword,
    HydraAdminHost: DefaultHydraAdminHost,
    HydraAdminPort: DefaultHydraAdminPort,
    HydraAdminBasePath: DefaultHydraAdminBasePath,
    AuthFunc: AuthNever,
    LoginPageTemplate: DefaultLoginPageTemplate,
    ConsentPageTemplate: DefaultConsentPageTemplate,
    ADPort: DefaultADPort,
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

func (c *Config) Dump() {

  fmt.Printf("Port: %d\n", c.Port)
  fmt.Printf("PrefillUser: %s\n", c.PrefillUser)
  fmt.Printf("PrefillPassword: %s\n", c.PrefillPassword)
  fmt.Printf("AuthMode: %s\n", c.AuthMode)
  if c.AuthMode == AUTHMODE_AD {
    fmt.Printf("ADDomainControllers: %s\n", c.ADDomainControllers)
    fmt.Printf("ADDomain: %s\n", c.ADDomain)
    fmt.Printf("ADPort: %d\n", c.ADPort)
    fmt.Printf("ADUserIdentifierProperty: %s\n", c.ADUserIdentifierProperty)
  }
  fmt.Printf("HydraAdminHost: %s\n", c.HydraAdminHost)
  fmt.Printf("HydraAdminPort: %d\n", c.HydraAdminPort)
  fmt.Printf("HydraAdminBasePath: %s\n", c.HydraAdminBasePath)

}

func ValidateAuthMode(mode string) bool {
  switch(mode) {
  case AUTHMODE_BUILTIN, AUTHMODE_AD:
    return true
  default:
    return false
  }
}
