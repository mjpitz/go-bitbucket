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

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams() *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug issues issue ID attachments path operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams struct {

	/*IssueID*/
	IssueID string
	/*Path*/
	Path string
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

// WithTimeout adds the timeout to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssueID adds the issueID to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithIssueID(issueID string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WithPath adds the path to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithPath(path string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetPath(path string) {
	o.Path = path
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug issues issue ID attachments path params
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugIssuesIssueIDAttachmentsPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param issue_id
	if err := r.SetPathParam("issue_id", o.IssueID); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
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
