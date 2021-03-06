// Code generated by go-swagger; DO NOT EDIT.

package issue_tracker

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

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug issues issue ID watch operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams struct {

	/*IssueID
	  The issue id

	*/
	IssueID string
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

// WithTimeout adds the timeout to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssueID adds the issueID to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithIssueID(issueID string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID watch params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param issue_id
	if err := r.SetPathParam("issue_id", o.IssueID); err != nil {
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
