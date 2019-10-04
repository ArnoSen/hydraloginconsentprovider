package activedirectory

import (
  "strings"
  "fmt"

  "gopkg.in/ldap.v2"
)

const (
  WRONGCREDENTIALSERROR = iota // this includes all kind of auth failures
  PASSWORDEXPIRED              // this is a particular fail
  UNWILLINGTOPERFORM
  PASSWORDPOLICYVIOLATION
  PASSWORDMUSTBERESET
  USERNOTFOUND
)

const (

  authWrongPasswordErrorString = "LDAP Result Code 49 \"Invalid Credentials\": 80090308: LdapErr: DSID-0C09042F, comment: AcceptSecurityContext error, data 52e, v2580"

  authPasswordExpiredErrorString = "LDAP Result Code 49 \"Invalid Credentials\": 80090308: LdapErr: DSID-0C09042F, comment: AcceptSecurityContext error, data 532, v2580"

  authUserMustResetErrorString = "LDAP Result Code 49 \"Invalid Credentials\": 80090308: LdapErr: DSID-0C09042F, comment: AcceptSecurityContext error, data 773, v2580"

  unwillToPerformErrorString = "LDAP Result Code 53 \"Unwilling To Perform\": 0000052D: SvcErr: DSID-031A12D2, problem 5003 (WILL_NOT_PERFORM), data 0"

  passwordPolicyViolationErrorString = "LDAP Result Code 19 \"Constraint Violation\": 0000052D: AtrErr: DSID-03191083, #1:\n\t0: 0000052D: DSID-03191083, problem 1005 (CONSTRAINT_ATT_TYPE), data 0, Att 9005a (unicodePwd)"

)

func IsErrorType(e error, errorType int) bool {

  if errorType == USERNOTFOUND {
    return e.Error() == fmt.Sprintf("%d", USERNOTFOUND)
  }

  _, ok := e.(*ldap.Error)
  if !ok {
    return false
  }

  switch errorType {
  case WRONGCREDENTIALSERROR:
    return ldap.IsErrorWithCode(e, 49)
  case PASSWORDEXPIRED:
    return strings.HasPrefix(e.Error(), authPasswordExpiredErrorString)
  case UNWILLINGTOPERFORM:
    return strings.HasPrefix(e.Error(), unwillToPerformErrorString)
  case PASSWORDPOLICYVIOLATION:
    return strings.HasPrefix(e.Error(), passwordPolicyViolationErrorString)
  case PASSWORDMUSTBERESET:
    return strings.HasPrefix(e.Error(), authUserMustResetErrorString)
  default:
    return false
  }
}
