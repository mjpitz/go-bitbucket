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

// GetPipelineStepLogForRepositoryReader is a Reader for the GetPipelineStepLogForRepository structure.
type GetPipelineStepLogForRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineStepLogForRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineStepLogForRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetPipelineStepLogForRepositoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineStepLogForRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 416:
		result := NewGetPipelineStepLogForRepositoryRequestRangeNotSatisfiable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineStepLogForRepositoryOK creates a GetPipelineStepLogForRepositoryOK with default headers values
func NewGetPipelineStepLogForRepositoryOK() *GetPipelineStepLogForRepositoryOK {
	return &GetPipelineStepLogForRepositoryOK{}
}

/*GetPipelineStepLogForRepositoryOK handles this case with default header values.

The raw log file for this pipeline step.
*/
type GetPipelineStepLogForRepositoryOK struct {
}

func (o *GetPipelineStepLogForRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log][%d] getPipelineStepLogForRepositoryOK ", 200)
}

func (o *GetPipelineStepLogForRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineStepLogForRepositoryNotModified creates a GetPipelineStepLogForRepositoryNotModified with default headers values
func NewGetPipelineStepLogForRepositoryNotModified() *GetPipelineStepLogForRepositoryNotModified {
	return &GetPipelineStepLogForRepositoryNotModified{}
}

/*GetPipelineStepLogForRepositoryNotModified handles this case with default header values.

The log has the same etag as the provided If-None-Match header.
*/
type GetPipelineStepLogForRepositoryNotModified struct {
	Payload *models.Error
}

func (o *GetPipelineStepLogForRepositoryNotModified) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log][%d] getPipelineStepLogForRepositoryNotModified  %+v", 304, o.Payload)
}

func (o *GetPipelineStepLogForRepositoryNotModified) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineStepLogForRepositoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineStepLogForRepositoryNotFound creates a GetPipelineStepLogForRepositoryNotFound with default headers values
func NewGetPipelineStepLogForRepositoryNotFound() *GetPipelineStepLogForRepositoryNotFound {
	return &GetPipelineStepLogForRepositoryNotFound{}
}

/*GetPipelineStepLogForRepositoryNotFound handles this case with default header values.

A pipeline with the given UUID does not exist, a step with the given UUID does not exist in the pipeline or a log file does not exist for the given step.
*/
type GetPipelineStepLogForRepositoryNotFound struct {
	Payload *models.Error
}

func (o *GetPipelineStepLogForRepositoryNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log][%d] getPipelineStepLogForRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineStepLogForRepositoryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineStepLogForRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineStepLogForRepositoryRequestRangeNotSatisfiable creates a GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable with default headers values
func NewGetPipelineStepLogForRepositoryRequestRangeNotSatisfiable() *GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable {
	return &GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable{}
}

/*GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable handles this case with default header values.

The requested range does not exist for requests that specified the [HTTP Range header](https://tools.ietf.org/html/rfc7233#section-3.1).
*/
type GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable struct {
	Payload *models.Error
}

func (o *GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log][%d] getPipelineStepLogForRepositoryRequestRangeNotSatisfiable  %+v", 416, o.Payload)
}

func (o *GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineStepLogForRepositoryRequestRangeNotSatisfiable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
