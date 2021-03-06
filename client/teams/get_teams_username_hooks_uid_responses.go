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

// GetTeamsUsernameHooksUIDReader is a Reader for the GetTeamsUsernameHooksUID structure.
type GetTeamsUsernameHooksUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamsUsernameHooksUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamsUsernameHooksUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTeamsUsernameHooksUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamsUsernameHooksUIDOK creates a GetTeamsUsernameHooksUIDOK with default headers values
func NewGetTeamsUsernameHooksUIDOK() *GetTeamsUsernameHooksUIDOK {
	return &GetTeamsUsernameHooksUIDOK{}
}

/*GetTeamsUsernameHooksUIDOK handles this case with default header values.

The webhook subscription object.
*/
type GetTeamsUsernameHooksUIDOK struct {
	Payload *models.WebhookSubscription
}

func (o *GetTeamsUsernameHooksUIDOK) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/hooks/{uid}][%d] getTeamsUsernameHooksUidOK  %+v", 200, o.Payload)
}

func (o *GetTeamsUsernameHooksUIDOK) GetPayload() *models.WebhookSubscription {
	return o.Payload
}

func (o *GetTeamsUsernameHooksUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamsUsernameHooksUIDNotFound creates a GetTeamsUsernameHooksUIDNotFound with default headers values
func NewGetTeamsUsernameHooksUIDNotFound() *GetTeamsUsernameHooksUIDNotFound {
	return &GetTeamsUsernameHooksUIDNotFound{}
}

/*GetTeamsUsernameHooksUIDNotFound handles this case with default header values.

If the webhook or team does not exist.
*/
type GetTeamsUsernameHooksUIDNotFound struct {
	Payload *models.Error
}

func (o *GetTeamsUsernameHooksUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/hooks/{uid}][%d] getTeamsUsernameHooksUidNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamsUsernameHooksUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTeamsUsernameHooksUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
