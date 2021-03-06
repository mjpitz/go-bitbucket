// Code generated by go-swagger; DO NOT EDIT.

package issue_tracker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader is a Reader for the GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatch structure.
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent handles this case with default header values.

If the authenticated user is watching this issue.
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchNoContent  %+v", 204, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized handles this case with default header values.

When the request wasn't authenticated.
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound handles this case with default header values.

If the authenticated user is not watching this issue, or when the repo does not exist, or does not have an issue tracker.
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
