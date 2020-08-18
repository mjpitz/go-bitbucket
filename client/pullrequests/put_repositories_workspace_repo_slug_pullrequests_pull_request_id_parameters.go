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

	"github.com/mjpitz/go-bitbucket/models"
)

// NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams creates a new PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams object
// with the default values initialized.
func NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams() *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithTimeout creates a new PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams{

		timeout: timeout,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithContext creates a new PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams{

		Context: ctx,
	}
}

// NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithHTTPClient creates a new PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParamsWithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	var ()
	return &PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams{
		HTTPClient: client,
	}
}

/*PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams contains all the parameters to send to the API endpoint
for the put repositories workspace repo slug pullrequests pull request ID operation typically these are written to a http.Request
*/
type PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams struct {

	/*Body
	  The pull request that is to be updated.

	*/
	Body *models.Pullrequest
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

// WithTimeout adds the timeout to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithTimeout(timeout time.Duration) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithContext(ctx context.Context) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithHTTPClient(client *http.Client) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithBody(body *models.Pullrequest) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetBody(body *models.Pullrequest) {
	o.Body = body
}

// WithPullRequestID adds the pullRequestID to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithPullRequestID(pullRequestID int64) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetPullRequestID(pullRequestID)
	return o
}

// SetPullRequestID adds the pullRequestId to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetPullRequestID(pullRequestID int64) {
	o.PullRequestID = pullRequestID
}

// WithRepoSlug adds the repoSlug to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithRepoSlug(repoSlug string) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WithWorkspace(workspace string) *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the put repositories workspace repo slug pullrequests pull request ID params
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PutRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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