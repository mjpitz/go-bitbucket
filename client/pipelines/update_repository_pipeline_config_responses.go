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

// UpdateRepositoryPipelineConfigReader is a Reader for the UpdateRepositoryPipelineConfig structure.
type UpdateRepositoryPipelineConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryPipelineConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryPipelineConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryPipelineConfigOK creates a UpdateRepositoryPipelineConfigOK with default headers values
func NewUpdateRepositoryPipelineConfigOK() *UpdateRepositoryPipelineConfigOK {
	return &UpdateRepositoryPipelineConfigOK{}
}

/*UpdateRepositoryPipelineConfigOK handles this case with default header values.

The repository pipelines configuration was updated.
*/
type UpdateRepositoryPipelineConfigOK struct {
	Payload *models.PipelinesConfig
}

func (o *UpdateRepositoryPipelineConfigOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pipelines_config][%d] updateRepositoryPipelineConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateRepositoryPipelineConfigOK) GetPayload() *models.PipelinesConfig {
	return o.Payload
}

func (o *UpdateRepositoryPipelineConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelinesConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
