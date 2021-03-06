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

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams() *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug pullrequests pull request ID commits operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPullRequestID adds the pullRequestID to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithPullRequestID(pullRequestID int64) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetPullRequestID(pullRequestID)
	return o
}

// SetPullRequestID adds the pullRequestId to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetPullRequestID(pullRequestID int64) {
	o.PullRequestID = pullRequestID
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug pullrequests pull request ID commits params
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDCommitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
