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

// CreatePipelineForRepositoryReader is a Reader for the CreatePipelineForRepository structure.
type CreatePipelineForRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineForRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePipelineForRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePipelineForRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePipelineForRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePipelineForRepositoryCreated creates a CreatePipelineForRepositoryCreated with default headers values
func NewCreatePipelineForRepositoryCreated() *CreatePipelineForRepositoryCreated {
	return &CreatePipelineForRepositoryCreated{}
}

/*CreatePipelineForRepositoryCreated handles this case with default header values.

The initiated pipeline.
*/
type CreatePipelineForRepositoryCreated struct {
	/*The URL of the newly created pipeline.
	 */
	Location string

	Payload *models.Pipeline
}

func (o *CreatePipelineForRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/][%d] createPipelineForRepositoryCreated  %+v", 201, o.Payload)
}

func (o *CreatePipelineForRepositoryCreated) GetPayload() *models.Pipeline {
	return o.Payload
}

func (o *CreatePipelineForRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineForRepositoryBadRequest creates a CreatePipelineForRepositoryBadRequest with default headers values
func NewCreatePipelineForRepositoryBadRequest() *CreatePipelineForRepositoryBadRequest {
	return &CreatePipelineForRepositoryBadRequest{}
}

/*CreatePipelineForRepositoryBadRequest handles this case with default header values.

The account or repository is not enabled, the yml file does not exist in the repository for the given revision, or the request body contained invalid properties.
*/
type CreatePipelineForRepositoryBadRequest struct {
	Payload *models.Error
}

func (o *CreatePipelineForRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/][%d] createPipelineForRepositoryBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePipelineForRepositoryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePipelineForRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineForRepositoryNotFound creates a CreatePipelineForRepositoryNotFound with default headers values
func NewCreatePipelineForRepositoryNotFound() *CreatePipelineForRepositoryNotFound {
	return &CreatePipelineForRepositoryNotFound{}
}

/*CreatePipelineForRepositoryNotFound handles this case with default header values.

The account or repository was not found.
*/
type CreatePipelineForRepositoryNotFound struct {
	Payload *models.Error
}

func (o *CreatePipelineForRepositoryNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pipelines/][%d] createPipelineForRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *CreatePipelineForRepositoryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePipelineForRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}