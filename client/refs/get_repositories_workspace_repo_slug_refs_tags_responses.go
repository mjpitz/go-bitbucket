// Code generated by go-swagger; DO NOT EDIT.

package refs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugRefsTagsReader is a Reader for the GetRepositoriesWorkspaceRepoSlugRefsTags structure.
type GetRepositoriesWorkspaceRepoSlugRefsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugRefsTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRepositoriesWorkspaceRepoSlugRefsTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugRefsTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugRefsTagsOK creates a GetRepositoriesWorkspaceRepoSlugRefsTagsOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugRefsTagsOK() *GetRepositoriesWorkspaceRepoSlugRefsTagsOK {
	return &GetRepositoriesWorkspaceRepoSlugRefsTagsOK{}
}

/*GetRepositoriesWorkspaceRepoSlugRefsTagsOK handles this case with default header values.

A paginated list of tags matching any filter criteria that were provided.
*/
type GetRepositoriesWorkspaceRepoSlugRefsTagsOK struct {
	Payload *models.PaginatedTags
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/refs/tags][%d] getRepositoriesWorkspaceRepoSlugRefsTagsOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsOK) GetPayload() *models.PaginatedTags {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedTags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugRefsTagsForbidden creates a GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden with default headers values
func NewGetRepositoriesWorkspaceRepoSlugRefsTagsForbidden() *GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden {
	return &GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden{}
}

/*GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden handles this case with default header values.

If the repository is private and the authenticated user does not have
access to it.

*/
type GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/refs/tags][%d] getRepositoriesWorkspaceRepoSlugRefsTagsForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugRefsTagsNotFound creates a GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugRefsTagsNotFound() *GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound {
	return &GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound handles this case with default header values.

The specified repository does not exist.
*/
type GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/refs/tags][%d] getRepositoriesWorkspaceRepoSlugRefsTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugRefsTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}