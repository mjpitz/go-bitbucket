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

// GetPipelineStepsForRepositoryReader is a Reader for the GetPipelineStepsForRepository structure.
type GetPipelineStepsForRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineStepsForRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineStepsForRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineStepsForRepositoryOK creates a GetPipelineStepsForRepositoryOK with default headers values
func NewGetPipelineStepsForRepositoryOK() *GetPipelineStepsForRepositoryOK {
	return &GetPipelineStepsForRepositoryOK{}
}

/*GetPipelineStepsForRepositoryOK handles this case with default header values.

The steps.
*/
type GetPipelineStepsForRepositoryOK struct {
	Payload *models.PaginatedPipelineSteps
}

func (o *GetPipelineStepsForRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/][%d] getPipelineStepsForRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetPipelineStepsForRepositoryOK) GetPayload() *models.PaginatedPipelineSteps {
	return o.Payload
}

func (o *GetPipelineStepsForRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPipelineSteps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
