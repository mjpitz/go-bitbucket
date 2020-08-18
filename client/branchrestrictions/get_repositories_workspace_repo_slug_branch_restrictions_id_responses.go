// Code generated by go-swagger; DO NOT EDIT.

package branchrestrictions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader is a Reader for the GetRepositoriesWorkspaceRepoSlugBranchRestrictionsID structure.
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK creates a GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK() *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK {
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK handles this case with default header values.

The branch restriction rule
*/
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK struct {
	Payload *models.Branchrestriction
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] getRepositoriesWorkspaceRepoSlugBranchRestrictionsIdOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) GetPayload() *models.Branchrestriction {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branchrestriction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized creates a GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized() *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized handles this case with default header values.

If the request was not authenticated
*/
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] getRepositoriesWorkspaceRepoSlugBranchRestrictionsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden creates a GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden() *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden {
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden handles this case with default header values.

If the authenticated user does not have admin access to the repository
*/
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] getRepositoriesWorkspaceRepoSlugBranchRestrictionsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound creates a GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound() *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound {
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound handles this case with default header values.

If the repository or branch restriction id does not exist
*/
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] getRepositoriesWorkspaceRepoSlugBranchRestrictionsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}