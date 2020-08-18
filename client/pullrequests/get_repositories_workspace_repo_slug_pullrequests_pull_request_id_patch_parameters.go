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
	"github.com/go-openapi/swag"
)

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug pullrequests pull request ID patch operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams struct {

	/*PullRequestID
	  The id of the pull request.

	*/
	PullRequestID int64
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

// WithTimeout adds the timeout to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPullRequestID adds the pullRequestID to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithPullRequestID(pullRequestID int64) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetPullRequestID(pullRequestID)
	return o
}

// SetPullRequestID adds the pullRequestId to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetPullRequestID(pullRequestID int64) {
	o.PullRequestID = pullRequestID
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug pullrequests pull request ID patch params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pull_request_id
	if err := r.SetPathParam("pull_request_id", swag.FormatInt64(o.PullRequestID)); err != nil {
		return err
	}

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