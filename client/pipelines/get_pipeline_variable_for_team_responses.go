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

// GetPipelineVariableForTeamReader is a Reader for the GetPipelineVariableForTeam structure.
type GetPipelineVariableForTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineVariableForTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineVariableForTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPipelineVariableForTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineVariableForTeamOK creates a GetPipelineVariableForTeamOK with default headers values
func NewGetPipelineVariableForTeamOK() *GetPipelineVariableForTeamOK {
	return &GetPipelineVariableForTeamOK{}
}

/*GetPipelineVariableForTeamOK handles this case with default header values.

The variable.
*/
type GetPipelineVariableForTeamOK struct {
	Payload *models.PipelineVariable
}

func (o *GetPipelineVariableForTeamOK) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/pipelines_config/variables/{variable_uuid}][%d] getPipelineVariableForTeamOK  %+v", 200, o.Payload)
}

func (o *GetPipelineVariableForTeamOK) GetPayload() *models.PipelineVariable {
	return o.Payload
}

func (o *GetPipelineVariableForTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineVariableForTeamNotFound creates a GetPipelineVariableForTeamNotFound with default headers values
func NewGetPipelineVariableForTeamNotFound() *GetPipelineVariableForTeamNotFound {
	return &GetPipelineVariableForTeamNotFound{}
}

/*GetPipelineVariableForTeamNotFound handles this case with default header values.

The account or variable with the given UUID was not found.
*/
type GetPipelineVariableForTeamNotFound struct {
	Payload *models.Error
}

func (o *GetPipelineVariableForTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{username}/pipelines_config/variables/{variable_uuid}][%d] getPipelineVariableForTeamNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineVariableForTeamNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineVariableForTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
