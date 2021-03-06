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

// UpdateRepositoryPipelineKeyPairReader is a Reader for the UpdateRepositoryPipelineKeyPair structure.
type UpdateRepositoryPipelineKeyPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryPipelineKeyPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryPipelineKeyPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateRepositoryPipelineKeyPairNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryPipelineKeyPairOK creates a UpdateRepositoryPipelineKeyPairOK with default headers values
func NewUpdateRepositoryPipelineKeyPairOK() *UpdateRepositoryPipelineKeyPairOK {
	return &UpdateRepositoryPipelineKeyPairOK{}
}

/*UpdateRepositoryPipelineKeyPairOK handles this case with default header values.

The SSH key pair was created or updated.
*/
type UpdateRepositoryPipelineKeyPairOK struct {
	Payload *models.PipelineSSHKeyPair
}

func (o *UpdateRepositoryPipelineKeyPairOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair][%d] updateRepositoryPipelineKeyPairOK  %+v", 200, o.Payload)
}

func (o *UpdateRepositoryPipelineKeyPairOK) GetPayload() *models.PipelineSSHKeyPair {
	return o.Payload
}

func (o *UpdateRepositoryPipelineKeyPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineSSHKeyPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryPipelineKeyPairNotFound creates a UpdateRepositoryPipelineKeyPairNotFound with default headers values
func NewUpdateRepositoryPipelineKeyPairNotFound() *UpdateRepositoryPipelineKeyPairNotFound {
	return &UpdateRepositoryPipelineKeyPairNotFound{}
}

/*UpdateRepositoryPipelineKeyPairNotFound handles this case with default header values.

The account, repository or SSH key pair was not found.
*/
type UpdateRepositoryPipelineKeyPairNotFound struct {
	Payload *models.Error
}

func (o *UpdateRepositoryPipelineKeyPairNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair][%d] updateRepositoryPipelineKeyPairNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRepositoryPipelineKeyPairNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryPipelineKeyPairNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
