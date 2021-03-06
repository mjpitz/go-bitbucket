// Code generated by go-swagger; DO NOT EDIT.

package deploy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDReader is a Reader for the DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyID structure.
type DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent creates a DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent() *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent {
	return &DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent{}
}

/*DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent handles this case with default header values.

The key has been deleted
*/
type DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent struct {
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id}][%d] deleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIdNoContent ", 204)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden creates a DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden() *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden {
	return &DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden{}
}

/*DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden handles this case with default header values.

If the current user does not have permission to delete a key for the specified user
*/
type DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden struct {
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id}][%d] deleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIdForbidden ", 403)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound creates a DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound() *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound {
	return &DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound{}
}

/*DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound handles this case with default header values.

If the specified user, repository, or deploy key does not exist
*/
type DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id}][%d] deleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRepositoriesWorkspaceRepoSlugDeployKeysKeyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
