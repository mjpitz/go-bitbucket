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

// PutSnippetsWorkspaceEncodedIDNodeIDReader is a Reader for the PutSnippetsWorkspaceEncodedIDNodeID structure.
type PutSnippetsWorkspaceEncodedIDNodeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSnippetsWorkspaceEncodedIDNodeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSnippetsWorkspaceEncodedIDNodeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutSnippetsWorkspaceEncodedIDNodeIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSnippetsWorkspaceEncodedIDNodeIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutSnippetsWorkspaceEncodedIDNodeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSnippetsWorkspaceEncodedIDNodeIDOK creates a PutSnippetsWorkspaceEncodedIDNodeIDOK with default headers values
func NewPutSnippetsWorkspaceEncodedIDNodeIDOK() *PutSnippetsWorkspaceEncodedIDNodeIDOK {
	return &PutSnippetsWorkspaceEncodedIDNodeIDOK{}
}

/*PutSnippetsWorkspaceEncodedIDNodeIDOK handles this case with default header values.

The updated snippet object.
*/
type PutSnippetsWorkspaceEncodedIDNodeIDOK struct {
	Payload *models.Snippet
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDOK) Error() string {
	return fmt.Sprintf("[PUT /snippets/{workspace}/{encoded_id}/{node_id}][%d] putSnippetsWorkspaceEncodedIdNodeIdOK  %+v", 200, o.Payload)
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDOK) GetPayload() *models.Snippet {
	return o.Payload
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snippet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSnippetsWorkspaceEncodedIDNodeIDUnauthorized creates a PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized with default headers values
func NewPutSnippetsWorkspaceEncodedIDNodeIDUnauthorized() *PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized {
	return &PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized{}
}

/*PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized handles this case with default header values.

If the snippet is private and the request was not authenticated.
*/
type PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /snippets/{workspace}/{encoded_id}/{node_id}][%d] putSnippetsWorkspaceEncodedIdNodeIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSnippetsWorkspaceEncodedIDNodeIDForbidden creates a PutSnippetsWorkspaceEncodedIDNodeIDForbidden with default headers values
func NewPutSnippetsWorkspaceEncodedIDNodeIDForbidden() *PutSnippetsWorkspaceEncodedIDNodeIDForbidden {
	return &PutSnippetsWorkspaceEncodedIDNodeIDForbidden{}
}

/*PutSnippetsWorkspaceEncodedIDNodeIDForbidden handles this case with default header values.

If authenticated user does not have permission to update the private snippet.
*/
type PutSnippetsWorkspaceEncodedIDNodeIDForbidden struct {
	Payload *models.Error
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /snippets/{workspace}/{encoded_id}/{node_id}][%d] putSnippetsWorkspaceEncodedIdNodeIdForbidden  %+v", 403, o.Payload)
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSnippetsWorkspaceEncodedIDNodeIDNotFound creates a PutSnippetsWorkspaceEncodedIDNodeIDNotFound with default headers values
func NewPutSnippetsWorkspaceEncodedIDNodeIDNotFound() *PutSnippetsWorkspaceEncodedIDNodeIDNotFound {
	return &PutSnippetsWorkspaceEncodedIDNodeIDNotFound{}
}

/*PutSnippetsWorkspaceEncodedIDNodeIDNotFound handles this case with default header values.

If the snippet or the revision does not exist.
*/
type PutSnippetsWorkspaceEncodedIDNodeIDNotFound struct {
	Payload *models.Error
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /snippets/{workspace}/{encoded_id}/{node_id}][%d] putSnippetsWorkspaceEncodedIdNodeIdNotFound  %+v", 404, o.Payload)
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed creates a PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed with default headers values
func NewPutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed() *PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed {
	return &PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed{}
}

/*PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed handles this case with default header values.

If `{node_id}` is not the latest revision.
*/
type PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed struct {
	Payload *models.Error
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /snippets/{workspace}/{encoded_id}/{node_id}][%d] putSnippetsWorkspaceEncodedIdNodeIdMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutSnippetsWorkspaceEncodedIDNodeIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
