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

// GetWorkspacesWorkspaceHooksReader is a Reader for the GetWorkspacesWorkspaceHooks structure.
type GetWorkspacesWorkspaceHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspacesWorkspaceHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspacesWorkspaceHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetWorkspacesWorkspaceHooksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkspacesWorkspaceHooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkspacesWorkspaceHooksOK creates a GetWorkspacesWorkspaceHooksOK with default headers values
func NewGetWorkspacesWorkspaceHooksOK() *GetWorkspacesWorkspaceHooksOK {
	return &GetWorkspacesWorkspaceHooksOK{}
}

/*GetWorkspacesWorkspaceHooksOK handles this case with default header values.

The paginated list of installed webhooks.
*/
type GetWorkspacesWorkspaceHooksOK struct {
	Payload *models.PaginatedWebhookSubscriptions
}

func (o *GetWorkspacesWorkspaceHooksOK) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspace}/hooks][%d] getWorkspacesWorkspaceHooksOK  %+v", 200, o.Payload)
}

func (o *GetWorkspacesWorkspaceHooksOK) GetPayload() *models.PaginatedWebhookSubscriptions {
	return o.Payload
}

func (o *GetWorkspacesWorkspaceHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedWebhookSubscriptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacesWorkspaceHooksForbidden creates a GetWorkspacesWorkspaceHooksForbidden with default headers values
func NewGetWorkspacesWorkspaceHooksForbidden() *GetWorkspacesWorkspaceHooksForbidden {
	return &GetWorkspacesWorkspaceHooksForbidden{}
}

/*GetWorkspacesWorkspaceHooksForbidden handles this case with default header values.

If the authenticated user is not an owner on the specified workspace.
*/
type GetWorkspacesWorkspaceHooksForbidden struct {
	Payload *models.Error
}

func (o *GetWorkspacesWorkspaceHooksForbidden) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspace}/hooks][%d] getWorkspacesWorkspaceHooksForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkspacesWorkspaceHooksForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkspacesWorkspaceHooksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacesWorkspaceHooksNotFound creates a GetWorkspacesWorkspaceHooksNotFound with default headers values
func NewGetWorkspacesWorkspaceHooksNotFound() *GetWorkspacesWorkspaceHooksNotFound {
	return &GetWorkspacesWorkspaceHooksNotFound{}
}

/*GetWorkspacesWorkspaceHooksNotFound handles this case with default header values.

If the specified workspace does not exist.
*/
type GetWorkspacesWorkspaceHooksNotFound struct {
	Payload *models.Error
}

func (o *GetWorkspacesWorkspaceHooksNotFound) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspace}/hooks][%d] getWorkspacesWorkspaceHooksNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkspacesWorkspaceHooksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkspacesWorkspaceHooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}