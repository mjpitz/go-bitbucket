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

// GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestID structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK handles this case with default header values.

The pull request object
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK struct {
	Payload *models.Pullrequest
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK) GetPayload() *models.Pullrequest {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pullrequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized handles this case with default header values.

If the repository is private and the request was not authenticated.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized struct {
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdUnauthorized ", 401)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound creates a GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound handles this case with default header values.

If the repository or pull request does not exist
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}][%d] getRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
