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

// CreatePipelineVariableForTeamReader is a Reader for the CreatePipelineVariableForTeam structure.
type CreatePipelineVariableForTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineVariableForTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePipelineVariableForTeamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreatePipelineVariableForTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePipelineVariableForTeamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePipelineVariableForTeamCreated creates a CreatePipelineVariableForTeamCreated with default headers values
func NewCreatePipelineVariableForTeamCreated() *CreatePipelineVariableForTeamCreated {
	return &CreatePipelineVariableForTeamCreated{}
}

/*CreatePipelineVariableForTeamCreated handles this case with default header values.

The created variable.
*/
type CreatePipelineVariableForTeamCreated struct {
	/*The URL of the newly created pipeline variable.
	 */
	Location string

	Payload *models.PipelineVariable
}

func (o *CreatePipelineVariableForTeamCreated) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/pipelines_config/variables/][%d] createPipelineVariableForTeamCreated  %+v", 201, o.Payload)
}

func (o *CreatePipelineVariableForTeamCreated) GetPayload() *models.PipelineVariable {
	return o.Payload
}

func (o *CreatePipelineVariableForTeamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.PipelineVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineVariableForTeamNotFound creates a CreatePipelineVariableForTeamNotFound with default headers values
func NewCreatePipelineVariableForTeamNotFound() *CreatePipelineVariableForTeamNotFound {
	return &CreatePipelineVariableForTeamNotFound{}
}

/*CreatePipelineVariableForTeamNotFound handles this case with default header values.

The account does not exist.
*/
type CreatePipelineVariableForTeamNotFound struct {
	Payload *models.Error
}

func (o *CreatePipelineVariableForTeamNotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/pipelines_config/variables/][%d] createPipelineVariableForTeamNotFound  %+v", 404, o.Payload)
}

func (o *CreatePipelineVariableForTeamNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePipelineVariableForTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineVariableForTeamConflict creates a CreatePipelineVariableForTeamConflict with default headers values
func NewCreatePipelineVariableForTeamConflict() *CreatePipelineVariableForTeamConflict {
	return &CreatePipelineVariableForTeamConflict{}
}

/*CreatePipelineVariableForTeamConflict handles this case with default header values.

A variable with the provided key already exists.
*/
type CreatePipelineVariableForTeamConflict struct {
	Payload *models.Error
}

func (o *CreatePipelineVariableForTeamConflict) Error() string {
	return fmt.Sprintf("[POST /teams/{username}/pipelines_config/variables/][%d] createPipelineVariableForTeamConflict  %+v", 409, o.Payload)
}

func (o *CreatePipelineVariableForTeamConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePipelineVariableForTeamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}