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

// PostSnippetsReader is a Reader for the PostSnippets structure.
type PostSnippetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSnippetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostSnippetsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostSnippetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSnippetsCreated creates a PostSnippetsCreated with default headers values
func NewPostSnippetsCreated() *PostSnippetsCreated {
	return &PostSnippetsCreated{}
}

/*PostSnippetsCreated handles this case with default header values.

The newly created snippet object.
*/
type PostSnippetsCreated struct {
	/*The URL of the newly created snippet.
	 */
	Location string

	Payload *models.Snippet
}

func (o *PostSnippetsCreated) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] postSnippetsCreated  %+v", 201, o.Payload)
}

func (o *PostSnippetsCreated) GetPayload() *models.Snippet {
	return o.Payload
}

func (o *PostSnippetsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.Snippet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSnippetsUnauthorized creates a PostSnippetsUnauthorized with default headers values
func NewPostSnippetsUnauthorized() *PostSnippetsUnauthorized {
	return &PostSnippetsUnauthorized{}
}

/*PostSnippetsUnauthorized handles this case with default header values.

If the request was not authenticated
*/
type PostSnippetsUnauthorized struct {
	Payload *models.Error
}

func (o *PostSnippetsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /snippets][%d] postSnippetsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSnippetsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSnippetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
