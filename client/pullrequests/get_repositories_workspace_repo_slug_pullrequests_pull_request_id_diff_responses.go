// Code generated by go-swagger; DO NOT EDIT.

package pullrequests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiff structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound handles this case with default header values.

Redirects to the [repository diff](../../diff/%7Bspec%7D) with the
revspec that corresponds to the pull request.

*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound struct {
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diff][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffFound ", 302)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
