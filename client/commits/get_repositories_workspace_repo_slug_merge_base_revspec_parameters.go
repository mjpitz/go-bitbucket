// Code generated by go-swagger; DO NOT EDIT.

package commits

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

// NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams creates a new GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams() *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug merge base revspec operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams struct {

	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Revspec
	  A commit range using double dot notation (e.g. `3a8b42..9ff173`).


	*/
	Revspec string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithRevspec adds the revspec to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithRevspec(revspec string) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetRevspec(revspec)
	return o
}

// SetRevspec adds the revspec to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetRevspec(revspec string) {
	o.Revspec = revspec
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug merge base revspec params
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugMergeBaseRevspecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param revspec
	if err := r.SetPathParam("revspec", o.Revspec); err != nil {
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
