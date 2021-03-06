// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// PostRepositoriesWorkspaceRepoSlugIssuesExportReader is a Reader for the PostRepositoriesWorkspaceRepoSlugIssuesExport structure.
type PostRepositoriesWorkspaceRepoSlugIssuesExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostRepositoriesWorkspaceRepoSlugIssuesExportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesWorkspaceRepoSlugIssuesExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRepositoriesWorkspaceRepoSlugIssuesExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportAccepted creates a PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted with default headers values
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportAccepted() *PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted {
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted{}
}

/*PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted handles this case with default header values.

The export job has been accepted
*/
type PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted struct {
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/issues/export][%d] postRepositoriesWorkspaceRepoSlugIssuesExportAccepted ", 202)
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized creates a PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized with default headers values
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized() *PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized {
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized{}
}

/*PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized handles this case with default header values.

The request wasn't authenticated properly
*/
type PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/issues/export][%d] postRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportForbidden creates a PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden with default headers values
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportForbidden() *PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden {
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden{}
}

/*PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden handles this case with default header values.

When the authenticated user does not have admin permission on the repo
*/
type PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/issues/export][%d] postRepositoriesWorkspaceRepoSlugIssuesExportForbidden  %+v", 403, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportNotFound creates a PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound with default headers values
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportNotFound() *PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound {
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound{}
}

/*PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound handles this case with default header values.

The repo does not exist or does not have an issue tracker
*/
type PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/issues/export][%d] postRepositoriesWorkspaceRepoSlugIssuesExportNotFound  %+v", 404, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
