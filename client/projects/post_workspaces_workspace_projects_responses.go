// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// PostWorkspacesWorkspaceProjectsReader is a Reader for the PostWorkspacesWorkspaceProjects structure.
type PostWorkspacesWorkspaceProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkspacesWorkspaceProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWorkspacesWorkspaceProjectsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostWorkspacesWorkspaceProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkspacesWorkspaceProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkspacesWorkspaceProjectsCreated creates a PostWorkspacesWorkspaceProjectsCreated with default headers values
func NewPostWorkspacesWorkspaceProjectsCreated() *PostWorkspacesWorkspaceProjectsCreated {
	return &PostWorkspacesWorkspaceProjectsCreated{}
}

/*PostWorkspacesWorkspaceProjectsCreated handles this case with default header values.

A new project has been created.
*/
type PostWorkspacesWorkspaceProjectsCreated struct {
	/*The location of the newly created project
	 */
	Location string

	Payload *models.Project
}

func (o *PostWorkspacesWorkspaceProjectsCreated) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/projects][%d] postWorkspacesWorkspaceProjectsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkspacesWorkspaceProjectsCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceProjectsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkspacesWorkspaceProjectsForbidden creates a PostWorkspacesWorkspaceProjectsForbidden with default headers values
func NewPostWorkspacesWorkspaceProjectsForbidden() *PostWorkspacesWorkspaceProjectsForbidden {
	return &PostWorkspacesWorkspaceProjectsForbidden{}
}

/*PostWorkspacesWorkspaceProjectsForbidden handles this case with default header values.

The requesting user isn't authorized to create the project.
*/
type PostWorkspacesWorkspaceProjectsForbidden struct {
	Payload *models.Error
}

func (o *PostWorkspacesWorkspaceProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/projects][%d] postWorkspacesWorkspaceProjectsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkspacesWorkspaceProjectsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkspacesWorkspaceProjectsNotFound creates a PostWorkspacesWorkspaceProjectsNotFound with default headers values
func NewPostWorkspacesWorkspaceProjectsNotFound() *PostWorkspacesWorkspaceProjectsNotFound {
	return &PostWorkspacesWorkspaceProjectsNotFound{}
}

/*PostWorkspacesWorkspaceProjectsNotFound handles this case with default header values.

A workspace doesn't exist at this location.
*/
type PostWorkspacesWorkspaceProjectsNotFound struct {
	Payload *models.Error
}

func (o *PostWorkspacesWorkspaceProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /workspaces/{workspace}/projects][%d] postWorkspacesWorkspaceProjectsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkspacesWorkspaceProjectsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWorkspacesWorkspaceProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
