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

// DeleteTeamsUsernameHooksUIDReader is a Reader for the DeleteTeamsUsernameHooksUID structure.
type DeleteTeamsUsernameHooksUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamsUsernameHooksUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTeamsUsernameHooksUIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteTeamsUsernameHooksUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamsUsernameHooksUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTeamsUsernameHooksUIDNoContent creates a DeleteTeamsUsernameHooksUIDNoContent with default headers values
func NewDeleteTeamsUsernameHooksUIDNoContent() *DeleteTeamsUsernameHooksUIDNoContent {
	return &DeleteTeamsUsernameHooksUIDNoContent{}
}

/*DeleteTeamsUsernameHooksUIDNoContent handles this case with default header values.

When the webhook was deleted successfully
*/
type DeleteTeamsUsernameHooksUIDNoContent struct {
}

func (o *DeleteTeamsUsernameHooksUIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{username}/hooks/{uid}][%d] deleteTeamsUsernameHooksUidNoContent ", 204)
}

func (o *DeleteTeamsUsernameHooksUIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTeamsUsernameHooksUIDForbidden creates a DeleteTeamsUsernameHooksUIDForbidden with default headers values
func NewDeleteTeamsUsernameHooksUIDForbidden() *DeleteTeamsUsernameHooksUIDForbidden {
	return &DeleteTeamsUsernameHooksUIDForbidden{}
}

/*DeleteTeamsUsernameHooksUIDForbidden handles this case with default header values.

If the authenticated user does not have permission to delete the webhook.
*/
type DeleteTeamsUsernameHooksUIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteTeamsUsernameHooksUIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{username}/hooks/{uid}][%d] deleteTeamsUsernameHooksUidForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamsUsernameHooksUIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTeamsUsernameHooksUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamsUsernameHooksUIDNotFound creates a DeleteTeamsUsernameHooksUIDNotFound with default headers values
func NewDeleteTeamsUsernameHooksUIDNotFound() *DeleteTeamsUsernameHooksUIDNotFound {
	return &DeleteTeamsUsernameHooksUIDNotFound{}
}

/*DeleteTeamsUsernameHooksUIDNotFound handles this case with default header values.

If the webhook or team does not exist.
*/
type DeleteTeamsUsernameHooksUIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteTeamsUsernameHooksUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{username}/hooks/{uid}][%d] deleteTeamsUsernameHooksUidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamsUsernameHooksUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTeamsUsernameHooksUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}