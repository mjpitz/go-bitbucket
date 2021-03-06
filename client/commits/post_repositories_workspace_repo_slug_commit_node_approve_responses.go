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

// PostRepositoriesWorkspaceRepoSlugCommitNodeApproveReader is a Reader for the PostRepositoriesWorkspaceRepoSlugCommitNodeApprove structure.
type PostRepositoriesWorkspaceRepoSlugCommitNodeApproveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK creates a PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK with default headers values
func NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK() *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK {
	return &PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK{}
}

/*PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK handles this case with default header values.

The `participant` object recording that the authenticated user approved the commit.
*/
type PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK struct {
	Payload *models.Participant
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/commit/{node}/approve][%d] postRepositoriesWorkspaceRepoSlugCommitNodeApproveOK  %+v", 200, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK) GetPayload() *models.Participant {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Participant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound creates a PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound with default headers values
func NewPostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound() *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound {
	return &PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound{}
}

/*PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound handles this case with default header values.

If the specified commit, or the repository does not exist.
*/
type PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/commit/{node}/approve][%d] postRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound  %+v", 404, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeApproveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
