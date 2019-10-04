package server

// The result consists of 1) login successful 2) details when the login was not succesful 3) any error that prevented from making an authentication call
type Authenticator interface {
  Authenticate(username, password string) (bool, string, error)
}
