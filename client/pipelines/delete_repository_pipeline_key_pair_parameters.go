// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRepositoryPipelineKeyPairParams creates a new DeleteRepositoryPipelineKeyPairParams object
// with the default values initialized.
func NewDeleteRepositoryPipelineKeyPairParams() *DeleteRepositoryPipelineKeyPairParams {
	var ()
	return &DeleteRepositoryPipelineKeyPairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoryPipelineKeyPairParamsWithTimeout creates a new DeleteRepositoryPipelineKeyPairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoryPipelineKeyPairParamsWithTimeout(timeout time.Duration) *DeleteRepositoryPipelineKeyPairParams {
	var ()
	return &DeleteRepositoryPipelineKeyPairParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoryPipelineKeyPairParamsWithContext creates a new DeleteRepositoryPipelineKeyPairParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoryPipelineKeyPairParamsWithContext(ctx context.Context) *DeleteRepositoryPipelineKeyPairParams {
	var ()
	return &DeleteRepositoryPipelineKeyPairParams{

		Context: ctx,
	}
}

// NewDeleteRepositoryPipelineKeyPairParamsWithHTTPClient creates a new DeleteRepositoryPipelineKeyPairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoryPipelineKeyPairParamsWithHTTPClient(client *http.Client) *DeleteRepositoryPipelineKeyPairParams {
	var ()
	return &DeleteRepositoryPipelineKeyPairParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoryPipelineKeyPairParams contains all the parameters to send to the API endpoint
for the delete repository pipeline key pair operation typically these are written to a http.Request
*/
type DeleteRepositoryPipelineKeyPairParams struct {

	/*RepoSlug
	  The repository.

	*/
	RepoSlug string
	/*Username
	  The account.

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) WithTimeout(timeout time.Duration) *DeleteRepositoryPipelineKeyPairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) WithContext(ctx context.Context) *DeleteRepositoryPipelineKeyPairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) WithHTTPClient(client *http.Client) *DeleteRepositoryPipelineKeyPairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) WithRepoSlug(repoSlug string) *DeleteRepositoryPipelineKeyPairParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) WithUsername(username string) *DeleteRepositoryPipelineKeyPairParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete repository pipeline key pair params
func (o *DeleteRepositoryPipelineKeyPairParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoryPipelineKeyPairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}