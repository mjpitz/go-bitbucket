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

// PostRepositoriesWorkspaceRepoSlugRefsTagsReader is a Reader for the PostRepositoriesWorkspaceRepoSlugRefsTags structure.
type PostRepositoriesWorkspaceRepoSlugRefsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesWorkspaceRepoSlugRefsTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugRefsTagsCreated creates a PostRepositoriesWorkspaceRepoSlugRefsTagsCreated with default headers values
func NewPostRepositoriesWorkspaceRepoSlugRefsTagsCreated() *PostRepositoriesWorkspaceRepoSlugRefsTagsCreated {
	return &PostRepositoriesWorkspaceRepoSlugRefsTagsCreated{}
}

/*PostRepositoriesWorkspaceRepoSlugRefsTagsCreated handles this case with default header values.

The newly created tag.
*/
type PostRepositoriesWorkspaceRepoSlugRefsTagsCreated struct {
	Payload *models.Tag
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/refs/tags][%d] postRepositoriesWorkspaceRepoSlugRefsTagsCreated  %+v", 201, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsCreated) GetPayload() *models.Tag {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest creates a PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest with default headers values
func NewPostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest() *PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest {
	return &PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest{}
}

/*PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest handles this case with default header values.

If the target hash is missing, ambiguous, or invalid, or if the name is not provided.
*/
type PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/refs/tags][%d] postRepositoriesWorkspaceRepoSlugRefsTagsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugRefsTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}