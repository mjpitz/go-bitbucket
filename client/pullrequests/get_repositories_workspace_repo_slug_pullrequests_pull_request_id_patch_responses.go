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

// GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatch structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault(code int) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault{
		_statusCode: code,
	}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault handles this case with default header values.

Unexpected error.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get repositories workspace repo slug pullrequests pull request ID patch default response
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault) Code() int {
	return o._statusCode
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/patch][%d] GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatch default  %+v", o._statusCode, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}