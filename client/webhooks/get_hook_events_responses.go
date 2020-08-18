// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetHookEventsReader is a Reader for the GetHookEvents structure.
type GetHookEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHookEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHookEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHookEventsOK creates a GetHookEventsOK with default headers values
func NewGetHookEventsOK() *GetHookEventsOK {
	return &GetHookEventsOK{}
}

/*GetHookEventsOK handles this case with default header values.

A mapping of resource/subject types pointing to their individual event types.
*/
type GetHookEventsOK struct {
	Payload *models.SubjectTypes
}

func (o *GetHookEventsOK) Error() string {
	return fmt.Sprintf("[GET /hook_events][%d] getHookEventsOK  %+v", 200, o.Payload)
}

func (o *GetHookEventsOK) GetPayload() *models.SubjectTypes {
	return o.Payload
}

func (o *GetHookEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubjectTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
