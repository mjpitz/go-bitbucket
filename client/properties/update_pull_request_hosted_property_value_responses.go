// Code generated by go-swagger; DO NOT EDIT.

package properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePullRequestHostedPropertyValueReader is a Reader for the UpdatePullRequestHostedPropertyValue structure.
type UpdatePullRequestHostedPropertyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePullRequestHostedPropertyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdatePullRequestHostedPropertyValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePullRequestHostedPropertyValueNoContent creates a UpdatePullRequestHostedPropertyValueNoContent with default headers values
func NewUpdatePullRequestHostedPropertyValueNoContent() *UpdatePullRequestHostedPropertyValueNoContent {
	return &UpdatePullRequestHostedPropertyValueNoContent{}
}

/*UpdatePullRequestHostedPropertyValueNoContent handles this case with default header values.

An empty response.
*/
type UpdatePullRequestHostedPropertyValueNoContent struct {
}

func (o *UpdatePullRequestHostedPropertyValueNoContent) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name}][%d] updatePullRequestHostedPropertyValueNoContent ", 204)
}

func (o *UpdatePullRequestHostedPropertyValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}