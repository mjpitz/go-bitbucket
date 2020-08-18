// Code generated by go-swagger; DO NOT EDIT.

package downloads

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

// NewGetRepositoriesWorkspaceRepoSlugDownloadsParams creates a new GetRepositoriesWorkspaceRepoSlugDownloadsParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugDownloadsParams() *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDownloadsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugDownloadsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDownloadsParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugDownloadsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDownloadsParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugDownloadsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugDownloadsParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDownloadsParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugDownloadsParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug downloads operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugDownloadsParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugDownloadsParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug downloads params
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugDownloadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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