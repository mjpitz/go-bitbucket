// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAddonUsersTargetUserEventsEventKeyReader is a Reader for the PostAddonUsersTargetUserEventsEventKey structure.
type PostAddonUsersTargetUserEventsEventKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddonUsersTargetUserEventsEventKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAddonUsersTargetUserEventsEventKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostAddonUsersTargetUserEventsEventKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAddonUsersTargetUserEventsEventKeyNoContent creates a PostAddonUsersTargetUserEventsEventKeyNoContent with default headers values
func NewPostAddonUsersTargetUserEventsEventKeyNoContent() *PostAddonUsersTargetUserEventsEventKeyNoContent {
	return &PostAddonUsersTargetUserEventsEventKeyNoContent{}
}

/*PostAddonUsersTargetUserEventsEventKeyNoContent handles this case with default header values.

Event successfully submitted
*/
type PostAddonUsersTargetUserEventsEventKeyNoContent struct {
}

func (o *PostAddonUsersTargetUserEventsEventKeyNoContent) Error() string {
	return fmt.Sprintf("[POST /addon/users/{target_user}/events/{event_key}][%d] postAddonUsersTargetUserEventsEventKeyNoContent ", 204)
}

func (o *PostAddonUsersTargetUserEventsEventKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAddonUsersTargetUserEventsEventKeyNotFound creates a PostAddonUsersTargetUserEventsEventKeyNotFound with default headers values
func NewPostAddonUsersTargetUserEventsEventKeyNotFound() *PostAddonUsersTargetUserEventsEventKeyNotFound {
	return &PostAddonUsersTargetUserEventsEventKeyNotFound{}
}

/*PostAddonUsersTargetUserEventsEventKeyNotFound handles this case with default header values.

Connect app not installed or event does not exist
*/
type PostAddonUsersTargetUserEventsEventKeyNotFound struct {
}

func (o *PostAddonUsersTargetUserEventsEventKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /addon/users/{target_user}/events/{event_key}][%d] postAddonUsersTargetUserEventsEventKeyNotFound ", 404)
}

func (o *PostAddonUsersTargetUserEventsEventKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
