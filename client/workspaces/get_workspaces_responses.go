// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetWorkspacesReader is a Reader for the GetWorkspaces structure.
type GetWorkspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkspacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkspacesOK creates a GetWorkspacesOK with default headers values
func NewGetWorkspacesOK() *GetWorkspacesOK {
	return &GetWorkspacesOK{}
}

/*GetWorkspacesOK handles this case with default header values.

The list of workspaces accessible by the user.
*/
type GetWorkspacesOK struct {
	Payload *models.PaginatedWorkspaces
}

func (o *GetWorkspacesOK) Error() string {
	return fmt.Sprintf("[GET /workspaces][%d] getWorkspacesOK  %+v", 200, o.Payload)
}

func (o *GetWorkspacesOK) GetPayload() *models.PaginatedWorkspaces {
	return o.Payload
}

func (o *GetWorkspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedWorkspaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacesUnauthorized creates a GetWorkspacesUnauthorized with default headers values
func NewGetWorkspacesUnauthorized() *GetWorkspacesUnauthorized {
	return &GetWorkspacesUnauthorized{}
}

/*GetWorkspacesUnauthorized handles this case with default header values.

The request wasn't authenticated.
*/
type GetWorkspacesUnauthorized struct {
	Payload *models.Error
}

func (o *GetWorkspacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workspaces][%d] getWorkspacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkspacesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWorkspacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}