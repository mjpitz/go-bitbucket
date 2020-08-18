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

// GetSnippetsWorkspaceEncodedIDCommentsReader is a Reader for the GetSnippetsWorkspaceEncodedIDComments structure.
type GetSnippetsWorkspaceEncodedIDCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsWorkspaceEncodedIDCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnippetsWorkspaceEncodedIDCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSnippetsWorkspaceEncodedIDCommentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnippetsWorkspaceEncodedIDCommentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsWorkspaceEncodedIDCommentsOK creates a GetSnippetsWorkspaceEncodedIDCommentsOK with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommentsOK() *GetSnippetsWorkspaceEncodedIDCommentsOK {
	return &GetSnippetsWorkspaceEncodedIDCommentsOK{}
}

/*GetSnippetsWorkspaceEncodedIDCommentsOK handles this case with default header values.

A paginated list of snippet comments, ordered by creation date.
*/
type GetSnippetsWorkspaceEncodedIDCommentsOK struct {
	Payload *models.PaginatedSnippetComments
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsOK) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/comments][%d] getSnippetsWorkspaceEncodedIdCommentsOK  %+v", 200, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsOK) GetPayload() *models.PaginatedSnippetComments {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedSnippetComments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDCommentsForbidden creates a GetSnippetsWorkspaceEncodedIDCommentsForbidden with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommentsForbidden() *GetSnippetsWorkspaceEncodedIDCommentsForbidden {
	return &GetSnippetsWorkspaceEncodedIDCommentsForbidden{}
}

/*GetSnippetsWorkspaceEncodedIDCommentsForbidden handles this case with default header values.

If the authenticated user does not have access to the snippet.
*/
type GetSnippetsWorkspaceEncodedIDCommentsForbidden struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsForbidden) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/comments][%d] getSnippetsWorkspaceEncodedIdCommentsForbidden  %+v", 403, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDCommentsNotFound creates a GetSnippetsWorkspaceEncodedIDCommentsNotFound with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommentsNotFound() *GetSnippetsWorkspaceEncodedIDCommentsNotFound {
	return &GetSnippetsWorkspaceEncodedIDCommentsNotFound{}
}

/*GetSnippetsWorkspaceEncodedIDCommentsNotFound handles this case with default header values.

If the snippet does not exist.
*/
type GetSnippetsWorkspaceEncodedIDCommentsNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/comments][%d] getSnippetsWorkspaceEncodedIdCommentsNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}