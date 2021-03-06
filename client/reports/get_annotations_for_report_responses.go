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

// GetAnnotationsForReportReader is a Reader for the GetAnnotationsForReport structure.
type GetAnnotationsForReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnnotationsForReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnnotationsForReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnnotationsForReportOK creates a GetAnnotationsForReportOK with default headers values
func NewGetAnnotationsForReportOK() *GetAnnotationsForReportOK {
	return &GetAnnotationsForReportOK{}
}

/*GetAnnotationsForReportOK handles this case with default header values.

OK
*/
type GetAnnotationsForReportOK struct {
	Payload *models.PaginatedAnnotations
}

func (o *GetAnnotationsForReportOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations][%d] getAnnotationsForReportOK  %+v", 200, o.Payload)
}

func (o *GetAnnotationsForReportOK) GetPayload() *models.PaginatedAnnotations {
	return o.Payload
}

func (o *GetAnnotationsForReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedAnnotations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
