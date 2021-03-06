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

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug issues issue ID changes change ID operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams struct {

	/*ChangeID
	  The issue change id

	*/
	ChangeID string
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

// WithTimeout adds the timeout to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeID adds the changeID to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithChangeID(changeID string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetChangeID(changeID)
	return o
}

// SetChangeID adds the changeId to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetChangeID(changeID string) {
	o.ChangeID = changeID
}

// WithIssueID adds the issueID to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithIssueID(issueID string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID changes change ID params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesChangeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param change_id
	if err := r.SetPathParam("change_id", o.ChangeID); err != nil {
		return err
	}

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
