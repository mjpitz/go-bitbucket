// Code generated by go-swagger; DO NOT EDIT.

package pullrequests

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

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams() *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsActivityParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug pullrequests activity operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug pullrequests activity params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
