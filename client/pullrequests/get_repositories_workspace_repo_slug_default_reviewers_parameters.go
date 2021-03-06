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

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParams creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParams() *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug default reviewers operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug default reviewers params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
