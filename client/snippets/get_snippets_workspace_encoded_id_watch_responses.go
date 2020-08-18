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

// GetSnippetsWorkspaceEncodedIDWatchReader is a Reader for the GetSnippetsWorkspaceEncodedIDWatch structure.
type GetSnippetsWorkspaceEncodedIDWatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsWorkspaceEncodedIDWatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetSnippetsWorkspaceEncodedIDWatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSnippetsWorkspaceEncodedIDWatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchNoContent creates a GetSnippetsWorkspaceEncodedIDWatchNoContent with default headers values
func NewGetSnippetsWorkspaceEncodedIDWatchNoContent() *GetSnippetsWorkspaceEncodedIDWatchNoContent {
	return &GetSnippetsWorkspaceEncodedIDWatchNoContent{}
}

/*GetSnippetsWorkspaceEncodedIDWatchNoContent handles this case with default header values.

If the authenticated user is watching the snippet.
*/
type GetSnippetsWorkspaceEncodedIDWatchNoContent struct {
	Payload *models.PaginatedUsers
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNoContent) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/watch][%d] getSnippetsWorkspaceEncodedIdWatchNoContent  %+v", 204, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNoContent) GetPayload() *models.PaginatedUsers {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDWatchNotFound creates a GetSnippetsWorkspaceEncodedIDWatchNotFound with default headers values
func NewGetSnippetsWorkspaceEncodedIDWatchNotFound() *GetSnippetsWorkspaceEncodedIDWatchNotFound {
	return &GetSnippetsWorkspaceEncodedIDWatchNotFound{}
}

/*GetSnippetsWorkspaceEncodedIDWatchNotFound handles this case with default header values.

If the snippet does not exist, or if the authenticated user is not watching the snippet.
*/
type GetSnippetsWorkspaceEncodedIDWatchNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/watch][%d] getSnippetsWorkspaceEncodedIdWatchNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDWatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}