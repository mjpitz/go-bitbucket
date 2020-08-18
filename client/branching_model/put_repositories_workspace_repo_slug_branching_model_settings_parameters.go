// Code generated by go-swagger; DO NOT EDIT.

package branching_model

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

// NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams creates a new PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams object
// with the default values initialized.
func NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams() *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithTimeout creates a new PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams{

		timeout: timeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithContext creates a new PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams{

		Context: ctx,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithHTTPClient creates a new PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParamsWithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams{
		HTTPClient: client,
	}
}

/*PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams contains all the parameters to send to the API endpoint
for the put repositories workspace repo slug branching model settings operation typically these are written to a http.Request
*/
type PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams struct {

	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WithRepoSlug(repoSlug string) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WithWorkspace(workspace string) *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the put repositories workspace repo slug branching model settings params
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesWorkspaceRepoSlugBranchingModelSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param workspace
	if err := r.SetPathParam("workspace", o.Workspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
