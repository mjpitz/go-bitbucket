// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// DeleteRepositoriesWorkspaceRepoSlugHooksUIDReader is a Reader for the DeleteRepositoriesWorkspaceRepoSlugHooksUID structure.
type DeleteRepositoriesWorkspaceRepoSlugHooksUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent creates a DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent() *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent {
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent{}
}

/*DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent handles this case with default header values.

When the webhook was deleted successfully
*/
type DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent struct {
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/hooks/{uid}][%d] deleteRepositoriesWorkspaceRepoSlugHooksUidNoContent ", 204)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden creates a DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden() *DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden {
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden{}
}

/*DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden handles this case with default header values.

If the authenticated user does not have permission to delete the webhook.
*/
type DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/hooks/{uid}][%d] deleteRepositoriesWorkspaceRepoSlugHooksUidForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound creates a DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound() *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound {
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound{}
}

/*DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound handles this case with default header values.

If the webhook or repository does not exist.
*/
type DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/hooks/{uid}][%d] deleteRepositoriesWorkspaceRepoSlugHooksUidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}