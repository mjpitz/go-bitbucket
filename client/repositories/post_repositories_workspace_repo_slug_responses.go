// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mjpitz/go-bitbucket/models"
)

// PostRepositoriesWorkspaceRepoSlugReader is a Reader for the PostRepositoriesWorkspaceRepoSlug structure.
type PostRepositoriesWorkspaceRepoSlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesWorkspaceRepoSlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRepositoriesWorkspaceRepoSlugOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRepositoriesWorkspaceRepoSlugBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRepositoriesWorkspaceRepoSlugUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRepositoriesWorkspaceRepoSlugOK creates a PostRepositoriesWorkspaceRepoSlugOK with default headers values
func NewPostRepositoriesWorkspaceRepoSlugOK() *PostRepositoriesWorkspaceRepoSlugOK {
	return &PostRepositoriesWorkspaceRepoSlugOK{}
}

/*PostRepositoriesWorkspaceRepoSlugOK handles this case with default header values.

The newly created repository.
*/
type PostRepositoriesWorkspaceRepoSlugOK struct {
	Payload *models.Repository
}

func (o *PostRepositoriesWorkspaceRepoSlugOK) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}][%d] postRepositoriesWorkspaceRepoSlugOK  %+v", 200, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugBadRequest creates a PostRepositoriesWorkspaceRepoSlugBadRequest with default headers values
func NewPostRepositoriesWorkspaceRepoSlugBadRequest() *PostRepositoriesWorkspaceRepoSlugBadRequest {
	return &PostRepositoriesWorkspaceRepoSlugBadRequest{}
}

/*PostRepositoriesWorkspaceRepoSlugBadRequest handles this case with default header values.

If the input document was invalid, or if the caller lacks the privilege to create repositories under the targeted account.
*/
type PostRepositoriesWorkspaceRepoSlugBadRequest struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugBadRequest) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}][%d] postRepositoriesWorkspaceRepoSlugBadRequest  %+v", 400, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRepositoriesWorkspaceRepoSlugUnauthorized creates a PostRepositoriesWorkspaceRepoSlugUnauthorized with default headers values
func NewPostRepositoriesWorkspaceRepoSlugUnauthorized() *PostRepositoriesWorkspaceRepoSlugUnauthorized {
	return &PostRepositoriesWorkspaceRepoSlugUnauthorized{}
}

/*PostRepositoriesWorkspaceRepoSlugUnauthorized handles this case with default header values.

If the request was not authenticated.
*/
type PostRepositoriesWorkspaceRepoSlugUnauthorized struct {
	Payload *models.Error
}

func (o *PostRepositoriesWorkspaceRepoSlugUnauthorized) Error() string {
	return fmt.Sprintf("[POST /repositories/{workspace}/{repo_slug}][%d] postRepositoriesWorkspaceRepoSlugUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRepositoriesWorkspaceRepoSlugUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostRepositoriesWorkspaceRepoSlugUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
