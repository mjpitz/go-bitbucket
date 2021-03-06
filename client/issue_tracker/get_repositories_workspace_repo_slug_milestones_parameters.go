// Code generated by go-swagger; DO NOT EDIT.

package issue_tracker

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

// NewGetRepositoriesWorkspaceRepoSlugMilestonesParams creates a new GetRepositoriesWorkspaceRepoSlugMilestonesParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugMilestonesParams() *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugMilestonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugMilestonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugMilestonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugMilestonesParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug milestones operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugMilestonesParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugMilestonesParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug milestones params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
