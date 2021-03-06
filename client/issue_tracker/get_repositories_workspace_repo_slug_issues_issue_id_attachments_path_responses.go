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

// GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathReader is a Reader for the GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPath structure.
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound handles this case with default header values.

A redirect to the file's contents
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound struct {
	Location string
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments/{path}][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathFound ", 302)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized handles this case with default header values.

If the issue tracker is private and the request was not authenticated.
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized struct {
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments/{path}][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathUnauthorized ", 401)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound creates a GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound {
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound handles this case with default header values.

The specified repository or issue does not exist or does not have the issue tracker enabled.
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments/{path}][%d] getRepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
