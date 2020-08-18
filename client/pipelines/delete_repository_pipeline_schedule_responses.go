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

// DeleteRepositoryPipelineScheduleReader is a Reader for the DeleteRepositoryPipelineSchedule structure.
type DeleteRepositoryPipelineScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoryPipelineScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRepositoryPipelineScheduleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteRepositoryPipelineScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRepositoryPipelineScheduleNoContent creates a DeleteRepositoryPipelineScheduleNoContent with default headers values
func NewDeleteRepositoryPipelineScheduleNoContent() *DeleteRepositoryPipelineScheduleNoContent {
	return &DeleteRepositoryPipelineScheduleNoContent{}
}

/*DeleteRepositoryPipelineScheduleNoContent handles this case with default header values.

The schedule was deleted.
*/
type DeleteRepositoryPipelineScheduleNoContent struct {
}

func (o *DeleteRepositoryPipelineScheduleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid}][%d] deleteRepositoryPipelineScheduleNoContent ", 204)
}

func (o *DeleteRepositoryPipelineScheduleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoryPipelineScheduleNotFound creates a DeleteRepositoryPipelineScheduleNotFound with default headers values
func NewDeleteRepositoryPipelineScheduleNotFound() *DeleteRepositoryPipelineScheduleNotFound {
	return &DeleteRepositoryPipelineScheduleNotFound{}
}

/*DeleteRepositoryPipelineScheduleNotFound handles this case with default header values.

The account, repository or schedule was not found.
*/
type DeleteRepositoryPipelineScheduleNotFound struct {
	Payload *models.Error
}

func (o *DeleteRepositoryPipelineScheduleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid}][%d] deleteRepositoryPipelineScheduleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRepositoryPipelineScheduleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRepositoryPipelineScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}