// Code generated by go-swagger; DO NOT EDIT.

package pullrequests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstat structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound handles this case with default header values.

Redirects to the [repository diffstat](../../diffstat/%7Bspec%7D) with
the revspec that corresponds to pull request.

*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound struct {
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diffstat][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatFound ", 302)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDDiffstatFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
