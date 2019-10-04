package loginconsentclient 

import (
  "fmt"
  "context"
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "crypto/tls"
  "reflect"
  "net/http/httputil"
)

type Client struct {
  Host string
  Port uint16
  Basepath string
  skipSSLCheck bool
}

func NewClient(host, basepath string, port uint16) *Client {
  return &Client {
    Host: host,
    Basepath: basepath,
    Port: port,
  }
}

func (c *Client) SkipSSLCheck() {
  c.skipSSLCheck = true
}

func (c *Client) GetURL(path string) string {
  return fmt.Sprintf("https://%s:%d/%s/%s", c.Host, c.Port, c.Basepath, path)
}

func (c *Client) httpRequest(ctx context.Context, method, path string, payLoad interface{}, responseStruct interface{}) (int, []byte, error) {

  //Executes an http request. Returns the statuscode, the body and an error (if applicable)
  urlString := c.GetURL(path)

  //fmt.Printf("URL: %s\n", urlString)

  var br *bytes.Reader

  if isInterfaceEmpty(payLoad) || method == "GET"{
    br = bytes.NewReader([]byte{})
  } else {
    payloadBytes, err := json.Marshal(payLoad)
    if err != nil {
      return 0, []byte{}, err
    }
 
    br = bytes.NewReader(payloadBytes)
  }

  req, requestErr := http.NewRequest(method, urlString, br)

  if (requestErr != nil) {
    return 0, []byte{}, requestErr
  }

  req = req.WithContext(ctx)
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "*/*")

  tp := &http.Transport{}
  if c.skipSSLCheck {
    tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
  }

  httpClient := &http.Client{ Transport: tp }

  reqText, _ := httputil.DumpRequest(req, false)
  fmt.Printf("Req: %s\n", reqText)

  resp, err := httpClient.Do(req)

  if err != nil {
    return 0, []byte{}, fmt.Errorf("Invalid request: %s", err)
  }
  defer resp.Body.Close()

  bodyBytes, bodyReadError := ioutil.ReadAll(resp.Body)

  if bodyReadError != nil {
    return resp.StatusCode, []byte{}, bodyReadError
  }

  if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

    respText, _ := httputil.DumpResponse(resp, false)
    fmt.Printf("Resp: %s\n", respText)
    

    if len(bodyBytes) == 0 {
      return resp.StatusCode, bodyBytes, nil
    }

    fmt.Printf("Body: %s\n", bodyBytes)

    if responseStruct == nil {
      return resp.StatusCode, bodyBytes, nil
    }

    if unmarshalError := json.Unmarshal(bodyBytes, &responseStruct); unmarshalError != nil {
      return resp.StatusCode, bodyBytes, unmarshalError
    }

    return resp.StatusCode, bodyBytes, nil
  }

  if len(bodyBytes) > 0 {
    /*
    errObject := Error{}

    if unmarshalError := json.Unmarshal(bodyBytes, &errObject); unmarshalError != nil {
      return resp.StatusCode, bodyBytes, unmarshalError
    }
    */
    return resp.StatusCode, bodyBytes, fmt.Errorf("Call return statuscode %d", resp.StatusCode)
  }

  return resp.StatusCode, bodyBytes, fmt.Errorf("Status code %d returned", resp.StatusCode)
}

func isInterfaceEmpty(c interface{}) bool {
  return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}
