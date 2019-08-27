// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// OauthAuthReader is a Reader for the OauthAuth structure.
type OauthAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OauthAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 302:
		result := NewOauthAuthFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOauthAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewOauthAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOauthAuthFound creates a OauthAuthFound with default headers values
func NewOauthAuthFound() *OauthAuthFound {
	return &OauthAuthFound{}
}

/*OauthAuthFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type OauthAuthFound struct {
}

func (o *OauthAuthFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth][%d] oauthAuthFound ", 302)
}

func (o *OauthAuthFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOauthAuthUnauthorized creates a OauthAuthUnauthorized with default headers values
func NewOauthAuthUnauthorized() *OauthAuthUnauthorized {
	return &OauthAuthUnauthorized{}
}

/*OauthAuthUnauthorized handles this case with default header values.

genericError
*/
type OauthAuthUnauthorized struct {
	Payload *models.GenericError
}

func (o *OauthAuthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth][%d] oauthAuthUnauthorized  %+v", 401, o.Payload)
}

func (o *OauthAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOauthAuthInternalServerError creates a OauthAuthInternalServerError with default headers values
func NewOauthAuthInternalServerError() *OauthAuthInternalServerError {
	return &OauthAuthInternalServerError{}
}

/*OauthAuthInternalServerError handles this case with default header values.

genericError
*/
type OauthAuthInternalServerError struct {
	Payload *models.GenericError
}

func (o *OauthAuthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth][%d] oauthAuthInternalServerError  %+v", 500, o.Payload)
}

func (o *OauthAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
