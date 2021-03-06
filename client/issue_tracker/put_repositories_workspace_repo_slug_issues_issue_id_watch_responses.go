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

// PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader is a Reader for the PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatch structure.
type PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent creates a PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent with default headers values
func NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent() *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent {
	return &PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent{}
}

/*PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent handles this case with default header values.

Indicates that the authenticated user successfully started watching this issue.
*/
type PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] putRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchNoContent  %+v", 204, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized creates a PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized with default headers values
func NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized() *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized {
	return &PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized{}
}

/*PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized handles this case with default header values.

When the request wasn't authenticated.
*/
type PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] putRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound creates a PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound with default headers values
func NewPutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound() *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound {
	return &PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound{}
}

/*PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound handles this case with default header values.

If the authenticated user is not watching this issue, or when the repo does not exist, or does not have an issue tracker.
*/
type PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch][%d] putRepositoriesWorkspaceRepoSlugIssuesIssueIdWatchNotFound  %+v", 404, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
