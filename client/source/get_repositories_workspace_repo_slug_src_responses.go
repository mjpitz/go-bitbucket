// Code generated by go-swagger; DO NOT EDIT.

package source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugSrcReader is a Reader for the GetRepositoriesWorkspaceRepoSlugSrc structure.
type GetRepositoriesWorkspaceRepoSlugSrcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugSrcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugSrcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugSrcNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugSrcOK creates a GetRepositoriesWorkspaceRepoSlugSrcOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugSrcOK() *GetRepositoriesWorkspaceRepoSlugSrcOK {
	return &GetRepositoriesWorkspaceRepoSlugSrcOK{}
}

/*GetRepositoriesWorkspaceRepoSlugSrcOK handles this case with default header values.

If the path matches a file, then the raw contents of the file are
returned (unless the `format=meta` query parameter was provided,
in which case a json document containing the file's meta data is
returned). If the path matches a directory, then a paginated
list of file and directory entries is returned (if the
`format=meta` query parameter was provided, then the json document
containing the directory's meta data is returned).

*/
type GetRepositoriesWorkspaceRepoSlugSrcOK struct {
	Payload *models.PaginatedTreeentries
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/src][%d] getRepositoriesWorkspaceRepoSlugSrcOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcOK) GetPayload() *models.PaginatedTreeentries {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedTreeentries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugSrcNotFound creates a GetRepositoriesWorkspaceRepoSlugSrcNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugSrcNotFound() *GetRepositoriesWorkspaceRepoSlugSrcNotFound {
	return &GetRepositoriesWorkspaceRepoSlugSrcNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugSrcNotFound handles this case with default header values.

If the path or commit in the URL does not exist.
*/
type GetRepositoriesWorkspaceRepoSlugSrcNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/src][%d] getRepositoriesWorkspaceRepoSlugSrcNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugSrcNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
