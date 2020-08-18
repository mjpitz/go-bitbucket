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

// GetRepositoriesWorkspaceRepoSlugCommitsReader is a Reader for the GetRepositoriesWorkspaceRepoSlugCommits structure.
type GetRepositoriesWorkspaceRepoSlugCommitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugCommitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetRepositoriesWorkspaceRepoSlugCommitsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetRepositoriesWorkspaceRepoSlugCommitsDefault creates a GetRepositoriesWorkspaceRepoSlugCommitsDefault with default headers values
func NewGetRepositoriesWorkspaceRepoSlugCommitsDefault(code int) *GetRepositoriesWorkspaceRepoSlugCommitsDefault {
	return &GetRepositoriesWorkspaceRepoSlugCommitsDefault{
		_statusCode: code,
	}
}

/*GetRepositoriesWorkspaceRepoSlugCommitsDefault handles this case with default header values.

Unexpected error.
*/
type GetRepositoriesWorkspaceRepoSlugCommitsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get repositories workspace repo slug commits default response
func (o *GetRepositoriesWorkspaceRepoSlugCommitsDefault) Code() int {
	return o._statusCode
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitsDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commits][%d] GetRepositoriesWorkspaceRepoSlugCommits default  %+v", o._statusCode, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugCommitsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}