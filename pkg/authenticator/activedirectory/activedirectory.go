package activedirectory

import (
  "strings"
  "fmt"

  "gopkg.in/ldap.v2"
)

type ActiveDirectory struct {
  DomainControllers []string 
  Domain string
  Port uint16
  UserIdentifierProperty string
}

func NewActiveDirectoryAuthenticator(dcs []string, domain string, port uint16, uip string) *ActiveDirectory {
  return &ActiveDirectory{
    DomainControllers: dcs,
    Domain: domain,
    Port: port,
    UserIdentifierProperty: uip,
  }
}

func (a *ActiveDirectory) Authenticate(username, password string) (bool, error) {

  if len(a.DomainControllers) == 0 {
    return false, fmt.Errorf("No domain controllers configured")
  } 

  l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", a.DomainControllers[0], a.Port))
  if err != nil {
    return false, err
  }
  defer l.Close()

  err = l.Bind(fmt.Sprintf("%s=%s,%s", a.UserIdentifierProperty, domainToDC(a.Domain)), password)
  if err != nil {
    return false, err
  }
  return true, nil
}

//Converts a domain e.g. mydomain.com to 'DC=mydomain,DC=com'
func domainToDC(d string) string {

  dots := strings.Split(d, ".")
 
  var result []string

  for _, dotPart := range dots {
    result=append(result, fmt.Sprintf("dc=%s", dotPart))
  }

  return strings.Join(result, ",")
}
