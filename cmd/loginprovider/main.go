package main

import (
  "path/filepath"
  "os"
  "fmt"

  "github.com/ArnoSen/hydraloginconsentprovider/pkg/server"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/config"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

const (
  VERSION="0.0.2"
  VERSIONTEMPLATE=`{{printf "Version: %s\n" .Version}}`

  // command defaults
  SYSTEMCONFIGDIRECTORY="/etc/loginprovider"
  LOCALCONFIGDIRECTORY="./etc"
  DEFAULTCONFIGFILENAMEPREFIX="loginprovider"
  ENVPREFIX="LP"
)

var myviper *viper.Viper

var RootCmd = &cobra.Command{
  Use:   "loginprovider",
  Short: "Hydra login and consent provider",
  ValidArgs: []string{"show-config", "start"},
  Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage()
  },
  Version: VERSION,
}

var ShowConfigCmd = &cobra.Command{
  Use:   "show-config",
  Short: "Show config variables",
  Run: func(cmd *cobra.Command, args []string) {
    cfg, err := parseConfig(myviper)
    if err != nil {
      fmt.Printf("Error in config: %s\n", err)
      os.Exit(1)
    }
    cfg.Dump()
  },
}

var StartCmd = &cobra.Command{
  Use:   "start",
  Short: "Start vROPs pollerd",
  Run: func(cmd *cobra.Command, args []string) {

    cfg, err := parseConfig(myviper)
    if err != nil {
      fmt.Printf("Error in config: %s\n", err)
      os.Exit(1)
    }

    server.New(cfg).Start()
  },
}

func parseConfig(v *viper.Viper) (*config.Config, error)  {

  if validateErr := validateConfig(v); validateErr != nil {
    return nil, validateErr
  }

  c := config.DefaultConfig()

  if i := v.GetInt("Port"); i != 0 { 
    c.Port = uint16(i)
  }
  if i := v.GetBool("PrefillBuiltinCredentials"); i {
    c.PrefillBuiltinCredentials = i
  }
  if i := v.GetString("HydraAdminHost"); i != "" {
    c.HydraAdminHost = i
  }
  if i := v.GetString("PrivateKeyLocation"); i != "" {
    c.PrivateKeyLocation = i
  }
  if i := v.GetString("CertLocation"); i != "" {
    c.CertLocation = i
  }
  if i := v.GetInt("HydraAdminPort"); i != 0 {
    c.HydraAdminPort = uint16(i)
  }
  if v.GetBool("SkipSSLCheck") {
    c.SkipSSLCheck = true
  }
  if i := v.GetString("AuthMode"); i != "" {
    c.AuthMode = i
  }
  if i := v.GetStringSlice("ADDomainControllers"); len(i) > 0 {
    c.ADDomainControllers = i
  }
  if i := v.GetString("ADDomain"); i != "" {
    c.ADDomain = i
  }
  if i := v.GetInt("ADDomainPort"); i != 0 {
    c.ADPort = uint16(i)
  }
  if i := v.GetString("ADUserIdentifierProperty"); i != "" {
    c.ADUserIdentifierProperty = i
  }

  return c, nil
}

func validateConfig(v *viper.Viper) error {

  if mode := v.GetString("AuthMode"); mode != "" {
    if !config.ValidateAuthMode(mode) {
      return fmt.Errorf("Unsupported auth mode '%s'", mode)
    }
  }
  return nil
}

func main() {
  wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    fmt.Printf("Cannot get working directory: %s", err)
    os.Exit(1)
  }

  RootCmd.InitDefaultVersionFlag()
  RootCmd.SetVersionTemplate(VERSIONTEMPLATE)
  RootCmd.AddCommand(ShowConfigCmd)
  RootCmd.AddCommand(StartCmd)

  myviper = viper.GetViper()
  myviper.SetEnvPrefix(ENVPREFIX)
  myviper.AutomaticEnv()

  myviper.AddConfigPath(SYSTEMCONFIGDIRECTORY)
  myviper.AddConfigPath(fmt.Sprintf("%s/%s", wd, LOCALCONFIGDIRECTORY ))
  myviper.SetConfigName(DEFAULTCONFIGFILENAMEPREFIX)

  // Read any config from the search path
  err = myviper.ReadInConfig()
  if err != nil {
    if _, castConfigFileNotFoundError := err.(viper.ConfigFileNotFoundError); !castConfigFileNotFoundError {
      fmt.Printf("Unable to read configfile: %s\n", err)
      os.Exit(1)
    }
  }

  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
