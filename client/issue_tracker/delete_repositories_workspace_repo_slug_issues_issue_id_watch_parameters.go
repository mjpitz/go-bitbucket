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

// NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams creates a new DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized.
func NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams() *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithTimeout creates a new DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithContext creates a new DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{

		Context: ctx,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithHTTPClient creates a new DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParamsWithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams contains all the parameters to send to the API endpoint
for the delete repositories workspace repo slug issues issue ID watch operation typically these are written to a http.Request
*/
type DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams struct {

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

// WithTimeout adds the timeout to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssueID adds the issueID to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithIssueID(issueID string) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WithRepoSlug adds the repoSlug to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithRepoSlug(repoSlug string) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WithWorkspace(workspace string) *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete repositories workspace repo slug issues issue ID watch params
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoriesWorkspaceRepoSlugIssuesIssueIDWatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
