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

// GetRepositoryPipelineKnownHostReader is a Reader for the GetRepositoryPipelineKnownHost structure.
type GetRepositoryPipelineKnownHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryPipelineKnownHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryPipelineKnownHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRepositoryPipelineKnownHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoryPipelineKnownHostOK creates a GetRepositoryPipelineKnownHostOK with default headers values
func NewGetRepositoryPipelineKnownHostOK() *GetRepositoryPipelineKnownHostOK {
	return &GetRepositoryPipelineKnownHostOK{}
}

/*GetRepositoryPipelineKnownHostOK handles this case with default header values.

The known host.
*/
type GetRepositoryPipelineKnownHostOK struct {
	Payload *models.PipelineKnownHost
}

func (o *GetRepositoryPipelineKnownHostOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid}][%d] getRepositoryPipelineKnownHostOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryPipelineKnownHostOK) GetPayload() *models.PipelineKnownHost {
	return o.Payload
}

func (o *GetRepositoryPipelineKnownHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineKnownHost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryPipelineKnownHostNotFound creates a GetRepositoryPipelineKnownHostNotFound with default headers values
func NewGetRepositoryPipelineKnownHostNotFound() *GetRepositoryPipelineKnownHostNotFound {
	return &GetRepositoryPipelineKnownHostNotFound{}
}

/*GetRepositoryPipelineKnownHostNotFound handles this case with default header values.

The account, repository or known host with the specified UUID was not found.
*/
type GetRepositoryPipelineKnownHostNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoryPipelineKnownHostNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid}][%d] getRepositoryPipelineKnownHostNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryPipelineKnownHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoryPipelineKnownHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
