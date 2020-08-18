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

// NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams creates a new PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized.
func NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams() *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithTimeout creates a new PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		timeout: timeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithContext creates a new PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{

		Context: ctx,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithHTTPClient creates a new PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParamsWithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams{
		HTTPClient: client,
	}
}

/*PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams contains all the parameters to send to the API endpoint
for the put repositories workspace repo slug default reviewers target username operation typically these are written to a http.Request
*/
type PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams struct {

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

// WithTimeout adds the timeout to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithRepoSlug(repoSlug string) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithTargetUsername adds the targetUsername to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithTargetUsername(targetUsername string) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetTargetUsername(targetUsername)
	return o
}

// SetTargetUsername adds the targetUsername to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetTargetUsername(targetUsername string) {
	o.TargetUsername = targetUsername
}

// WithWorkspace adds the workspace to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WithWorkspace(workspace string) *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the put repositories workspace repo slug default reviewers target username params
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
