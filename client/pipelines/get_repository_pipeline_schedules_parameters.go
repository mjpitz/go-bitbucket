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

// NewGetRepositoryPipelineSchedulesParams creates a new GetRepositoryPipelineSchedulesParams object
// with the default values initialized.
func NewGetRepositoryPipelineSchedulesParams() *GetRepositoryPipelineSchedulesParams {
	var ()
	return &GetRepositoryPipelineSchedulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoryPipelineSchedulesParamsWithTimeout creates a new GetRepositoryPipelineSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoryPipelineSchedulesParamsWithTimeout(timeout time.Duration) *GetRepositoryPipelineSchedulesParams {
	var ()
	return &GetRepositoryPipelineSchedulesParams{

		timeout: timeout,
	}
}

// NewGetRepositoryPipelineSchedulesParamsWithContext creates a new GetRepositoryPipelineSchedulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoryPipelineSchedulesParamsWithContext(ctx context.Context) *GetRepositoryPipelineSchedulesParams {
	var ()
	return &GetRepositoryPipelineSchedulesParams{

		Context: ctx,
	}
}

// NewGetRepositoryPipelineSchedulesParamsWithHTTPClient creates a new GetRepositoryPipelineSchedulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoryPipelineSchedulesParamsWithHTTPClient(client *http.Client) *GetRepositoryPipelineSchedulesParams {
	var ()
	return &GetRepositoryPipelineSchedulesParams{
		HTTPClient: client,
	}
}

/*GetRepositoryPipelineSchedulesParams contains all the parameters to send to the API endpoint
for the get repository pipeline schedules operation typically these are written to a http.Request
*/
type GetRepositoryPipelineSchedulesParams struct {

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

// WithTimeout adds the timeout to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) WithTimeout(timeout time.Duration) *GetRepositoryPipelineSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) WithContext(ctx context.Context) *GetRepositoryPipelineSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) WithHTTPClient(client *http.Client) *GetRepositoryPipelineSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) WithRepoSlug(repoSlug string) *GetRepositoryPipelineSchedulesParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) WithUsername(username string) *GetRepositoryPipelineSchedulesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get repository pipeline schedules params
func (o *GetRepositoryPipelineSchedulesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoryPipelineSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
