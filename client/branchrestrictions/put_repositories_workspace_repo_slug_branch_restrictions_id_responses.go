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

// PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader is a Reader for the PutRepositoriesWorkspaceRepoSlugBranchRestrictionsID structure.
type PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK creates a PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK with default headers values
func NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK() *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK {
	return &PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK{}
}

/*PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK handles this case with default header values.

The updated branch restriction rule
*/
type PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK struct {
	Payload *models.Branchrestriction
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] putRepositoriesWorkspaceRepoSlugBranchRestrictionsIdOK  %+v", 200, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) GetPayload() *models.Branchrestriction {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branchrestriction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized creates a PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized with default headers values
func NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized() *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized {
	return &PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized{}
}

/*PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized handles this case with default header values.

If the request was not authenticated
*/
type PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] putRepositoriesWorkspaceRepoSlugBranchRestrictionsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden creates a PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden with default headers values
func NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden() *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden {
	return &PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden{}
}

/*PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden handles this case with default header values.

If the authenticated user does not have admin access to the repository
*/
type PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] putRepositoriesWorkspaceRepoSlugBranchRestrictionsIdForbidden  %+v", 403, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound creates a PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound with default headers values
func NewPutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound() *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound {
	return &PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound{}
}

/*PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound handles this case with default header values.

If the repository or branch restriction id does not exist
*/
type PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound struct {
	Payload *models.Error
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /repositories/{workspace}/{repo_slug}/branch-restrictions/{id}][%d] putRepositoriesWorkspaceRepoSlugBranchRestrictionsIdNotFound  %+v", 404, o.Payload)
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutRepositoriesWorkspaceRepoSlugBranchRestrictionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
