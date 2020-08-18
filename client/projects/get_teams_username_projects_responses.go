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

// GetTeamsUsernameProjectsReader is a Reader for the GetTeamsUsernameProjects structure.
type GetTeamsUsernameProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamsUsernameProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamsUsernameProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamsUsernameProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamsUsernameProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamsUsernameProjectsOK creates a GetTeamsUsernameProjectsOK with default headers values
func NewGetTeamsUsernameProjectsOK() *GetTeamsUsernameProjectsOK {
	return &GetTeamsUsernameProjectsOK{}
}

/*GetTeamsUsernameProjectsOK handles this case with default header values.

A paginated list of projects that belong to the specified team.
*/
type GetTeamsUsernameProjectsOK struct {
	Payload *models.PaginatedProjects
}

func (o *GetTeamsUsernameProjectsOK) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/projects/][%d] getTeamsUsernameProjectsOK  %+v", 200, o.Payload)
}

func (o *GetTeamsUsernameProjectsOK) GetPayload() *models.PaginatedProjects {
	return o.Payload
}

func (o *GetTeamsUsernameProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedProjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamsUsernameProjectsForbidden creates a GetTeamsUsernameProjectsForbidden with default headers values
func NewGetTeamsUsernameProjectsForbidden() *GetTeamsUsernameProjectsForbidden {
	return &GetTeamsUsernameProjectsForbidden{}
}

/*GetTeamsUsernameProjectsForbidden handles this case with default header values.

The requesting user isn't authorized to read the list of projects for the specified team.
*/
type GetTeamsUsernameProjectsForbidden struct {
	Payload *models.Error
}

func (o *GetTeamsUsernameProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/projects/][%d] getTeamsUsernameProjectsForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamsUsernameProjectsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTeamsUsernameProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamsUsernameProjectsNotFound creates a GetTeamsUsernameProjectsNotFound with default headers values
func NewGetTeamsUsernameProjectsNotFound() *GetTeamsUsernameProjectsNotFound {
	return &GetTeamsUsernameProjectsNotFound{}
}

/*GetTeamsUsernameProjectsNotFound handles this case with default header values.

A team doesn't exist at this location.
*/
type GetTeamsUsernameProjectsNotFound struct {
	Payload *models.Error
}

func (o *GetTeamsUsernameProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/projects/][%d] getTeamsUsernameProjectsNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamsUsernameProjectsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTeamsUsernameProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
