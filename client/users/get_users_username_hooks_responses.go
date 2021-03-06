// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetUsersUsernameHooksReader is a Reader for the GetUsersUsernameHooks structure.
type GetUsersUsernameHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUsernameHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUsernameHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUsersUsernameHooksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUsernameHooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersUsernameHooksOK creates a GetUsersUsernameHooksOK with default headers values
func NewGetUsersUsernameHooksOK() *GetUsersUsernameHooksOK {
	return &GetUsersUsernameHooksOK{}
}

/*GetUsersUsernameHooksOK handles this case with default header values.

The paginated list of installed webhooks.
*/
type GetUsersUsernameHooksOK struct {
	Payload *models.PaginatedWebhookSubscriptions
}

func (o *GetUsersUsernameHooksOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/hooks][%d] getUsersUsernameHooksOK  %+v", 200, o.Payload)
}

func (o *GetUsersUsernameHooksOK) GetPayload() *models.PaginatedWebhookSubscriptions {
	return o.Payload
}

func (o *GetUsersUsernameHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedWebhookSubscriptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUsernameHooksForbidden creates a GetUsersUsernameHooksForbidden with default headers values
func NewGetUsersUsernameHooksForbidden() *GetUsersUsernameHooksForbidden {
	return &GetUsersUsernameHooksForbidden{}
}

/*GetUsersUsernameHooksForbidden handles this case with default header values.

If the authenticated user is accessing an account other than their own.
*/
type GetUsersUsernameHooksForbidden struct {
	Payload *models.Error
}

func (o *GetUsersUsernameHooksForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{username}/hooks][%d] getUsersUsernameHooksForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersUsernameHooksForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsersUsernameHooksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUsernameHooksNotFound creates a GetUsersUsernameHooksNotFound with default headers values
func NewGetUsersUsernameHooksNotFound() *GetUsersUsernameHooksNotFound {
	return &GetUsersUsernameHooksNotFound{}
}

/*GetUsersUsernameHooksNotFound handles this case with default header values.

If the specified account does not exist.
*/
type GetUsersUsernameHooksNotFound struct {
	Payload *models.Error
}

func (o *GetUsersUsernameHooksNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/hooks][%d] getUsersUsernameHooksNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersUsernameHooksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsersUsernameHooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
