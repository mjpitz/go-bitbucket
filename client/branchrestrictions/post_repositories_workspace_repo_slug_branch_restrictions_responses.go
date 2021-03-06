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

// PostRepositoriesWorkspaceRepoSlugBranchRestrictionsReader is a Reader for the PostRepositoriesWorkspaceRepoSlugBranchRestrictions structure.
type PostRepositoriesWorkspaceRepoSlugBranchRestrictionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated creates a PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated with default headers values
func NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated() *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated {
	return &PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated{}
}

/*PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated handles this case with default header values.

A paginated list of branch restrictions
*/
type PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated struct {
	Payload *models.Branchrestriction
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/branch-restrictions][%d] postRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated  %+v", 201, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated) GetPayload() *models.Branchrestriction {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branchrestriction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized creates a PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized with default headers values
func NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized() *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized {
	return &PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized{}
}

/*PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized handles this case with default header values.

If the request was not authenticated
*/
type PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/branch-restrictions][%d] postRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden creates a PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden with default headers values
func NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden() *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden {
	return &PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden{}
}

/*PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden handles this case with default header values.

If the authenticated user does not have admin access to the repository
*/
type PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/branch-restrictions][%d] postRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden  %+v", 403, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound creates a PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound with default headers values
func NewPostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound() *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound {
	return &PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound{}
}

/*PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound handles this case with default header values.

If the repository does not exist
*/
type PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}/branch-restrictions][%d] postRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound  %+v", 404, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugBranchRestrictionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
