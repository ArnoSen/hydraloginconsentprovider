package loginconsentclient

import (
  "time"
  jose "gopkg.in/square/go-jose.v2"
)

type RequestHandlerResponse struct {
  // RedirectURL is the URL which you should redirect the user to once the authentication process is completed.
  RedirectTo string `json:"redirect_to"`
}

type RequestDeniedError struct {
  Name        string `json:"error"`
  Description string `json:"error_description"`
  Hint        string `json:"error_hint,omitempty"`
  Code        int    `json:"status_code,omitempty"`
  Debug       string `json:"error_debug,omitempty"`
}

type HandledConsentRequest struct {
  // GrantScope sets the scope the user authorized the client to use. Should be a subset of `requested_scope`.
  GrantedScope []string `json:"grant_scope"`

  // GrantedAudience sets the audience the user authorized the client to use. Should be a subset of `requested_access_token_audience`.
  GrantedAudience []string `json:"grant_access_token_audience"`

  // Session allows you to set (optional) session data for access and ID tokens.
  Session *ConsentRequestSessionData `json:"session"`

  // Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same
  // client asks the same user for the same, or a subset of, scope.
  Remember bool `json:"remember"`

  // RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the
  // authorization will be remembered indefinitely.
  RememberFor int `json:"remember_for"`

  ConsentRequest  *ConsentRequest     `json:"-"`
  Error           *RequestDeniedError `json:"-"`
  Challenge       string              `json:"-"`
  RequestedAt     time.Time           `json:"-"`
  AuthenticatedAt time.Time           `json:"-"`
  WasUsed         bool                `json:"-"`
}

type PreviousConsentSession struct {
  // GrantScope sets the scope the user authorized the client to use. Should be a subset of `requested_scope`
  GrantedScope []string `json:"grant_scope"`

  // GrantedAudience sets the audience the user authorized the client to use. Should be a subset of `requested_access_token_audience`.
  GrantedAudience []string `json:"grant_access_token_audience"`

  // Session allows you to set (optional) session data for access and ID tokens.
  Session *ConsentRequestSessionData `json:"session"`

  // Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same
  // client asks the same user for the same, or a subset of, scope.
  Remember bool `json:"remember"`

  // RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the
  // authorization will be remembered indefinitely.
  RememberFor int `json:"remember_for"`

  ConsentRequest  *ConsentRequest     `json:"consent_request"`
}

type HandledLoginRequest struct {
  // Remember, if set to true, tells ORY Hydra to remember this user by telling the user agent (browser) to store
  // a cookie with authentication data. If the same user performs another OAuth 2.0 Authorization Request, he/she
  // will not be asked to log in again.
  Remember bool `json:"remember"`

  // RememberFor sets how long the authentication should be remembered for in seconds. If set to `0`, the
  // authorization will be remembered indefinitely.
  RememberFor int `json:"remember_for"`

  // ACR sets the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it
  // to express that, for example, a user authenticated using two factor authentication.
  ACR string `json:"acr"`

  // Subject is the user ID of the end-user that authenticated.
  // required: true
  Subject string `json:"subject"`

  // ForceSubjectIdentifier forces the "pairwise" user ID of the end-user that authenticated. The "pairwise" user ID refers to the
  // (Pairwise Identifier Algorithm)[http://openid.net/specs/openid-connect-core-1_0.html#PairwiseAlg] of the OpenID
  // Connect specification. It allows you to set an obfuscated subject ("user") identifier that is unique to the client.
  //
  // Please note that this changes the user ID on endpoint /userinfo and sub claim of the ID Token. It does not change the
  // sub claim in the OAuth 2.0 Introspection.
  //
  // Per default, ORY Hydra handles this value with its own algorithm. In case you want to set this yourself
  // you can use this field. Please note that setting this field has no effect if `pairwise` is not configured in
  // ORY Hydra or the OAuth 2.0 Client does not expect a pairwise identifier (set via `subject_type` key in the client's
  // configuration).
  //
  // Please also be aware that ORY Hydra is unable to properly compute this value during authentication. This implies
  // that you have to compute this value on every authentication process (probably depending on the client ID or some
  // other unique value).
  //
  // If you fail to compute the proper value, then authentication processes which have id_token_hint set might fail.
  ForceSubjectIdentifier string `json:"force_subject_identifier"`

  // Context is an optional object which can hold arbitrary data. The data will be made available when fetching the
  // consent request under the "context" field. This is useful in scenarios where login and consent endpoints share
  // data.
  Context map[string]interface{} `json:"context"`
}

type LogoutRequest struct {
  // Subject is the user for whom the logout was request.
  Subject string `json:"subject"`

  // SessionID is the login session ID that was requested to log out.
  SessionID string `json:"sid,omitempty"`

  // RequestURL is the original Logout URL requested.
  RequestURL string `json:"request_url"`

  // RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.
  RPInitiated bool `json:"rp_initiated"`
}

