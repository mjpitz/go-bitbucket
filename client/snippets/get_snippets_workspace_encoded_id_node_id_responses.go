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

// GetSnippetsWorkspaceEncodedIDNodeIDReader is a Reader for the GetSnippetsWorkspaceEncodedIDNodeID structure.
type GetSnippetsWorkspaceEncodedIDNodeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsWorkspaceEncodedIDNodeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnippetsWorkspaceEncodedIDNodeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSnippetsWorkspaceEncodedIDNodeIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSnippetsWorkspaceEncodedIDNodeIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnippetsWorkspaceEncodedIDNodeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsWorkspaceEncodedIDNodeIDOK creates a GetSnippetsWorkspaceEncodedIDNodeIDOK with default headers values
func NewGetSnippetsWorkspaceEncodedIDNodeIDOK() *GetSnippetsWorkspaceEncodedIDNodeIDOK {
	return &GetSnippetsWorkspaceEncodedIDNodeIDOK{}
}

/*GetSnippetsWorkspaceEncodedIDNodeIDOK handles this case with default header values.

The snippet object.
*/
type GetSnippetsWorkspaceEncodedIDNodeIDOK struct {
	Payload *models.Snippet
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDOK) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{node_id}][%d] getSnippetsWorkspaceEncodedIdNodeIdOK  %+v", 200, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDOK) GetPayload() *models.Snippet {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snippet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDNodeIDUnauthorized creates a GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized with default headers values
func NewGetSnippetsWorkspaceEncodedIDNodeIDUnauthorized() *GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized {
	return &GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized{}
}

/*GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized handles this case with default header values.

If the snippet is private and the request was not authenticated.
*/
type GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{node_id}][%d] getSnippetsWorkspaceEncodedIdNodeIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDNodeIDForbidden creates a GetSnippetsWorkspaceEncodedIDNodeIDForbidden with default headers values
func NewGetSnippetsWorkspaceEncodedIDNodeIDForbidden() *GetSnippetsWorkspaceEncodedIDNodeIDForbidden {
	return &GetSnippetsWorkspaceEncodedIDNodeIDForbidden{}
}

/*GetSnippetsWorkspaceEncodedIDNodeIDForbidden handles this case with default header values.

If authenticated user does not have access to the private snippet.
*/
type GetSnippetsWorkspaceEncodedIDNodeIDForbidden struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDForbidden) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{node_id}][%d] getSnippetsWorkspaceEncodedIdNodeIdForbidden  %+v", 403, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDNodeIDNotFound creates a GetSnippetsWorkspaceEncodedIDNodeIDNotFound with default headers values
func NewGetSnippetsWorkspaceEncodedIDNodeIDNotFound() *GetSnippetsWorkspaceEncodedIDNodeIDNotFound {
	return &GetSnippetsWorkspaceEncodedIDNodeIDNotFound{}
}

/*GetSnippetsWorkspaceEncodedIDNodeIDNotFound handles this case with default header values.

If the snippet, or the revision does not exist.
*/
type GetSnippetsWorkspaceEncodedIDNodeIDNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{node_id}][%d] getSnippetsWorkspaceEncodedIdNodeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDNodeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}