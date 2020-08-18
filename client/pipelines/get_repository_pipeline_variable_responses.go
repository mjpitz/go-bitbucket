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

// GetRepositoryPipelineVariableReader is a Reader for the GetRepositoryPipelineVariable structure.
type GetRepositoryPipelineVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryPipelineVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryPipelineVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRepositoryPipelineVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoryPipelineVariableOK creates a GetRepositoryPipelineVariableOK with default headers values
func NewGetRepositoryPipelineVariableOK() *GetRepositoryPipelineVariableOK {
	return &GetRepositoryPipelineVariableOK{}
}

/*GetRepositoryPipelineVariableOK handles this case with default header values.

The variable.
*/
type GetRepositoryPipelineVariableOK struct {
	Payload *models.PipelineVariable
}

func (o *GetRepositoryPipelineVariableOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid}][%d] getRepositoryPipelineVariableOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryPipelineVariableOK) GetPayload() *models.PipelineVariable {
	return o.Payload
}

func (o *GetRepositoryPipelineVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryPipelineVariableNotFound creates a GetRepositoryPipelineVariableNotFound with default headers values
func NewGetRepositoryPipelineVariableNotFound() *GetRepositoryPipelineVariableNotFound {
	return &GetRepositoryPipelineVariableNotFound{}
}

/*GetRepositoryPipelineVariableNotFound handles this case with default header values.

The account, repository or variable with the specified UUID was not found.
*/
type GetRepositoryPipelineVariableNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoryPipelineVariableNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid}][%d] getRepositoryPipelineVariableNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryPipelineVariableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoryPipelineVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}