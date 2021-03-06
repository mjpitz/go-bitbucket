// Code generated by go-swagger; DO NOT EDIT.

package branching_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// GetRepositoriesWorkspaceRepoSlugBranchingModelReader is a Reader for the GetRepositoriesWorkspaceRepoSlugBranchingModel structure.
type GetRepositoriesWorkspaceRepoSlugBranchingModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchingModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchingModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRepositoriesWorkspaceRepoSlugBranchingModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelOK creates a GetRepositoriesWorkspaceRepoSlugBranchingModelOK with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelOK() *GetRepositoriesWorkspaceRepoSlugBranchingModelOK {
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelOK{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchingModelOK handles this case with default header values.

The branching model object
*/
type GetRepositoriesWorkspaceRepoSlugBranchingModelOK struct {
	Payload *models.BranchingModel
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branching-model][%d] getRepositoriesWorkspaceRepoSlugBranchingModelOK  %+v", 200, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelOK) GetPayload() *models.BranchingModel {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BranchingModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized creates a GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized() *GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized {
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized handles this case with default header values.

If the request was not authenticated
*/
type GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branching-model][%d] getRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelForbidden creates a GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelForbidden() *GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden {
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden handles this case with default header values.

If the authenticated user does not have read access to the repository
*/
type GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branching-model][%d] getRepositoriesWorkspaceRepoSlugBranchingModelForbidden  %+v", 403, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelNotFound creates a GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound with default headers values
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelNotFound() *GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound {
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound{}
}

/*GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound handles this case with default header values.

If the repository does not exist
*/
type GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound struct {
	Payload *models.Error
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/branching-model][%d] getRepositoriesWorkspaceRepoSlugBranchingModelNotFound  %+v", 404, o.Payload)
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
