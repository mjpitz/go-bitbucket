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

// GetRepositoriesWorkspaceRepoSlugVersionsVersionIDReader is a Reader for the GetRepositoriesWorkspaceRepoSlugVersionsVersionID structure.
type GetRepositoriesWorkspaceRepoSlugVersionsVersionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK creates a GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK() *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK {
	return &GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK{}
}

/*GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK handles this case with default header values.

The specified version object.
*/
type GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK struct {
	Payload *models.Version
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/versions/{version_id}][%d] getRepositoriesWorkspaceRepoSlugVersionsVersionIdOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound creates a GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound() *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound {
	return &GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound handles this case with default header values.

The specified repository or version does not exist or does not have the issue tracker enabled.
*/
type GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/versions/{version_id}][%d] getRepositoriesWorkspaceRepoSlugVersionsVersionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugVersionsVersionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
