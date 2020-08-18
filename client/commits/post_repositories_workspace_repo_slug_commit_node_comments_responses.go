// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsReader is a Reader for the PostRepositoriesWorkspaceRepoSlugCommitNodeComments structure.
type PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated creates a PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated with default headers values
func NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated() *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated {
	return &PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated{}
}

/*PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated handles this case with default header values.

The newly created comment.
*/
type PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated struct {
	/*The location of the newly created comment.
	 */
	Location string
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/commit/{node}/comments][%d] postRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated ", 201)
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest creates a PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest with default headers values
func NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest() *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest {
	return &PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest{}
}

/*PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest handles this case with default header values.

If the comment was detected as spam, or if the parent comment is not attached to the same node as the new comment
*/
type PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest struct {
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/commit/{node}/comments][%d] postRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest ", 400)
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound creates a PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound with default headers values
func NewPostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound() *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound {
	return &PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound{}
}

/*PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound handles this case with default header values.

If a parent ID was passed in that cannot be found
*/
type PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound struct {
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/commit/{node}/comments][%d] postRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound ", 404)
}

func (o *PostRepositoriesWorkspaceRepoSlugCommitNodeCommentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
