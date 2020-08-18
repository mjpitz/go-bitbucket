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

// GetPipelinesForRepositoryReader is a Reader for the GetPipelinesForRepository structure.
type GetPipelinesForRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelinesForRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelinesForRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelinesForRepositoryOK creates a GetPipelinesForRepositoryOK with default headers values
func NewGetPipelinesForRepositoryOK() *GetPipelinesForRepositoryOK {
	return &GetPipelinesForRepositoryOK{}
}

/*GetPipelinesForRepositoryOK handles this case with default header values.

The matching pipelines.
*/
type GetPipelinesForRepositoryOK struct {
	Payload *models.PaginatedPipelines
}

func (o *GetPipelinesForRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pipelines/][%d] getPipelinesForRepositoryOK  %+v", 200, o.Payload)
}

func (o *GetPipelinesForRepositoryOK) GetPayload() *models.PaginatedPipelines {
	return o.Payload
}

func (o *GetPipelinesForRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPipelines)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
