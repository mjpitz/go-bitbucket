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

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug issues issue ID changes operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams struct {

	/*IssueID
	  The issue id

	*/
	IssueID string
	/*Q

	Query string to narrow down the response. See
	[filtering and sorting](../../../meta/filtering) for details.

	*/
	Q *string
	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Sort

	Name of a response property to sort results. See
	[filtering and sorting](../../../meta/filtering#query-sort)
	for details.


	*/
	Sort *string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssueID adds the issueID to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithIssueID(issueID string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WithQ adds the q to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithQ(q *string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetQ(q *string) {
	o.Q = q
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithSort adds the sort to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithSort(sort *string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID changes params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param issue_id
	if err := r.SetPathParam("issue_id", o.IssueID); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

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
