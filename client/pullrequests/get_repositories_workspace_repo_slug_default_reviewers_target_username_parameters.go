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

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams() *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug default reviewers target username operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams struct {

	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*TargetUsername
	  This can either be the username or the UUID of the default reviewer,
	surrounded by curly-braces, for example: `{account UUID}`.


	*/
	TargetUsername string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithTargetUsername adds the targetUsername to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithTargetUsername(targetUsername string) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetTargetUsername(targetUsername)
	return o
}

// SetTargetUsername adds the targetUsername to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetTargetUsername(targetUsername string) {
	o.TargetUsername = targetUsername
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug default reviewers target username params
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param target_username
	if err := r.SetPathParam("target_username", o.TargetUsername); err != nil {
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
