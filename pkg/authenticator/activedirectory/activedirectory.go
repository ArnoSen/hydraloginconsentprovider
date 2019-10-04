package activedirectory

import (
  "strings"
  "fmt"
  "crypto/tls"

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

func (a *ActiveDirectory) Authenticate(username, password string) (bool, string, error) {

  if len(a.DomainControllers) == 0 {
    return false, "", fmt.Errorf("No domain controllers configured")
  } 

  l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", a.DomainControllers[0], a.Port), &tls.Config{ ServerName: a.DomainControllers[0], InsecureSkipVerify: true } )

  if err != nil {
    return false, "", err
  }
  defer l.Close()

  ldapUsername := fmt.Sprintf("%s@%s", username, a.Domain)

  err = l.Bind(ldapUsername, password)
  if err != nil {
    var errorText string

    if IsErrorType(err, WRONGCREDENTIALSERROR) {
      errorText = "Invalid username/password" 
    }
    if IsErrorType(err, PASSWORDEXPIRED) {
      errorText = "Password has expired" 
    }
    if IsErrorType(err, PASSWORDMUSTBERESET) {
      errorText = "Password must be set at next logon" 
    }
    if errorText != "" {
      return false, errorText, nil
    }

    return false, "", err

  }
  return true, "", nil
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
