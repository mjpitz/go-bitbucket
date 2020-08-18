// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDReader is a Reader for the GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentID structure.
type GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK creates a GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK() *GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK {
	return &GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK{}
}

/*GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK handles this case with default header values.

The commit comment.
*/
type GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK struct {
	Payload *models.CommitComment
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{node}/comments/{comment_id}][%d] getRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIdOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK) GetPayload() *models.CommitComment {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitNodeCommentsCommentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommitComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}