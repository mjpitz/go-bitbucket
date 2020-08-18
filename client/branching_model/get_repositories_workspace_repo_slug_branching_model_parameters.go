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

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelParams creates a new GetRepositoriesWorkspaceRepoSlugBranchingModelParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelParams() *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugBranchingModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugBranchingModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugBranchingModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugBranchingModelParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchingModelParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugBranchingModelParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug branching model operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugBranchingModelParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugBranchingModelParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug branching model params
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugBranchingModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
