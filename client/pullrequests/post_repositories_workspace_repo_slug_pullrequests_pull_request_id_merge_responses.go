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

// PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeReader is a Reader for the PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMerge structure.
type PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 555:
		result := NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK creates a PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK with default headers values
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK() *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK {
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK{}
}

/*PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK handles this case with default header values.

The pull request object.
*/
type PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK struct {
	Payload *models.Pullrequest
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge][%d] postRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeOK  %+v", 200, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK) GetPayload() *models.Pullrequest {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pullrequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted creates a PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted with default headers values
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted() *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted {
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted{}
}

/*PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted handles this case with default header values.

In the Location header, the URL to poll for the pull request merge status
*/
type PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted struct {
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge][%d] postRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeAccepted ", 202)
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555 creates a PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555 with default headers values
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555() *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555 {
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555{}
}

/*PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555 handles this case with default header values.

If the merge took too long and timed out.
In this case the caller should retry the request later
*/
type PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555 struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge][%d] postRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeStatus555  %+v", 555, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDMergeStatus555) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}