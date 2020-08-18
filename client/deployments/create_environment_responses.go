// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// CreateEnvironmentReader is a Reader for the CreateEnvironment structure.
type CreateEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEnvironmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEnvironmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnvironmentCreated creates a CreateEnvironmentCreated with default headers values
func NewCreateEnvironmentCreated() *CreateEnvironmentCreated {
	return &CreateEnvironmentCreated{}
}

/*CreateEnvironmentCreated handles this case with default header values.

The environment was created.
*/
type CreateEnvironmentCreated struct {
	/*The URL of the newly created environment.
	 */
	Location string

	Payload *models.DeploymentEnvironment
}

func (o *CreateEnvironmentCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/environments/][%d] createEnvironmentCreated  %+v", 201, o.Payload)
}

func (o *CreateEnvironmentCreated) GetPayload() *models.DeploymentEnvironment {
	return o.Payload
}

func (o *CreateEnvironmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.DeploymentEnvironment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentNotFound creates a CreateEnvironmentNotFound with default headers values
func NewCreateEnvironmentNotFound() *CreateEnvironmentNotFound {
	return &CreateEnvironmentNotFound{}
}

/*CreateEnvironmentNotFound handles this case with default header values.

The account or repository does not exist.
*/
type CreateEnvironmentNotFound struct {
	Payload *models.Error
}

func (o *CreateEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/environments/][%d] createEnvironmentNotFound  %+v", 404, o.Payload)
}

func (o *CreateEnvironmentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentConflict creates a CreateEnvironmentConflict with default headers values
func NewCreateEnvironmentConflict() *CreateEnvironmentConflict {
	return &CreateEnvironmentConflict{}
}

/*CreateEnvironmentConflict handles this case with default header values.

An environment host with the provided name already exists.
*/
type CreateEnvironmentConflict struct {
	Payload *models.Error
}

func (o *CreateEnvironmentConflict) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/environments/][%d] createEnvironmentConflict  %+v", 409, o.Payload)
}

func (o *CreateEnvironmentConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateEnvironmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
