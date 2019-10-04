package config

import (
  "fmt"

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/htmltemplates"
)

const (
  DefaultPort = 8888
  DefaultPrefillBuiltinCredentials = true
  DefaultBuiltinUser = "test"
  DefaultBuiltinPassword = "test"
  DefaultHydraAdminHost = "localhost"
  DefaultHydraAdminPort = 9001
  DefaultHydraAdminBasePath = "/"
  DefaultLoginPageTemplate = htmltemplates.DefaultLoginTemplate
  DefaultConsentPageTemplate = htmltemplates.DefaultConsentTemplate
  DefaultADPort = 636
  DefaultPrivateKeyLocation = "etc/loginprovider.key"
  DefaultCertLocation = "etc/loginprovider.crt"

  AUTHMODE_BUILTIN = "builtin"
  AUTHMODE_AD = "ad"
)

type Config struct {
  PrivateKeyLocation string
  CertLocation string
  Port uint16
  PrefillBuiltinCredentials bool
  BuiltinUser string
  BuiltinPassword string
  HydraAdminHost string
  HydraAdminPort uint16
  HydraAdminBasePath string
  SkipSSLCheck bool
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
    PrivateKeyLocation: DefaultPrivateKeyLocation,
    CertLocation: DefaultCertLocation,
    Port: DefaultPort,
    AuthMode: AUTHMODE_BUILTIN,
    PrefillBuiltinCredentials: DefaultPrefillBuiltinCredentials,
    BuiltinUser: DefaultBuiltinUser,
    BuiltinPassword: DefaultBuiltinPassword,
    HydraAdminHost: DefaultHydraAdminHost,
    HydraAdminPort: DefaultHydraAdminPort,
    HydraAdminBasePath: DefaultHydraAdminBasePath,
    LoginPageTemplate: DefaultLoginPageTemplate,
    ConsentPageTemplate: DefaultConsentPageTemplate,
    ADPort: DefaultADPort,
  }
}

func (c *Config) SetPrivateKeyLocation(l string) {
  c.PrivateKeyLocation = l
}

func (c *Config) SetCertLocation(l string) {
  c.CertLocation = l
}

func (c *Config) SetSkipSSLCheck() {
  c.SkipSSLCheck = true
}

func (c *Config) GetHydraAdminHostname() string {
  return fmt.Sprintf("%s:%d", c.HydraAdminHost, c.HydraAdminPort)
}

func (c *Config) Dump() {

  fmt.Printf("PrivateKeyLocation: %s\n", c.PrivateKeyLocation)
  fmt.Printf("CertificateLocation: %s\n", c.CertLocation)
  fmt.Printf("Port: %d\n", c.Port)
  fmt.Printf("AuthMode: %s\n", c.AuthMode)
  if c.AuthMode == AUTHMODE_AD {
    fmt.Printf("ADDomainControllers: %s\n", c.ADDomainControllers)
    fmt.Printf("ADDomain: %s\n", c.ADDomain)
    fmt.Printf("ADPort: %d\n", c.ADPort)
    fmt.Printf("ADUserIdentifierProperty: %s\n", c.ADUserIdentifierProperty)
  } else {
    fmt.Printf("PrefillBuiltinCredentials: %t\n", c.PrefillBuiltinCredentials)
  }
  fmt.Printf("HydraAdminHost: %s\n", c.HydraAdminHost)
  fmt.Printf("HydraAdminPort: %d\n", c.HydraAdminPort)
  fmt.Printf("HydraAdminBasePath: %s\n", c.HydraAdminBasePath)
  fmt.Printf("SkipSSLCheck: %t\n", c.SkipSSLCheck)
}

func ValidateAuthMode(mode string) bool {
  switch(mode) {
  case AUTHMODE_BUILTIN, AUTHMODE_AD:
    return true
  default:
    return false
  }
}
