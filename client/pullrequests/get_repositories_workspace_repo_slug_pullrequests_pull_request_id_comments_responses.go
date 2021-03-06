// Code generated by go-swagger; DO NOT EDIT.

package pullrequests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDComments structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK handles this case with default header values.

A paginated list of comments made on the given pull request, in chronological order.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK struct {
	Payload *models.PaginatedPullrequestComments
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK) GetPayload() *models.PaginatedPullrequestComments {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPullrequestComments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden handles this case with default header values.

If the authenticated user does not have access to the pull request.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound handles this case with default header values.

If the pull request does not exist.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
