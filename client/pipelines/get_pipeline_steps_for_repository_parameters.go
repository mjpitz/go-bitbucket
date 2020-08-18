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

// NewGetPipelineStepsForRepositoryParams creates a new GetPipelineStepsForRepositoryParams object
// with the default values initialized.
func NewGetPipelineStepsForRepositoryParams() *GetPipelineStepsForRepositoryParams {
	var ()
	return &GetPipelineStepsForRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineStepsForRepositoryParamsWithTimeout creates a new GetPipelineStepsForRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPipelineStepsForRepositoryParamsWithTimeout(timeout time.Duration) *GetPipelineStepsForRepositoryParams {
	var ()
	return &GetPipelineStepsForRepositoryParams{

		timeout: timeout,
	}
}

// NewGetPipelineStepsForRepositoryParamsWithContext creates a new GetPipelineStepsForRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPipelineStepsForRepositoryParamsWithContext(ctx context.Context) *GetPipelineStepsForRepositoryParams {
	var ()
	return &GetPipelineStepsForRepositoryParams{

		Context: ctx,
	}
}

// NewGetPipelineStepsForRepositoryParamsWithHTTPClient creates a new GetPipelineStepsForRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPipelineStepsForRepositoryParamsWithHTTPClient(client *http.Client) *GetPipelineStepsForRepositoryParams {
	var ()
	return &GetPipelineStepsForRepositoryParams{
		HTTPClient: client,
	}
}

/*GetPipelineStepsForRepositoryParams contains all the parameters to send to the API endpoint
for the get pipeline steps for repository operation typically these are written to a http.Request
*/
type GetPipelineStepsForRepositoryParams struct {

	/*PipelineUUID
	  The UUID of the pipeline.

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

// WithTimeout adds the timeout to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithTimeout(timeout time.Duration) *GetPipelineStepsForRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithContext(ctx context.Context) *GetPipelineStepsForRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithHTTPClient(client *http.Client) *GetPipelineStepsForRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPipelineUUID adds the pipelineUUID to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithPipelineUUID(pipelineUUID string) *GetPipelineStepsForRepositoryParams {
	o.SetPipelineUUID(pipelineUUID)
	return o
}

// SetPipelineUUID adds the pipelineUuid to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetPipelineUUID(pipelineUUID string) {
	o.PipelineUUID = pipelineUUID
}

// WithRepoSlug adds the repoSlug to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithRepoSlug(repoSlug string) *GetPipelineStepsForRepositoryParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) WithUsername(username string) *GetPipelineStepsForRepositoryParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get pipeline steps for repository params
func (o *GetPipelineStepsForRepositoryParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineStepsForRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
