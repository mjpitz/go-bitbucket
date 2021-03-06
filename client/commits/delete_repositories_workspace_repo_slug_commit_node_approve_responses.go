// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveReader is a Reader for the DeleteRepositoriesWorkspaceRepoSlugCommitNodeApprove structure.
type DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent creates a DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent() *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent {
	return &DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent{}
}

/*DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent handles this case with default header values.

An empty response indicating the authenticated user's approval has been withdrawn.
*/
type DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent struct {
}

func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/commit/{node}/approve][%d] deleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent ", 204)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound creates a DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound with default headers values
func NewDeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound() *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound {
	return &DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound{}
}

/*DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound handles this case with default header values.

If the specified commit, or the repository does not exist.
*/
type DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound struct {
	Payload *models.Error
}

func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{workspace}/{repo_slug}/commit/{node}/approve][%d] deleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
