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

// UpdateDeploymentVariableReader is a Reader for the UpdateDeploymentVariable structure.
type UpdateDeploymentVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeploymentVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeploymentVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateDeploymentVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeploymentVariableOK creates a UpdateDeploymentVariableOK with default headers values
func NewUpdateDeploymentVariableOK() *UpdateDeploymentVariableOK {
	return &UpdateDeploymentVariableOK{}
}

/*UpdateDeploymentVariableOK handles this case with default header values.

The deployment variable was updated.
*/
type UpdateDeploymentVariableOK struct {
	Payload *models.DeploymentVariable
}

func (o *UpdateDeploymentVariableOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid}][%d] updateDeploymentVariableOK  %+v", 200, o.Payload)
}

func (o *UpdateDeploymentVariableOK) GetPayload() *models.DeploymentVariable {
	return o.Payload
}

func (o *UpdateDeploymentVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeploymentVariableNotFound creates a UpdateDeploymentVariableNotFound with default headers values
func NewUpdateDeploymentVariableNotFound() *UpdateDeploymentVariableNotFound {
	return &UpdateDeploymentVariableNotFound{}
}

/*UpdateDeploymentVariableNotFound handles this case with default header values.

The account, repository, environment or variable with the given UUID was not found.
*/
type UpdateDeploymentVariableNotFound struct {
	Payload *models.Error
}

func (o *UpdateDeploymentVariableNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid}][%d] updateDeploymentVariableNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDeploymentVariableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDeploymentVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}