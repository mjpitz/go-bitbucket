// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// CreateOrUpdateAnnotationReader is a Reader for the CreateOrUpdateAnnotation structure.
type CreateOrUpdateAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrUpdateAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrUpdateAnnotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrUpdateAnnotationOK creates a CreateOrUpdateAnnotationOK with default headers values
func NewCreateOrUpdateAnnotationOK() *CreateOrUpdateAnnotationOK {
	return &CreateOrUpdateAnnotationOK{}
}

/*CreateOrUpdateAnnotationOK handles this case with default header values.

OK
*/
type CreateOrUpdateAnnotationOK struct {
	Payload *models.ReportAnnotation
}

func (o *CreateOrUpdateAnnotationOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId}][%d] createOrUpdateAnnotationOK  %+v", 200, o.Payload)
}

func (o *CreateOrUpdateAnnotationOK) GetPayload() *models.ReportAnnotation {
	return o.Payload
}

func (o *CreateOrUpdateAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportAnnotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrUpdateAnnotationBadRequest creates a CreateOrUpdateAnnotationBadRequest with default headers values
func NewCreateOrUpdateAnnotationBadRequest() *CreateOrUpdateAnnotationBadRequest {
	return &CreateOrUpdateAnnotationBadRequest{}
}

/*CreateOrUpdateAnnotationBadRequest handles this case with default header values.

The provided Annotation object is malformed or incomplete.
*/
type CreateOrUpdateAnnotationBadRequest struct {
	Payload *models.Error
}

func (o *CreateOrUpdateAnnotationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId}][%d] createOrUpdateAnnotationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrUpdateAnnotationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOrUpdateAnnotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
