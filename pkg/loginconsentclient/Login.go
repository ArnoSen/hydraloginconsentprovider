package loginconsentclient

import (
  "context"
  "fmt"
)

func (c *Client) GetLoginRequest(ctx context.Context, challenge string) (*LoginRequest , error) {
  //GET /oauth2/auth/requests/login?login_challenge=aabe78bd82f34ecba750ab52250a67eb

  path := fmt.Sprintf("/oauth2/auth/requests/login?login_challenge=%s", challenge)

  resp := &LoginRequest{}

  _, _, err := c.httpRequest(ctx, "GET", path, nil, resp)  

  return resp, err 
}

func (c *Client) AcceptLoginRequest(ctx context.Context, r *HandledLoginRequest, challenge string) (*RequestHandlerResponse, error) {

  path := fmt.Sprintf("/oauth2/auth/requests/login/accept?login_challenge=%s", challenge)

  resp := &RequestHandlerResponse{}

  _, _, err := c.httpRequest(ctx, "PUT", path, r, resp)  

  return resp, err 
}
