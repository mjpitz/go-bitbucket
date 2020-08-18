// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetUserEmailsEmailReader is a Reader for the GetUserEmailsEmail structure.
type GetUserEmailsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserEmailsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetUserEmailsEmailDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetUserEmailsEmailDefault creates a GetUserEmailsEmailDefault with default headers values
func NewGetUserEmailsEmailDefault(code int) *GetUserEmailsEmailDefault {
	return &GetUserEmailsEmailDefault{
		_statusCode: code,
	}
}

/*GetUserEmailsEmailDefault handles this case with default header values.

Unexpected error.
*/
type GetUserEmailsEmailDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get user emails email default response
func (o *GetUserEmailsEmailDefault) Code() int {
	return o._statusCode
}

func (o *GetUserEmailsEmailDefault) Error() string {
	return fmt.Sprintf("[GET /user/emails/{email}][%d] GetUserEmailsEmail default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserEmailsEmailDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserEmailsEmailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}