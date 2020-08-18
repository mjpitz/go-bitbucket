// Code generated by go-swagger; DO NOT EDIT.

package properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateCommitHostedPropertyValueReader is a Reader for the UpdateCommitHostedPropertyValue structure.
type UpdateCommitHostedPropertyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCommitHostedPropertyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateCommitHostedPropertyValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCommitHostedPropertyValueNoContent creates a UpdateCommitHostedPropertyValueNoContent with default headers values
func NewUpdateCommitHostedPropertyValueNoContent() *UpdateCommitHostedPropertyValueNoContent {
	return &UpdateCommitHostedPropertyValueNoContent{}
}

/*UpdateCommitHostedPropertyValueNoContent handles this case with default header values.

An empty response.
*/
type UpdateCommitHostedPropertyValueNoContent struct {
}

func (o *UpdateCommitHostedPropertyValueNoContent) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name}][%d] updateCommitHostedPropertyValueNoContent ", 204)
}

func (o *UpdateCommitHostedPropertyValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}