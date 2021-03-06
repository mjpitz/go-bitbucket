// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// CreateDeploymentVariableReader is a Reader for the CreateDeploymentVariable structure.
type CreateDeploymentVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeploymentVariableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateDeploymentVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDeploymentVariableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeploymentVariableCreated creates a CreateDeploymentVariableCreated with default headers values
func NewCreateDeploymentVariableCreated() *CreateDeploymentVariableCreated {
	return &CreateDeploymentVariableCreated{}
}

/*CreateDeploymentVariableCreated handles this case with default header values.

The variable was created.
*/
type CreateDeploymentVariableCreated struct {
	/*The URL of the newly created variable.
	 */
	Location string

	Payload *models.DeploymentVariable
}

func (o *CreateDeploymentVariableCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables][%d] createDeploymentVariableCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentVariableCreated) GetPayload() *models.DeploymentVariable {
	return o.Payload
}

func (o *CreateDeploymentVariableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.DeploymentVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentVariableNotFound creates a CreateDeploymentVariableNotFound with default headers values
func NewCreateDeploymentVariableNotFound() *CreateDeploymentVariableNotFound {
	return &CreateDeploymentVariableNotFound{}
}

/*CreateDeploymentVariableNotFound handles this case with default header values.

The account, repository, environment or variable with the given UUID was not found.
*/
type CreateDeploymentVariableNotFound struct {
	Payload *models.Error
}

func (o *CreateDeploymentVariableNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables][%d] createDeploymentVariableNotFound  %+v", 404, o.Payload)
}

func (o *CreateDeploymentVariableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeploymentVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentVariableConflict creates a CreateDeploymentVariableConflict with default headers values
func NewCreateDeploymentVariableConflict() *CreateDeploymentVariableConflict {
	return &CreateDeploymentVariableConflict{}
}

/*CreateDeploymentVariableConflict handles this case with default header values.

A variable with the provided key already exists.
*/
type CreateDeploymentVariableConflict struct {
	Payload *models.Error
}

func (o *CreateDeploymentVariableConflict) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables][%d] createDeploymentVariableConflict  %+v", 409, o.Payload)
}

func (o *CreateDeploymentVariableConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeploymentVariableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
