// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesReader is a Reader for the GetRepositories structure.
type GetRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesOK creates a GetRepositoriesOK with default headers values
func NewGetRepositoriesOK() *GetRepositoriesOK {
	return &GetRepositoriesOK{}
}

/*GetRepositoriesOK handles this case with default header values.

All public repositories.
*/
type GetRepositoriesOK struct {
	Payload *models.PaginatedRepositories
}

func (o *GetRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /repositories][%d] getRepositoriesOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesOK) GetPayload() *models.PaginatedRepositories {
	return o.Payload
}

func (o *GetRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedRepositories)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
