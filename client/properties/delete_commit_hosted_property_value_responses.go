// Code generated by go-swagger; DO NOT EDIT.

package properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCommitHostedPropertyValueReader is a Reader for the DeleteCommitHostedPropertyValue structure.
type DeleteCommitHostedPropertyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCommitHostedPropertyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCommitHostedPropertyValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCommitHostedPropertyValueNoContent creates a DeleteCommitHostedPropertyValueNoContent with default headers values
func NewDeleteCommitHostedPropertyValueNoContent() *DeleteCommitHostedPropertyValueNoContent {
	return &DeleteCommitHostedPropertyValueNoContent{}
}

/*DeleteCommitHostedPropertyValueNoContent handles this case with default header values.

An empty response.
*/
type DeleteCommitHostedPropertyValueNoContent struct {
}

func (o *DeleteCommitHostedPropertyValueNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name}][%d] deleteCommitHostedPropertyValueNoContent ", 204)
}

func (o *DeleteCommitHostedPropertyValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
