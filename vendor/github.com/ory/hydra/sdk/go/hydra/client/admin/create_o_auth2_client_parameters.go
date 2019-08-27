// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// NewCreateOAuth2ClientParams creates a new CreateOAuth2ClientParams object
// with the default values initialized.
func NewCreateOAuth2ClientParams() *CreateOAuth2ClientParams {
	var ()
	return &CreateOAuth2ClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOAuth2ClientParamsWithTimeout creates a new CreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOAuth2ClientParamsWithTimeout(timeout time.Duration) *CreateOAuth2ClientParams {
	var ()
	return &CreateOAuth2ClientParams{

		timeout: timeout,
	}
}

// NewCreateOAuth2ClientParamsWithContext creates a new CreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOAuth2ClientParamsWithContext(ctx context.Context) *CreateOAuth2ClientParams {
	var ()
	return &CreateOAuth2ClientParams{

		Context: ctx,
	}
}

// NewCreateOAuth2ClientParamsWithHTTPClient creates a new CreateOAuth2ClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOAuth2ClientParamsWithHTTPClient(client *http.Client) *CreateOAuth2ClientParams {
	var ()
	return &CreateOAuth2ClientParams{
		HTTPClient: client,
	}
}

/*CreateOAuth2ClientParams contains all the parameters to send to the API endpoint
for the create o auth2 client operation typically these are written to a http.Request
*/
type CreateOAuth2ClientParams struct {

	/*Body*/
	Body *models.Client

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create o auth2 client params
func (o *CreateOAuth2ClientParams) WithTimeout(timeout time.Duration) *CreateOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create o auth2 client params
func (o *CreateOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create o auth2 client params
func (o *CreateOAuth2ClientParams) WithContext(ctx context.Context) *CreateOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create o auth2 client params
func (o *CreateOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create o auth2 client params
func (o *CreateOAuth2ClientParams) WithHTTPClient(client *http.Client) *CreateOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create o auth2 client params
func (o *CreateOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create o auth2 client params
func (o *CreateOAuth2ClientParams) WithBody(body *models.Client) *CreateOAuth2ClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create o auth2 client params
func (o *CreateOAuth2ClientParams) SetBody(body *models.Client) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
