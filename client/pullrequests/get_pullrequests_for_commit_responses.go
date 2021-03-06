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

// GetPullrequestsForCommitReader is a Reader for the GetPullrequestsForCommit structure.
type GetPullrequestsForCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPullrequestsForCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPullrequestsForCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetPullrequestsForCommitAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPullrequestsForCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPullrequestsForCommitOK creates a GetPullrequestsForCommitOK with default headers values
func NewGetPullrequestsForCommitOK() *GetPullrequestsForCommitOK {
	return &GetPullrequestsForCommitOK{}
}

/*GetPullrequestsForCommitOK handles this case with default header values.

The paginated list of pull requests.
*/
type GetPullrequestsForCommitOK struct {
	Payload *models.PaginatedPullrequests
}

func (o *GetPullrequestsForCommitOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/pullrequests][%d] getPullrequestsForCommitOK  %+v", 200, o.Payload)
}

func (o *GetPullrequestsForCommitOK) GetPayload() *models.PaginatedPullrequests {
	return o.Payload
}

func (o *GetPullrequestsForCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPullrequests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPullrequestsForCommitAccepted creates a GetPullrequestsForCommitAccepted with default headers values
func NewGetPullrequestsForCommitAccepted() *GetPullrequestsForCommitAccepted {
	return &GetPullrequestsForCommitAccepted{}
}

/*GetPullrequestsForCommitAccepted handles this case with default header values.

The repository's pull requests are still being indexed.
*/
type GetPullrequestsForCommitAccepted struct {
	Payload *models.PaginatedPullrequests
}

func (o *GetPullrequestsForCommitAccepted) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/pullrequests][%d] getPullrequestsForCommitAccepted  %+v", 202, o.Payload)
}

func (o *GetPullrequestsForCommitAccepted) GetPayload() *models.PaginatedPullrequests {
	return o.Payload
}

func (o *GetPullrequestsForCommitAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedPullrequests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPullrequestsForCommitNotFound creates a GetPullrequestsForCommitNotFound with default headers values
func NewGetPullrequestsForCommitNotFound() *GetPullrequestsForCommitNotFound {
	return &GetPullrequestsForCommitNotFound{}
}

/*GetPullrequestsForCommitNotFound handles this case with default header values.

Either the repository does not exist, or pull request commit links have not yet been indexed.
*/
type GetPullrequestsForCommitNotFound struct {
	Payload *models.Error
}

func (o *GetPullrequestsForCommitNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{workspace}/{repo_slug}/commit/{commit}/pullrequests][%d] getPullrequestsForCommitNotFound  %+v", 404, o.Payload)
}

func (o *GetPullrequestsForCommitNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPullrequestsForCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
