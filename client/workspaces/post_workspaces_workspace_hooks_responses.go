// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// PostWorkspacesWorkspaceHooksReader is a Reader for the PostWorkspacesWorkspaceHooks structure.
type PostWorkspacesWorkspaceHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkspacesWorkspaceHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWorkspacesWorkspaceHooksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostWorkspacesWorkspaceHooksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkspacesWorkspaceHooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkspacesWorkspaceHooksCreated creates a PostWorkspacesWorkspaceHooksCreated with default headers values
func NewPostWorkspacesWorkspaceHooksCreated() *PostWorkspacesWorkspaceHooksCreated {
	return &PostWorkspacesWorkspaceHooksCreated{}
}

/*PostWorkspacesWorkspaceHooksCreated handles this case with default header values.

The newly installed webhook.
*/
type PostWorkspacesWorkspaceHooksCreated struct {
	/*The URL of new newly created webhook.
	 */
	Location string

	Payload *models.WebhookSubscription
}

func (o *PostWorkspacesWorkspaceHooksCreated) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/hooks][%d] postWorkspacesWorkspaceHooksCreated  %+v", 201, o.Payload)
}

func (o *PostWorkspacesWorkspaceHooksCreated) GetPayload() *models.WebhookSubscription {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceHooksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.WebhookSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkspacesWorkspaceHooksForbidden creates a PostWorkspacesWorkspaceHooksForbidden with default headers values
func NewPostWorkspacesWorkspaceHooksForbidden() *PostWorkspacesWorkspaceHooksForbidden {
	return &PostWorkspacesWorkspaceHooksForbidden{}
}

/*PostWorkspacesWorkspaceHooksForbidden handles this case with default header values.

If the authenticated user is not an owner on the specified workspace.
*/
type PostWorkspacesWorkspaceHooksForbidden struct {
	Payload *models.Error
}

func (o *PostWorkspacesWorkspaceHooksForbidden) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/hooks][%d] postWorkspacesWorkspaceHooksForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkspacesWorkspaceHooksForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceHooksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkspacesWorkspaceHooksNotFound creates a PostWorkspacesWorkspaceHooksNotFound with default headers values
func NewPostWorkspacesWorkspaceHooksNotFound() *PostWorkspacesWorkspaceHooksNotFound {
	return &PostWorkspacesWorkspaceHooksNotFound{}
}

/*PostWorkspacesWorkspaceHooksNotFound handles this case with default header values.

If the specified workspace does not exist.
*/
type PostWorkspacesWorkspaceHooksNotFound struct {
	Payload *models.Error
}

func (o *PostWorkspacesWorkspaceHooksNotFound) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/hooks][%d] postWorkspacesWorkspaceHooksNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkspacesWorkspaceHooksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceHooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
