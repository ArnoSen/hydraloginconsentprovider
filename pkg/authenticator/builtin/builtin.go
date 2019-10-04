package builtin

type BuiltIn struct {
  username string
  password string
}

func NewBuiltInAuthorizer(username, password string) *BuiltIn {
  return &BuiltIn{
    username: username,
    password: password,
  }
}

func (b *BuiltIn) Authenticate(username, password string) (bool, string, error) {
  if username==b.username && password==b.password {
    return true, "", nil
  } else {
    return false, "Incorrect username and/or password", nil
  }
}
