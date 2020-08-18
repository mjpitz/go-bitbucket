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

// GetSnippetsWorkspaceEncodedIDRevisionDiffReader is a Reader for the GetSnippetsWorkspaceEncodedIDRevisionDiff structure.
type GetSnippetsWorkspaceEncodedIDRevisionDiffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnippetsWorkspaceEncodedIDRevisionDiffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSnippetsWorkspaceEncodedIDRevisionDiffForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnippetsWorkspaceEncodedIDRevisionDiffNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnippetsWorkspaceEncodedIDRevisionDiffOK creates a GetSnippetsWorkspaceEncodedIDRevisionDiffOK with default headers values
func NewGetSnippetsWorkspaceEncodedIDRevisionDiffOK() *GetSnippetsWorkspaceEncodedIDRevisionDiffOK {
	return &GetSnippetsWorkspaceEncodedIDRevisionDiffOK{}
}

/*GetSnippetsWorkspaceEncodedIDRevisionDiffOK handles this case with default header values.

The raw diff contents.
*/
type GetSnippetsWorkspaceEncodedIDRevisionDiffOK struct {
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffOK) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{revision}/diff][%d] getSnippetsWorkspaceEncodedIdRevisionDiffOK ", 200)
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDRevisionDiffForbidden creates a GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden with default headers values
func NewGetSnippetsWorkspaceEncodedIDRevisionDiffForbidden() *GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden {
	return &GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden{}
}

/*GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden handles this case with default header values.

If the authenticated user does not have access to the snippet.
*/
type GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{revision}/diff][%d] getSnippetsWorkspaceEncodedIdRevisionDiffForbidden  %+v", 403, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnippetsWorkspaceEncodedIDRevisionDiffNotFound creates a GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound with default headers values
func NewGetSnippetsWorkspaceEncodedIDRevisionDiffNotFound() *GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound {
	return &GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound{}
}

/*GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound handles this case with default header values.

If the snippet does not exist.
*/
type GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound struct {
	Payload *models.Error
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound) Error() string {
	return fmt.Sprintf("[GET /snippets/{workspace}/{encoded_id}/{revision}/diff][%d] getSnippetsWorkspaceEncodedIdRevisionDiffNotFound  %+v", 404, o.Payload)
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSnippetsWorkspaceEncodedIDRevisionDiffNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}