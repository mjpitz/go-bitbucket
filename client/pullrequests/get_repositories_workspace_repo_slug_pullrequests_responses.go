// Code generated by go-swagger; DO NOT EDIT.

package pullrequests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugPullrequestsReader is a Reader for the GetRepositoriesWorkspaceRepoSlugPullrequests structure.
type GetRepositoriesWorkspaceRepoSlugPullrequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugPullrequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsOK creates a GetRepositoriesWorkspaceRepoSlugPullrequestsOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsOK() *GetRepositoriesWorkspaceRepoSlugPullrequestsOK {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsOK{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsOK handles this case with default header values.

All pull requests on the specified repository.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsOK struct {
	Payload *models.PaginatedPullrequests
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests][%d] getRepositoriesWorkspaceRepoSlugPullrequestsOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsOK) GetPayload() *models.PaginatedPullrequests {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPullrequests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized creates a GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized() *GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized handles this case with default header values.

If the repository is private and the request was not authenticated.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized struct {
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests][%d] getRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized ", 401)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsNotFound creates a GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsNotFound() *GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound {
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound handles this case with default header values.

If the specified repository does not exist.
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/pullrequests][%d] getRepositoriesWorkspaceRepoSlugPullrequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
