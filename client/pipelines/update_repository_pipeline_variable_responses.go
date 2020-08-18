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

// UpdateRepositoryPipelineVariableReader is a Reader for the UpdateRepositoryPipelineVariable structure.
type UpdateRepositoryPipelineVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryPipelineVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryPipelineVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateRepositoryPipelineVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryPipelineVariableOK creates a UpdateRepositoryPipelineVariableOK with default headers values
func NewUpdateRepositoryPipelineVariableOK() *UpdateRepositoryPipelineVariableOK {
	return &UpdateRepositoryPipelineVariableOK{}
}

/*UpdateRepositoryPipelineVariableOK handles this case with default header values.

The variable was updated.
*/
type UpdateRepositoryPipelineVariableOK struct {
	Payload *models.PipelineVariable
}

func (o *UpdateRepositoryPipelineVariableOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid}][%d] updateRepositoryPipelineVariableOK  %+v", 200, o.Payload)
}

func (o *UpdateRepositoryPipelineVariableOK) GetPayload() *models.PipelineVariable {
	return o.Payload
}

func (o *UpdateRepositoryPipelineVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryPipelineVariableNotFound creates a UpdateRepositoryPipelineVariableNotFound with default headers values
func NewUpdateRepositoryPipelineVariableNotFound() *UpdateRepositoryPipelineVariableNotFound {
	return &UpdateRepositoryPipelineVariableNotFound{}
}

/*UpdateRepositoryPipelineVariableNotFound handles this case with default header values.

The account, repository or variable with the given UUID was not found.
*/
type UpdateRepositoryPipelineVariableNotFound struct {
	Payload *models.Error
}

func (o *UpdateRepositoryPipelineVariableNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid}][%d] updateRepositoryPipelineVariableNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRepositoryPipelineVariableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryPipelineVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
