package server

import (
  "net/http"
  "fmt"
  "log"
  "text/template"
  "crypto/tls"
  "time"
  "strings"

  "github.com/sirupsen/logrus"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/htmltemplates"
  "github.com/ArnoSen/hydraloginconsentprovider/pkg/config"
  hydraclient "github.com/ory/hydra/sdk/go/hydra/client"
  admin "github.com/ory/hydra/sdk/go/hydra/client/admin"
  "github.com/ory/hydra/sdk/go/hydra/models"
)

type Server struct {
  logger *logrus.Logger
  config *config.Config
}

type LoginPage struct {
  Title string
  PrefillUser string
  PrefillPassword string
  LoginFailed bool
  LoginError bool
  Challenge string
}

type ConsentPage struct {
  Title string
  User string
  Challenge string
  ClientName string
  ClientId string
  Scope []string
}

func New(c *config.Config) *Server {
  l := logrus.New()
  l.SetLevel(logrus.DebugLevel)

  return &Server{
    logger: l,
    config: c,
  }
}

func (s *Server) Start() {

  http.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request) {

    logFields := make(map[string]interface{})
    logFields["path"] = r.URL.String()

    s.logger.WithFields(logFields).Infof("Request received")

    switch(r.Method) { 
    case "GET":
      s.logger.Infof("GET Request received")

      t := template.Must(template.New("login").Parse(htmltemplates.Login))

      lp := &LoginPage{
        PrefillUser: s.config.PrefillUser,
        PrefillPassword: s.config.PrefillPassword,
        Challenge: r.URL.Query().Get("login_challenge"),
      }

      err := t.Execute(w, lp)

      if err != nil {
        s.logger.Errorf("Error rendering network page: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
      }

    case "POST":
      s.logger.Infof("POST Request received")

      r.ParseForm()

      username, usernamePresent := r.Form["username"]
      password, passwordPresent := r.Form["password"]

      authSuccess := true
      var authErr error

      if usernamePresent && passwordPresent {
        if s.config.AuthFunc != nil {
          authSuccess, authErr = s.config.AuthFunc(username[0], password[0]) 
        }
      }

      if authErr != nil {
        s.logger.Error("AuthErr credentials")

        t := template.Must(template.New("login").Parse(htmltemplates.Login))

        lp := &LoginPage{
          PrefillUser: s.config.PrefillUser,
          PrefillPassword: s.config.PrefillPassword,
          LoginError: true,
          Challenge: r.URL.Query().Get("challenge"),
        }

        err := t.Execute(w, lp)

        if err != nil {
          s.logger.Errorf("Error rendering network page: %s", err)
          w.WriteHeader(http.StatusInternalServerError)
          return
        }
        return
      }

      if !authSuccess {
        s.logger.Error("Incorrect credentials")

        t := template.Must(template.New("login").Parse(htmltemplates.Login))

        lp := &LoginPage{
          PrefillUser: s.config.PrefillUser,
          PrefillPassword: s.config.PrefillPassword,
          LoginFailed: true,
          Challenge: r.URL.Query().Get("challenge"),
        }

        err := t.Execute(w, lp)

        if err != nil {
          s.logger.Errorf("Error rendering network page: %s", err)
          w.WriteHeader(http.StatusInternalServerError)
          return
        }
        return
      }

      s.logger.Infof("User '%s' has been authenticated", username[0])

      var challengeValue string
      challengeFromForm, challengeInForm := r.Form["challenge"]
      if challengeInForm {
        challengeValue = challengeFromForm[0]
      }
      s.logger.Debugf("Challenge value: '%s'", challengeValue)

      httpclient := &http.Client{}
      if s.config.SkipSSLCheck {
        httpclient.Transport = &http.Transport{
          TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
      }

      trans := &hydraclient.TransportConfig{
        Host: s.config.GetHydraAdminHostname(),
        BasePath: s.config.HydraAdminBasePath,
        Schemes: []string{"https"},
      }
      hydra := hydraclient.NewHTTPClientWithConfig(nil, trans)

      getLoginRequestParamsRequest := admin.NewGetLoginRequestParamsWithHTTPClient(httpclient)
      getLoginRequestParamsRequest.SetTimeout( 10 * time.Second)
      getLoginRequestParamsRequest.LoginChallenge = challengeValue

      resp, err := hydra.Admin.GetLoginRequest(getLoginRequestParamsRequest)
      if err != nil {
        s.logger.Errorf("Error getting login request details: %s", err)
        fmt.Fprintf(w, err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
      }
      s.logger.Debugf("Response object getting login request: %+v", resp.Payload)

      loginOKRequest := admin.NewAcceptLoginRequestParamsWithHTTPClient(httpclient)

      b := &models.HandledLoginRequest{
        Subject: &username[0], 
      }

      loginOKRequest.SetBody(b)
      loginOKRequest.SetTimeout(10 * time.Second )
      loginOKRequest.LoginChallenge = resp.Payload.Challenge

      s.logger.Debugf("Sending accept login okay")

      loginOKResponse, err := hydra.Admin.AcceptLoginRequest(loginOKRequest) 
      if err != nil {
        s.logger.Errorf("Error getting login ok: %s", err)
        fmt.Fprintf(w, err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
      }

      //now send the redirect we just received
      http.Redirect(w, r, loginOKResponse.Payload.RedirectTo, http.StatusFound)

    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
    }
  })

  http.HandleFunc("/consent", func (w http.ResponseWriter, r *http.Request) {

    logFields := make(map[string]interface{})
    logFields["path"] = r.URL.String()

    s.logger.WithFields(logFields).Infof("Request received")

    httpclient := &http.Client{}
    if s.config.SkipSSLCheck {
      httpclient.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
      }
    }

    trans := &hydraclient.TransportConfig{
      Host: s.config.GetHydraAdminHostname(),
      BasePath: s.config.HydraAdminBasePath,
      Schemes: []string{"https"},
    }
    hydra := hydraclient.NewHTTPClientWithConfig(nil, trans)

    switch(r.Method) { 
    case "GET":
      s.logger.Infof("GET Request received")

      t := template.Must(template.New("consent").Parse(htmltemplates.Consent))

      // get the consent request
      getConsentRequest := admin.NewGetConsentRequestParamsWithHTTPClient(httpclient)
      getConsentRequest.SetTimeout(10 * time.Second) 
      getConsentRequest.ConsentChallenge = r.URL.Query().Get("consent_challenge")

      getConsentRequestResponse, err := hydra.Admin.GetConsentRequest(getConsentRequest)
      if err != nil {
        s.logger.Errorf("Error getting login ok: %s", err)
        fmt.Fprintf(w, err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
      }

      lp := &ConsentPage{
        Title: "ConsentPage",
        User: getConsentRequestResponse.Payload.Subject,
        ClientName: getConsentRequestResponse.Payload.Client.Name,
        ClientId: getConsentRequestResponse.Payload.Client.ClientID,
        Challenge: r.URL.Query().Get("challenge"),
        Scope: getConsentRequestResponse.Payload.RequestedScope,
      }

      err = t.Execute(w, lp)

      if err != nil {
        s.logger.Errorf("Error rendering network page: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
      }

    case "POST":
      s.logger.Infof("POST Request received")

      r.ParseForm()

      allowed, found := r.Form["userconsent"]
      if !found {
        s.logger.Errorf("Variable 'userconsent' not found in http post so form must have been tampered with")
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Bad request!")
        return
      } 

      switch allowed[0] {
      case "accept":
        r.ParseForm()

        s.logger.Debugf("Handling that user accepted the scope")

        s.logger.Debugf("Getting the consent request")

        getConsentRequest := admin.NewGetConsentRequestParamsWithHTTPClient(httpclient)
        getConsentRequest.SetTimeout(10 * time.Second) 
        getConsentRequest.ConsentChallenge = r.URL.Query().Get("consent_challenge")

        getConsentRequestResponse, err := hydra.Admin.GetConsentRequest(getConsentRequest)
        if err != nil {
          s.logger.Errorf("Error getting the consent request: %s", err)
          fmt.Fprintf(w, err.Error())
          w.WriteHeader(http.StatusInternalServerError)
          return
        }

        s.logger.Debugf("Preparing the accept consent request")

        consentOKRequest := admin.NewAcceptConsentRequestParamsWithHTTPClient(httpclient)

        var grantedScope []string

        for _, requestedScope := range getConsentRequestResponse.Payload.RequestedScope {
          if val, found := r.Form[fmt.Sprintf("scope_%s", requestedScope)]; found && strings.ToLower(val[0]) == "checked" {
            grantedScope = append(grantedScope, requestedScope)
          }
        }

        s.logger.Debugf("RequestedScope: %s", getConsentRequestResponse.Payload.RequestedScope) 
        s.logger.Debugf("GrantedScope: %s", grantedScope) 

        b := &models.HandledConsentRequest{
          GrantedScope: grantedScope,
          GrantedAudience: getConsentRequestResponse.Payload.RequestedAudience,
        }

        consentOKRequest.SetBody(b)
        consentOKRequest.SetTimeout(10 * time.Second )
        consentOKRequest.ConsentChallenge = r.URL.Query().Get("consent_challenge")

        consentOKResponse, err := hydra.Admin.AcceptConsentRequest(consentOKRequest) 
        if err != nil {
          s.logger.Errorf("Error getting login ok: %s", err)
          fmt.Fprintf(w, err.Error())
          w.WriteHeader(http.StatusInternalServerError)
          return
        }
        s.logger.Debugf("Succesfully called the accept consent endpoint")

        //now send the redirect we just received
        http.Redirect(w, r, consentOKResponse.Payload.RedirectTo, http.StatusFound)
        return

      case "reject":
        consentDeniedRequest := admin.NewRejectConsentRequestParamsWithHTTPClient(httpclient)

        b := &models.RequestDeniedError{
          Name: "access_denied",
          Description: "The resource owner denied the request",
        }

        consentDeniedRequest.SetBody(b)
        consentDeniedRequest.SetTimeout(10 * time.Second )
        consentDeniedRequest.ConsentChallenge = r.URL.Query().Get("consent_challenge")

        consentDenyResponse, err := hydra.Admin.RejectConsentRequest(consentDeniedRequest)
        if err != nil {
          s.logger.Errorf("Error submitting a deny consent: %s", err)
          fmt.Fprintf(w, err.Error())
          w.WriteHeader(http.StatusInternalServerError)
          return
        }
        s.logger.Debugf("Succesfully called the deny consent endpoint")

        //now send the redirect we just received
        http.Redirect(w, r, consentDenyResponse.Payload.RedirectTo, http.StatusFound)
        return

      default:
        s.logger.Errorf("Variable 'userconsent' not found in http post so form must have been tampered with")
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Bad request!")
        return
      }

    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
    }
  })
 
  s.logger.Infof("Starting server on port %d", s.config.Port)

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), nil))
}
