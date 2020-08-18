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

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams creates a new PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams object
// with the default values initialized.
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams() *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithTimeout creates a new PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithContext creates a new PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams{

		Context: ctx,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithHTTPClient creates a new PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParamsWithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams contains all the parameters to send to the API endpoint
for the post repositories workspace repo slug pullrequests pull request ID approve operation typically these are written to a http.Request
*/
type PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams struct {

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

// WithTimeout adds the timeout to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPullRequestID adds the pullRequestID to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithPullRequestID(pullRequestID int64) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetPullRequestID(pullRequestID)
	return o
}

// SetPullRequestID adds the pullRequestId to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetPullRequestID(pullRequestID int64) {
	o.PullRequestID = pullRequestID
}

// WithRepoSlug adds the repoSlug to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithRepoSlug(repoSlug string) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WithWorkspace(workspace string) *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the post repositories workspace repo slug pullrequests pull request ID approve params
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDApproveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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