type LoginRequest struct {
  // Challenge is the identifier ("login challenge") of the login request. It is used to
  // identify the session.
  Challenge string `json:"challenge"`

  // RequestedScope contains the OAuth 2.0 Scope requested by the OAuth 2.0 Client.
  RequestedScope []string `json:"requested_scope"`

  // RequestedScope contains the access token audience as requested by the OAuth 2.0 Client.
  RequestedAudience []string `json:"requested_access_token_audience"`

  // Skip, if true, implies that the client has requested the same scopes from the same user previously.
  // If true, you can skip asking the user to grant the requested scopes, and simply forward the user to the redirect URL.
  //
  // This feature allows you to update / set session information.
  Skip bool `json:"skip"`

  // Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope
  // requested by the OAuth 2.0 client. If this value is set and `skip` is true, you MUST include this subject type
  // when accepting the login request, or the request will fail.
  Subject string `json:"subject"`

  // OpenIDConnectContext provides context for the (potential) OpenID Connect context. Implementation of these
  // values in your app are optional but can be useful if you want to be fully compliant with the OpenID Connect spec.
  OpenIDConnectContext *OpenIDConnectContext `json:"oidc_context"`

  // Client is the OAuth 2.0 Client that initiated the request.
  Client *OAuthClient `json:"client"`

  // RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which
  // initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but
  // might come in handy if you want to deal with additional request parameters.
  RequestURL string `json:"request_url"`

  // SessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag)
  // this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false)
  // this will be a new random value. This value is used as the "sid" parameter in the ID Token and in OIDC Front-/Back-
  // channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
  SessionID string `json:"session_id"`
}

type ConsentRequest struct {
  // Challenge is the identifier ("authorization challenge") of the consent authorization request. It is used to
  // identify the session.
  Challenge string `json:"challenge"`

  // RequestedScope contains the OAuth 2.0 Scope requested by the OAuth 2.0 Client.
  RequestedScope []string `json:"requested_scope"`

  // RequestedScope contains the access token audience as requested by the OAuth 2.0 Client.
  RequestedAudience []string `json:"requested_access_token_audience"`

  // Skip, if true, implies that the client has requested the same scopes from the same user previously.
  // If true, you must not ask the user to grant the requested scopes. You must however either allow or deny the
  // consent request using the usual API call.
  Skip bool `json:"skip"`

  // Subject is the user ID of the end-user that authenticated. Now, that end user needs to grant or deny the scope
  // requested by the OAuth 2.0 client.
  Subject string `json:"subject"`

  // OpenIDConnectContext provides context for the (potential) OpenID Connect context. Implementation of these
  // values in your app are optional but can be useful if you want to be fully compliant with the OpenID Connect spec.
  OpenIDConnectContext *OpenIDConnectContext `json:"oidc_context"`

  // Client is the OAuth 2.0 Client that initiated the request.
  Client *OAuthClient `json:"client"`

  // RequestURL is the original OAuth 2.0 Authorization URL requested by the OAuth 2.0 client. It is the URL which
  // initiates the OAuth 2.0 Authorization Code or OAuth 2.0 Implicit flow. This URL is typically not needed, but
  // might come in handy if you want to deal with additional request parameters.
  RequestURL string `json:"request_url"`

  // LoginChallenge is the login challenge this consent challenge belongs to. It can be used to associate
  // a login and consent request in the login & consent app.
  LoginChallenge string `json:"login_challenge"`

  // LoginSessionID is the login session ID. If the user-agent reuses a login session (via cookie / remember flag)
  // this ID will remain the same. If the user-agent did not have an existing authentication session (e.g. remember is false)
  // this will be a new random value. This value is used as the "sid" parameter in the ID Token and in OIDC Front-/Back-
  // channel logout. It's value can generally be used to associate consecutive login requests by a certain user.
  LoginSessionID string `json:"login_session_id"`

  // ACR represents the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it
  // to express that, for example, a user authenticated using two factor authentication.
  ACR string `json:"acr"`

  // Context contains arbitrary information set by the login endpoint or is empty if not set.
  Context map[string]interface{} `json:"context,omitempty"`
}

type ConsentRequestSessionData struct {
  // AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the
  // refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection.
  // If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties
  // can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
  AccessToken map[string]interface{} `json:"access_token"`

  // IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable
  // by anyone that has access to the ID Challenge. Use with care!
  IDToken map[string]interface{} `json:"id_token"`

  //UserInfo map[string]interface{} `json:"userinfo"`
}

func NewConsentRequestSessionData() *ConsentRequestSessionData {
  return &ConsentRequestSessionData{
    AccessToken: map[string]interface{}{},
    IDToken:     map[string]interface{}{},
  }
}

