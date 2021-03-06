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

// PostTeamsUsernameProjectsReader is a Reader for the PostTeamsUsernameProjects structure.
type PostTeamsUsernameProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTeamsUsernameProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostTeamsUsernameProjectsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostTeamsUsernameProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTeamsUsernameProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTeamsUsernameProjectsCreated creates a PostTeamsUsernameProjectsCreated with default headers values
func NewPostTeamsUsernameProjectsCreated() *PostTeamsUsernameProjectsCreated {
	return &PostTeamsUsernameProjectsCreated{}
}

/*PostTeamsUsernameProjectsCreated handles this case with default header values.

A new project has been created.
*/
type PostTeamsUsernameProjectsCreated struct {
	/*The location of the newly created project
	 */
	Location string

	Payload *models.Project
}

func (o *PostTeamsUsernameProjectsCreated) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/projects/][%d] postTeamsUsernameProjectsCreated  %+v", 201, o.Payload)
}

func (o *PostTeamsUsernameProjectsCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *PostTeamsUsernameProjectsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsUsernameProjectsForbidden creates a PostTeamsUsernameProjectsForbidden with default headers values
func NewPostTeamsUsernameProjectsForbidden() *PostTeamsUsernameProjectsForbidden {
	return &PostTeamsUsernameProjectsForbidden{}
}

/*PostTeamsUsernameProjectsForbidden handles this case with default header values.

The requesting user isn't authorized to create the project.
*/
type PostTeamsUsernameProjectsForbidden struct {
	Payload *models.Error
}

func (o *PostTeamsUsernameProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/projects/][%d] postTeamsUsernameProjectsForbidden  %+v", 403, o.Payload)
}

func (o *PostTeamsUsernameProjectsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTeamsUsernameProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsUsernameProjectsNotFound creates a PostTeamsUsernameProjectsNotFound with default headers values
func NewPostTeamsUsernameProjectsNotFound() *PostTeamsUsernameProjectsNotFound {
	return &PostTeamsUsernameProjectsNotFound{}
}

/*PostTeamsUsernameProjectsNotFound handles this case with default header values.

A team doesn't exist at this location.
*/
type PostTeamsUsernameProjectsNotFound struct {
	Payload *models.Error
}

func (o *PostTeamsUsernameProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/projects/][%d] postTeamsUsernameProjectsNotFound  %+v", 404, o.Payload)
}

func (o *PostTeamsUsernameProjectsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTeamsUsernameProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
