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

// NewGetPipelineForRepositoryParams creates a new GetPipelineForRepositoryParams object
// with the default values initialized.
func NewGetPipelineForRepositoryParams() *GetPipelineForRepositoryParams {
	var ()
	return &GetPipelineForRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineForRepositoryParamsWithTimeout creates a new GetPipelineForRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPipelineForRepositoryParamsWithTimeout(timeout time.Duration) *GetPipelineForRepositoryParams {
	var ()
	return &GetPipelineForRepositoryParams{

		timeout: timeout,
	}
}

// NewGetPipelineForRepositoryParamsWithContext creates a new GetPipelineForRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPipelineForRepositoryParamsWithContext(ctx context.Context) *GetPipelineForRepositoryParams {
	var ()
	return &GetPipelineForRepositoryParams{

		Context: ctx,
	}
}

// NewGetPipelineForRepositoryParamsWithHTTPClient creates a new GetPipelineForRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPipelineForRepositoryParamsWithHTTPClient(client *http.Client) *GetPipelineForRepositoryParams {
	var ()
	return &GetPipelineForRepositoryParams{
		HTTPClient: client,
	}
}

/*GetPipelineForRepositoryParams contains all the parameters to send to the API endpoint
for the get pipeline for repository operation typically these are written to a http.Request
*/
type GetPipelineForRepositoryParams struct {

	/*PipelineUUID
	  The pipeline UUID.

	*/
	PipelineUUID string
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

// WithTimeout adds the timeout to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithTimeout(timeout time.Duration) *GetPipelineForRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithContext(ctx context.Context) *GetPipelineForRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithHTTPClient(client *http.Client) *GetPipelineForRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPipelineUUID adds the pipelineUUID to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithPipelineUUID(pipelineUUID string) *GetPipelineForRepositoryParams {
	o.SetPipelineUUID(pipelineUUID)
	return o
}

// SetPipelineUUID adds the pipelineUuid to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetPipelineUUID(pipelineUUID string) {
	o.PipelineUUID = pipelineUUID
}

// WithRepoSlug adds the repoSlug to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithRepoSlug(repoSlug string) *GetPipelineForRepositoryParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) WithUsername(username string) *GetPipelineForRepositoryParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get pipeline for repository params
func (o *GetPipelineForRepositoryParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineForRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pipeline_uuid
	if err := r.SetPathParam("pipeline_uuid", o.PipelineUUID); err != nil {
		return err
	}

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