type OpenIDConnectContext struct {
  // ACRValues is the Authentication AuthorizationContext Class Reference requested in the OAuth 2.0 Authorization request.
  // It is a parameter defined by OpenID Connect and expresses which level of authentication (e.g. 2FA) is required.
  //
  // OpenID Connect defines it as follows:
  // > Requested Authentication AuthorizationContext Class Reference values. Space-separated string that specifies the acr values
  // that the Authorization Server is being requested to use for processing this Authentication Request, with the
  // values appearing in order of preference. The Authentication AuthorizationContext Class satisfied by the authentication
  // performed is returned as the acr Claim Value, as specified in Section 2. The acr Claim is requested as a
  // Voluntary Claim by this parameter.
  ACRValues []string `json:"acr_values,omitempty"`

  // UILocales is the End-User'id preferred languages and scripts for the user interface, represented as a
  // space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. For instance, the value
  // "fr-CA fr en" represents a preference for French as spoken in Canada, then French (without a region designation),
  // followed by English (without a region designation). An error SHOULD NOT result if some or all of the requested
  // locales are not supported by the OpenID Provider.
  UILocales []string `json:"ui_locales,omitempty"`

  // Display is a string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User.
  // The defined values are:
  // - page: The Authorization Server SHOULD display the authentication and consent UI consistent with a full User Agent page view. If the display parameter is not specified, this is the default display mode.
  // - popup: The Authorization Server SHOULD display the authentication and consent UI consistent with a popup User Agent window. The popup User Agent window should be of an appropriate size for a login-focused dialog and should not obscure the entire window that it is popping up over.
  // - touch: The Authorization Server SHOULD display the authentication and consent UI consistent with a device that leverages a touch interface.
  // - wap: The Authorization Server SHOULD display the authentication and consent UI consistent with a "feature phone" type display.
  //
  // The Authorization Server MAY also attempt to detect the capabilities of the User Agent and present an appropriate display.
  Display string `json:"display,omitempty"`

  // IDTokenHintClaims are the claims of the ID Token previously issued by the Authorization Server being passed as a hint about the
  // End-User's current or past authenticated session with the Client.
  IDTokenHintClaims map[string]interface{} `json:"id_token_hint_claims,omitempty"`

  // LoginHint hints about the login identifier the End-User might use to log in (if necessary).
  // This hint can be used by an RP if it first asks the End-User for their e-mail address (or other identifier)
  // and then wants to pass that value as a hint to the discovered authorization service. This value MAY also be a
  // phone number in the format specified for the phone_number Claim. The use of this parameter is optional.
  LoginHint string `json:"login_hint,omitempty"`
}

