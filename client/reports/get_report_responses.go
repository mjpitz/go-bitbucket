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

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/*GetReportOK handles this case with default header values.

OK
*/
type GetReportOK struct {
	Payload *models.Report
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}][%d] getReportOK  %+v", 200, o.Payload)
}

func (o *GetReportOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportNotFound creates a GetReportNotFound with default headers values
func NewGetReportNotFound() *GetReportNotFound {
	return &GetReportNotFound{}
}

/*GetReportNotFound handles this case with default header values.

The report with the given ID was not found.
*/
type GetReportNotFound struct {
	Payload *models.Error
}

func (o *GetReportNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}][%d] getReportNotFound  %+v", 404, o.Payload)
}

func (o *GetReportNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}