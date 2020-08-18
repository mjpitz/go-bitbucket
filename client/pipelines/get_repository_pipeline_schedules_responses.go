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

// GetRepositoryPipelineSchedulesReader is a Reader for the GetRepositoryPipelineSchedules structure.
type GetRepositoryPipelineSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryPipelineSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryPipelineSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRepositoryPipelineSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoryPipelineSchedulesOK creates a GetRepositoryPipelineSchedulesOK with default headers values
func NewGetRepositoryPipelineSchedulesOK() *GetRepositoryPipelineSchedulesOK {
	return &GetRepositoryPipelineSchedulesOK{}
}

/*GetRepositoryPipelineSchedulesOK handles this case with default header values.

The list of schedules.
*/
type GetRepositoryPipelineSchedulesOK struct {
	Payload *models.PaginatedPipelineSchedules
}

func (o *GetRepositoryPipelineSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/][%d] getRepositoryPipelineSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryPipelineSchedulesOK) GetPayload() *models.PaginatedPipelineSchedules {
	return o.Payload
}

func (o *GetRepositoryPipelineSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPipelineSchedules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoryPipelineSchedulesNotFound creates a GetRepositoryPipelineSchedulesNotFound with default headers values
func NewGetRepositoryPipelineSchedulesNotFound() *GetRepositoryPipelineSchedulesNotFound {
	return &GetRepositoryPipelineSchedulesNotFound{}
}

/*GetRepositoryPipelineSchedulesNotFound handles this case with default header values.

The account or repository was not found.
*/
type GetRepositoryPipelineSchedulesNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoryPipelineSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/][%d] getRepositoryPipelineSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoryPipelineSchedulesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoryPipelineSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
