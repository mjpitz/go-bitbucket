// Code generated by go-swagger; DO NOT EDIT.

package snippets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetSnippetsReader is a Reader for the GetSnippets structure.
type GetSnippetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnippetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSnippetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsOK creates a GetSnippetsOK with default headers values
func NewGetSnippetsOK() *GetSnippetsOK {
	return &GetSnippetsOK{}
}

/*GetSnippetsOK handles this case with default header values.

A paginated list of snippets.
*/
type GetSnippetsOK struct {
	Payload *models.PaginatedSnippets
}

func (o *GetSnippetsOK) Error() string {
	return fmt.Sprintf("[GET /snippets][%d] getSnippetsOK  %+v", 200, o.Payload)
}

func (o *GetSnippetsOK) GetPayload() *models.PaginatedSnippets {
	return o.Payload
}

func (o *GetSnippetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedSnippets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsNotFound creates a GetSnippetsNotFound with default headers values
func NewGetSnippetsNotFound() *GetSnippetsNotFound {
	return &GetSnippetsNotFound{}
}

/*GetSnippetsNotFound handles this case with default header values.

If the snippet does not exist.
*/
type GetSnippetsNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets][%d] getSnippetsNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
