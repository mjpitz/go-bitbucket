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

// GetDeploymentVariablesReader is a Reader for the GetDeploymentVariables structure.
type GetDeploymentVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentVariablesOK creates a GetDeploymentVariablesOK with default headers values
func NewGetDeploymentVariablesOK() *GetDeploymentVariablesOK {
	return &GetDeploymentVariablesOK{}
}

/*GetDeploymentVariablesOK handles this case with default header values.

The retrieved deployment variables.
*/
type GetDeploymentVariablesOK struct {
	Payload *models.PaginatedDeploymentVariable
}

func (o *GetDeploymentVariablesOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables][%d] getDeploymentVariablesOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentVariablesOK) GetPayload() *models.PaginatedDeploymentVariable {
	return o.Payload
}

func (o *GetDeploymentVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedDeploymentVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