type OAuthClient struct {
  // ClientID  is the id for this client.
  ClientID string `json:"client_id"`

  // Name is the human-readable string name of the client to be presented to the
  // end-user during authorization.
  Name string `json:"client_name"`

  // Secret is the client's secret. The secret will be included in the create request as cleartext, and then
  // never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users
  // that they need to write the secret down as it will not be made available again.
  Secret string `json:"client_secret,omitempty"`

  // RedirectURIs is an array of allowed redirect urls for the client, for example http://mydomain/oauth/callback .
  RedirectURIs []string `json:"redirect_uris"`

  // GrantTypes is an array of grant types the client is allowed to use.
  //
  // Pattern: client_credentials|authorization_code|implicit|refresh_token
  GrantTypes []string `json:"grant_types"`

  // ResponseTypes is an array of the OAuth 2.0 response type strings that the client can
  // use at the authorization endpoint.
  //
  // Pattern: id_token|code|token
  ResponseTypes []string `json:"response_types"`

  // Scope is a string containing a space-separated list of scope values (as
  // described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client
  // can use when requesting access tokens.
  //
  // Pattern: ([a-zA-Z0-9\.\*]+\s?)+
  Scope string `json:"scope"`

  // Audience is a whitelist defining the audiences this client is allowed to request tokens for. An audience limits
  // the applicability of an OAuth 2.0 Access Token to, for example, certain API endpoints. The value is a list
  // of URLs. URLs MUST NOT contain whitespaces.
  Audience []string `json:"audience"`

  // Owner is a string identifying the owner of the OAuth 2.0 Client.
  Owner string `json:"owner"`

  // PolicyURI is a URL string that points to a human-readable privacy policy document
  // that describes how the deployment organization collects, uses,
  // retains, and discloses personal data.
  PolicyURI string `json:"policy_uri"`

  // AllowedCORSOrigins are one or more URLs (scheme://host[:port]) which are allowed to make CORS requests
  // to the /oauth/token endpoint. If this array is empty, the sever's CORS origin configuration (`CORS_ALLOWED_ORIGINS`)
  // will be used instead. If this array is set, the allowed origins are appended to the server's CORS origin configuration.
  // Be aware that environment variable `CORS_ENABLED` MUST be set to `true` for this to work.
  AllowedCORSOrigins []string `json:"allowed_cors_origins"`

  // TermsOfServiceURI is a URL string that points to a human-readable terms of service
  // document for the client that describes a contractual relationship
  // between the end-user and the client that the end-user accepts when
  // authorizing the client.
  TermsOfServiceURI string `json:"tos_uri"`

  // ClientURI is an URL string of a web page providing information about the client.
  // If present, the server SHOULD display this URL to the end-user in
  // a clickable fashion.
  ClientURI string `json:"client_uri"`

  // LogoURI is an URL string that references a logo for the client.
  LogoURI string `json:"logo_uri"`

  // Contacts is a array of strings representing ways to contact people responsible
  // for this client, typically email addresses.
  Contacts []string `json:"contacts"`

  // SecretExpiresAt is an integer holding the time at which the client
  // secret will expire or 0 if it will not expire. The time is
  // represented as the number of seconds from 1970-01-01T00:00:00Z as
  // measured in UTC until the date/time of expiration.
  //
  // This feature is currently not supported and it's value will always
  // be set to 0.
  SecretExpiresAt int `json:"client_secret_expires_at"`

  // SubjectType requested for responses to this Client. The subject_types_supported Discovery parameter contains a
  // list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.
  SubjectType string `json:"subject_type"`

  // URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a
  // file with a single JSON array of redirect_uri values.
  SectorIdentifierURI string `json:"sector_identifier_uri,omitempty"`

  // URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains
  // the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the
  // Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing
  // and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced
  // JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both
  // signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used
  // to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST
  // match those in the certificate.
  JSONWebKeysURI string `json:"jwks_uri,omitempty"`

  // Client's JSON Web Key Set [JWK] document, passed by value. The semantics of the jwks parameter are the same as
  // the jwks_uri parameter, other than that the JWK Set is passed by value, rather than by reference. This parameter
  // is intended only to be used by Clients that, for some reason, are unable to use the jwks_uri parameter, for
  // instance, by native applications that might not have a location to host the contents of the JWK Set. If a Client
  // can use jwks_uri, it MUST NOT use jwks. One significant downside of jwks is that it does not enable key rotation
  // (which jwks_uri does, as described in Section 10 of OpenID Connect Core 1.0 [OpenID.Core]). The jwks_uri and jwks
  // parameters MUST NOT be used together.
  JSONWebKeys *jose.JSONWebKeySet `json:"jwks,omitempty"`

  // Requested Client Authentication method for the Token Endpoint. The options are client_secret_post,
  // client_secret_basic, private_key_jwt, and none.
  TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"`

  // Array of request_uri values that are pre-registered by the RP for use at the OP. Servers MAY cache the
  // contents of the files referenced by these URIs and not retrieve them at the time they are used in a request.
  // OPs can require that request_uri values used be pre-registered with the require_request_uri_registration
  // discovery parameter.
  RequestURIs []string `json:"request_uris,omitempty"`

  // JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects
  // from this Client MUST be rejected, if not signed with this algorithm.
  RequestObjectSigningAlgorithm string `json:"request_object_signing_alg,omitempty"`

  // JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT
  // [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims
  // as a UTF-8 encoded JSON object using the application/json content-type.
  UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg,omitempty"`

  // CreatedAt returns the timestamp of the client's creation.
  CreatedAt time.Time `json:"created_at,omitempty"`

  // UpdatedAt returns the timestamp of the last update.
  UpdatedAt time.Time `json:"updated_at,omitempty"`

  // RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query
  // parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the
  // request and to determine which of the potentially multiple sessions is to be logged out; if either is
  // included, both MUST be.
  FrontChannelLogoutURI string `json:"frontchannel_logout_uri,omitempty"`

  // Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be
  // included to identify the RP session with the OP when the frontchannel_logout_uri is used.
  // If omitted, the default value is false.
  FrontChannelLogoutSessionRequired bool `json:"frontchannel_logout_session_required,omitempty"`

  // Array of URLs supplied by the RP to which it MAY request that the End-User's User Agent be redirected using the
  // post_logout_redirect_uri parameter after a logout has been performed.
  PostLogoutRedirectURIs []string `json:"post_logout_redirect_uris,omitempty"`

  // RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.
  BackChannelLogoutURI string `json:"backchannel_logout_uri,omitempty"`

  // Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout
  // Token to identify the RP session with the OP when the backchannel_logout_uri is used.
  // If omitted, the default value is false.
  BackChannelLogoutSessionRequired bool `json:"backchannel_logout_session_required,omitempty"`
}
