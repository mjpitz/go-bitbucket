// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetTeamsUsernameHooksReader is a Reader for the GetTeamsUsernameHooks structure.
type GetTeamsUsernameHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamsUsernameHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamsUsernameHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamsUsernameHooksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamsUsernameHooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamsUsernameHooksOK creates a GetTeamsUsernameHooksOK with default headers values
func NewGetTeamsUsernameHooksOK() *GetTeamsUsernameHooksOK {
	return &GetTeamsUsernameHooksOK{}
}

/*GetTeamsUsernameHooksOK handles this case with default header values.

The paginated list of installed webhooks.
*/
type GetTeamsUsernameHooksOK struct {
	Payload *models.PaginatedWebhookSubscriptions
}

func (o *GetTeamsUsernameHooksOK) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/hooks][%d] getTeamsUsernameHooksOK  %+v", 200, o.Payload)
}

func (o *GetTeamsUsernameHooksOK) GetPayload() *models.PaginatedWebhookSubscriptions {
	return o.Payload
}

func (o *GetTeamsUsernameHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedWebhookSubscriptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamsUsernameHooksForbidden creates a GetTeamsUsernameHooksForbidden with default headers values
func NewGetTeamsUsernameHooksForbidden() *GetTeamsUsernameHooksForbidden {
	return &GetTeamsUsernameHooksForbidden{}
}

/*GetTeamsUsernameHooksForbidden handles this case with default header values.

If the authenticated user is not an admin on the specified team.
*/
type GetTeamsUsernameHooksForbidden struct {
	Payload *models.Error
}

func (o *GetTeamsUsernameHooksForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/hooks][%d] getTeamsUsernameHooksForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamsUsernameHooksForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTeamsUsernameHooksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamsUsernameHooksNotFound creates a GetTeamsUsernameHooksNotFound with default headers values
func NewGetTeamsUsernameHooksNotFound() *GetTeamsUsernameHooksNotFound {
	return &GetTeamsUsernameHooksNotFound{}
}

/*GetTeamsUsernameHooksNotFound handles this case with default header values.

If the specified team does not exist.
*/
type GetTeamsUsernameHooksNotFound struct {
	Payload *models.Error
}

func (o *GetTeamsUsernameHooksNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/hooks][%d] getTeamsUsernameHooksNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamsUsernameHooksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTeamsUsernameHooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
