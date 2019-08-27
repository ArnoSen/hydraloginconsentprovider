// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SwaggerOAuthIntrospectionRequest swagger o auth introspection request
// swagger:model swaggerOAuthIntrospectionRequest
type SwaggerOAuthIntrospectionRequest struct {

	// An optional, space separated list of required scopes. If the access token was not granted one of the
	// scopes, the result of active will be false.
	//
	// in: formData
	Scope string `json:"scope,omitempty"`

	// The string value of the token. For access tokens, this
	// is the "access_token" value returned from the token endpoint
	// defined in OAuth 2.0. For refresh tokens, this is the "refresh_token"
	// value returned.
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this swagger o auth introspection request
func (m *SwaggerOAuthIntrospectionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwaggerOAuthIntrospectionRequest) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerOAuthIntrospectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerOAuthIntrospectionRequest) UnmarshalBinary(b []byte) error {
	var res SwaggerOAuthIntrospectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
