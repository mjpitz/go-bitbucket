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

// PutTeamsUsernameHooksUIDReader is a Reader for the PutTeamsUsernameHooksUID structure.
type PutTeamsUsernameHooksUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTeamsUsernameHooksUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTeamsUsernameHooksUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutTeamsUsernameHooksUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTeamsUsernameHooksUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTeamsUsernameHooksUIDOK creates a PutTeamsUsernameHooksUIDOK with default headers values
func NewPutTeamsUsernameHooksUIDOK() *PutTeamsUsernameHooksUIDOK {
	return &PutTeamsUsernameHooksUIDOK{}
}

/*PutTeamsUsernameHooksUIDOK handles this case with default header values.

The webhook subscription object.
*/
type PutTeamsUsernameHooksUIDOK struct {
	Payload *models.WebhookSubscription
}

func (o *PutTeamsUsernameHooksUIDOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{username}/hooks/{uid}][%d] putTeamsUsernameHooksUidOK  %+v", 200, o.Payload)
}

func (o *PutTeamsUsernameHooksUIDOK) GetPayload() *models.WebhookSubscription {
	return o.Payload
}

func (o *PutTeamsUsernameHooksUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTeamsUsernameHooksUIDForbidden creates a PutTeamsUsernameHooksUIDForbidden with default headers values
func NewPutTeamsUsernameHooksUIDForbidden() *PutTeamsUsernameHooksUIDForbidden {
	return &PutTeamsUsernameHooksUIDForbidden{}
}

/*PutTeamsUsernameHooksUIDForbidden handles this case with default header values.

If the authenticated user does not have permission to update the webhook.
*/
type PutTeamsUsernameHooksUIDForbidden struct {
	Payload *models.Error
}

func (o *PutTeamsUsernameHooksUIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{username}/hooks/{uid}][%d] putTeamsUsernameHooksUidForbidden  %+v", 403, o.Payload)
}

func (o *PutTeamsUsernameHooksUIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutTeamsUsernameHooksUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTeamsUsernameHooksUIDNotFound creates a PutTeamsUsernameHooksUIDNotFound with default headers values
func NewPutTeamsUsernameHooksUIDNotFound() *PutTeamsUsernameHooksUIDNotFound {
	return &PutTeamsUsernameHooksUIDNotFound{}
}

/*PutTeamsUsernameHooksUIDNotFound handles this case with default header values.

If the webhook or team does not exist.
*/
type PutTeamsUsernameHooksUIDNotFound struct {
	Payload *models.Error
}

func (o *PutTeamsUsernameHooksUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{username}/hooks/{uid}][%d] putTeamsUsernameHooksUidNotFound  %+v", 404, o.Payload)
}

func (o *PutTeamsUsernameHooksUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutTeamsUsernameHooksUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
