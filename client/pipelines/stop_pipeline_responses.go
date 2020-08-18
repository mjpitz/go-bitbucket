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

// StopPipelineReader is a Reader for the StopPipeline structure.
type StopPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopPipelineNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopPipelineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopPipelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopPipelineNoContent creates a StopPipelineNoContent with default headers values
func NewStopPipelineNoContent() *StopPipelineNoContent {
	return &StopPipelineNoContent{}
}

/*StopPipelineNoContent handles this case with default header values.

The pipeline has been signaled to stop.
*/
type StopPipelineNoContent struct {
}

func (o *StopPipelineNoContent) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/stopPipeline][%d] stopPipelineNoContent ", 204)
}

func (o *StopPipelineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopPipelineBadRequest creates a StopPipelineBadRequest with default headers values
func NewStopPipelineBadRequest() *StopPipelineBadRequest {
	return &StopPipelineBadRequest{}
}

/*StopPipelineBadRequest handles this case with default header values.

The specified pipeline has already completed.
*/
type StopPipelineBadRequest struct {
	Payload *models.Error
}

func (o *StopPipelineBadRequest) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/stopPipeline][%d] stopPipelineBadRequest  %+v", 400, o.Payload)
}

func (o *StopPipelineBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopPipelineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopPipelineNotFound creates a StopPipelineNotFound with default headers values
func NewStopPipelineNotFound() *StopPipelineNotFound {
	return &StopPipelineNotFound{}
}

/*StopPipelineNotFound handles this case with default header values.

Either the account, repository or pipeline with the given UUID does not exist.
*/
type StopPipelineNotFound struct {
	Payload *models.Error
}

func (o *StopPipelineNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/stopPipeline][%d] stopPipelineNotFound  %+v", 404, o.Payload)
}

func (o *StopPipelineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopPipelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
