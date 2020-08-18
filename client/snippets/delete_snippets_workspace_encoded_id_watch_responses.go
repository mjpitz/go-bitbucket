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

// DeleteSnippetsWorkspaceEncodedIDWatchReader is a Reader for the DeleteSnippetsWorkspaceEncodedIDWatch structure.
type DeleteSnippetsWorkspaceEncodedIDWatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnippetsWorkspaceEncodedIDWatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSnippetsWorkspaceEncodedIDWatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSnippetsWorkspaceEncodedIDWatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSnippetsWorkspaceEncodedIDWatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSnippetsWorkspaceEncodedIDWatchNoContent creates a DeleteSnippetsWorkspaceEncodedIDWatchNoContent with default headers values
func NewDeleteSnippetsWorkspaceEncodedIDWatchNoContent() *DeleteSnippetsWorkspaceEncodedIDWatchNoContent {
	return &DeleteSnippetsWorkspaceEncodedIDWatchNoContent{}
}

/*DeleteSnippetsWorkspaceEncodedIDWatchNoContent handles this case with default header values.

Indicates the user stopped watching the snippet successfully.
*/
type DeleteSnippetsWorkspaceEncodedIDWatchNoContent struct {
	Payload *models.PaginatedUsers
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{workspace}/{encoded_id}/watch][%d] deleteSnippetsWorkspaceEncodedIdWatchNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNoContent) GetPayload() *models.PaginatedUsers {
	return o.Payload
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnippetsWorkspaceEncodedIDWatchUnauthorized creates a DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized with default headers values
func NewDeleteSnippetsWorkspaceEncodedIDWatchUnauthorized() *DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized {
	return &DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized{}
}

/*DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized handles this case with default header values.

If the request was not authenticated.
*/
type DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{workspace}/{encoded_id}/watch][%d] deleteSnippetsWorkspaceEncodedIdWatchUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnippetsWorkspaceEncodedIDWatchNotFound creates a DeleteSnippetsWorkspaceEncodedIDWatchNotFound with default headers values
func NewDeleteSnippetsWorkspaceEncodedIDWatchNotFound() *DeleteSnippetsWorkspaceEncodedIDWatchNotFound {
	return &DeleteSnippetsWorkspaceEncodedIDWatchNotFound{}
}

/*DeleteSnippetsWorkspaceEncodedIDWatchNotFound handles this case with default header values.

If the snippet does not exist.
*/
type DeleteSnippetsWorkspaceEncodedIDWatchNotFound struct {
	Payload *models.Error
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{workspace}/{encoded_id}/watch][%d] deleteSnippetsWorkspaceEncodedIdWatchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSnippetsWorkspaceEncodedIDWatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}