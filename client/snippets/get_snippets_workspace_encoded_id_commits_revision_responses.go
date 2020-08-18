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

// GetSnippetsWorkspaceEncodedIDCommitsRevisionReader is a Reader for the GetSnippetsWorkspaceEncodedIDCommitsRevision structure.
type GetSnippetsWorkspaceEncodedIDCommitsRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnippetsWorkspaceEncodedIDCommitsRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsWorkspaceEncodedIDCommitsRevisionOK creates a GetSnippetsWorkspaceEncodedIDCommitsRevisionOK with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommitsRevisionOK() *GetSnippetsWorkspaceEncodedIDCommitsRevisionOK {
	return &GetSnippetsWorkspaceEncodedIDCommitsRevisionOK{}
}

/*GetSnippetsWorkspaceEncodedIDCommitsRevisionOK handles this case with default header values.

The specified snippet commit.
*/
type GetSnippetsWorkspaceEncodedIDCommitsRevisionOK struct {
	Payload *models.SnippetCommit
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionOK) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/commits/{revision}][%d] getSnippetsWorkspaceEncodedIdCommitsRevisionOK  %+v", 200, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionOK) GetPayload() *models.SnippetCommit {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnippetCommit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden creates a GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden() *GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden {
	return &GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden{}
}

/*GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden handles this case with default header values.

If the authenticated user does not have access to the snippet.
*/
type GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/commits/{revision}][%d] getSnippetsWorkspaceEncodedIdCommitsRevisionForbidden  %+v", 403, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound creates a GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound with default headers values
func NewGetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound() *GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound {
	return &GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound{}
}

/*GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound handles this case with default header values.

If the commit or the snippet does not exist.
*/
type GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/commits/{revision}][%d] getSnippetsWorkspaceEncodedIdCommitsRevisionNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDCommitsRevisionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
