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

// NewGetRepositoriesWorkspaceRepoSlugVersionsParams creates a new GetRepositoriesWorkspaceRepoSlugVersionsParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugVersionsParams() *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugVersionsParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugVersionsParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugVersionsParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugVersionsParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugVersionsParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug versions operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugVersionsParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugVersionsParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug versions params
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
