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

// DeleteDeploymentVariableReader is a Reader for the DeleteDeploymentVariable structure.
type DeleteDeploymentVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeploymentVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDeploymentVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDeploymentVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDeploymentVariableNoContent creates a DeleteDeploymentVariableNoContent with default headers values
func NewDeleteDeploymentVariableNoContent() *DeleteDeploymentVariableNoContent {
	return &DeleteDeploymentVariableNoContent{}
}

/*DeleteDeploymentVariableNoContent handles this case with default header values.

The variable was deleted.
*/
type DeleteDeploymentVariableNoContent struct {
}

func (o *DeleteDeploymentVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid}][%d] deleteDeploymentVariableNoContent ", 204)
}

func (o *DeleteDeploymentVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeploymentVariableNotFound creates a DeleteDeploymentVariableNotFound with default headers values
func NewDeleteDeploymentVariableNotFound() *DeleteDeploymentVariableNotFound {
	return &DeleteDeploymentVariableNotFound{}
}

/*DeleteDeploymentVariableNotFound handles this case with default header values.

The account, repository, environment or variable with given UUID was not found.
*/
type DeleteDeploymentVariableNotFound struct {
	Payload *models.Error
}

func (o *DeleteDeploymentVariableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid}][%d] deleteDeploymentVariableNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeploymentVariableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeploymentVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}