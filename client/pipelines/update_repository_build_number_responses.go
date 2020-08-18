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

// UpdateRepositoryBuildNumberReader is a Reader for the UpdateRepositoryBuildNumber structure.
type UpdateRepositoryBuildNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryBuildNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryBuildNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepositoryBuildNumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepositoryBuildNumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryBuildNumberOK creates a UpdateRepositoryBuildNumberOK with default headers values
func NewUpdateRepositoryBuildNumberOK() *UpdateRepositoryBuildNumberOK {
	return &UpdateRepositoryBuildNumberOK{}
}

/*UpdateRepositoryBuildNumberOK handles this case with default header values.

The build number has been configured.
*/
type UpdateRepositoryBuildNumberOK struct {
	Payload *models.PipelineBuildNumber
}

func (o *UpdateRepositoryBuildNumberOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/build_number][%d] updateRepositoryBuildNumberOK  %+v", 200, o.Payload)
}

func (o *UpdateRepositoryBuildNumberOK) GetPayload() *models.PipelineBuildNumber {
	return o.Payload
}

func (o *UpdateRepositoryBuildNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineBuildNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryBuildNumberBadRequest creates a UpdateRepositoryBuildNumberBadRequest with default headers values
func NewUpdateRepositoryBuildNumberBadRequest() *UpdateRepositoryBuildNumberBadRequest {
	return &UpdateRepositoryBuildNumberBadRequest{}
}

/*UpdateRepositoryBuildNumberBadRequest handles this case with default header values.

The update failed because the next number was invalid (it should be higher than the current number).
*/
type UpdateRepositoryBuildNumberBadRequest struct {
	Payload *models.Error
}

func (o *UpdateRepositoryBuildNumberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/build_number][%d] updateRepositoryBuildNumberBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRepositoryBuildNumberBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryBuildNumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryBuildNumberNotFound creates a UpdateRepositoryBuildNumberNotFound with default headers values
func NewUpdateRepositoryBuildNumberNotFound() *UpdateRepositoryBuildNumberNotFound {
	return &UpdateRepositoryBuildNumberNotFound{}
}

/*UpdateRepositoryBuildNumberNotFound handles this case with default header values.

The account or repository was not found.
*/
type UpdateRepositoryBuildNumberNotFound struct {
	Payload *models.Error
}

func (o *UpdateRepositoryBuildNumberNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config/build_number][%d] updateRepositoryBuildNumberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRepositoryBuildNumberNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryBuildNumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
