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

// NewStopPipelineParams creates a new StopPipelineParams object
// with the default values initialized.
func NewStopPipelineParams() *StopPipelineParams {
	var ()
	return &StopPipelineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPipelineParamsWithTimeout creates a new StopPipelineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPipelineParamsWithTimeout(timeout time.Duration) *StopPipelineParams {
	var ()
	return &StopPipelineParams{

		timeout: timeout,
	}
}

// NewStopPipelineParamsWithContext creates a new StopPipelineParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPipelineParamsWithContext(ctx context.Context) *StopPipelineParams {
	var ()
	return &StopPipelineParams{

		Context: ctx,
	}
}

// NewStopPipelineParamsWithHTTPClient creates a new StopPipelineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPipelineParamsWithHTTPClient(client *http.Client) *StopPipelineParams {
	var ()
	return &StopPipelineParams{
		HTTPClient: client,
	}
}

/*StopPipelineParams contains all the parameters to send to the API endpoint
for the stop pipeline operation typically these are written to a http.Request
*/
type StopPipelineParams struct {

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

// WithTimeout adds the timeout to the stop pipeline params
func (o *StopPipelineParams) WithTimeout(timeout time.Duration) *StopPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop pipeline params
func (o *StopPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop pipeline params
func (o *StopPipelineParams) WithContext(ctx context.Context) *StopPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop pipeline params
func (o *StopPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop pipeline params
func (o *StopPipelineParams) WithHTTPClient(client *http.Client) *StopPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop pipeline params
func (o *StopPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPipelineUUID adds the pipelineUUID to the stop pipeline params
func (o *StopPipelineParams) WithPipelineUUID(pipelineUUID string) *StopPipelineParams {
	o.SetPipelineUUID(pipelineUUID)
	return o
}

// SetPipelineUUID adds the pipelineUuid to the stop pipeline params
func (o *StopPipelineParams) SetPipelineUUID(pipelineUUID string) {
	o.PipelineUUID = pipelineUUID
}

// WithRepoSlug adds the repoSlug to the stop pipeline params
func (o *StopPipelineParams) WithRepoSlug(repoSlug string) *StopPipelineParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the stop pipeline params
func (o *StopPipelineParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the stop pipeline params
func (o *StopPipelineParams) WithUsername(username string) *StopPipelineParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the stop pipeline params
func (o *StopPipelineParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *StopPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
