package server

type Authenticator interface {
  Authenticate(username, password string) (bool, error)
}
