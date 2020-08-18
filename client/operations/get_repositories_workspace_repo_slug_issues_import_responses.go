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

// GetRepositoriesWorkspaceRepoSlugIssuesImportReader is a Reader for the GetRepositoriesWorkspaceRepoSlugIssuesImport structure.
type GetRepositoriesWorkspaceRepoSlugIssuesImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesImportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesImportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugIssuesImportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesImportOK creates a GetRepositoriesWorkspaceRepoSlugIssuesImportOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesImportOK() *GetRepositoriesWorkspaceRepoSlugIssuesImportOK {
	return &GetRepositoriesWorkspaceRepoSlugIssuesImportOK{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesImportOK handles this case with default header values.

Import job complete with either FAILURE or SUCCESS status
*/
type GetRepositoriesWorkspaceRepoSlugIssuesImportOK struct {
	Payload *models.IssueJobStatus
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/import][%d] getRepositoriesWorkspaceRepoSlugIssuesImportOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportOK) GetPayload() *models.IssueJobStatus {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssueJobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesImportAccepted creates a GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesImportAccepted() *GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted {
	return &GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted handles this case with default header values.

Import job started
*/
type GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted struct {
	Payload *models.IssueJobStatus
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/import][%d] getRepositoriesWorkspaceRepoSlugIssuesImportAccepted  %+v", 202, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted) GetPayload() *models.IssueJobStatus {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssueJobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized creates a GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized() *GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized handles this case with default header values.

The request wasn't authenticated properly
*/
type GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/import][%d] getRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesImportForbidden creates a GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesImportForbidden() *GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden {
	return &GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden handles this case with default header values.

When the authenticated user does not have admin permission on the repo
*/
type GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/import][%d] getRepositoriesWorkspaceRepoSlugIssuesImportForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesImportNotFound creates a GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugIssuesImportNotFound() *GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound {
	return &GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound handles this case with default header values.

No export job has begun
*/
type GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/issues/import][%d] getRepositoriesWorkspaceRepoSlugIssuesImportNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugIssuesImportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}