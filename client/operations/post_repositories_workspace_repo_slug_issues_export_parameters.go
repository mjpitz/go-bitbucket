// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/mjpitz/go-bitbucket/models"
)

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportParams creates a new PostRepositoriesWorkspaceRepoSlugIssuesExportParams object
// with the default values initialized.
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportParams() *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithTimeout creates a new PostRepositoriesWorkspaceRepoSlugIssuesExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithContext creates a new PostRepositoriesWorkspaceRepoSlugIssuesExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportParams{

		Context: ctx,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithHTTPClient creates a new PostRepositoriesWorkspaceRepoSlugIssuesExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesWorkspaceRepoSlugIssuesExportParamsWithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugIssuesExportParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesWorkspaceRepoSlugIssuesExportParams contains all the parameters to send to the API endpoint
for the post repositories workspace repo slug issues export operation typically these are written to a http.Request
*/
type PostRepositoriesWorkspaceRepoSlugIssuesExportParams struct {

	/*Body
	  The options to apply to the export. Available options include `project_key` and `project_name` which, if specified, are used as the project key and name in the exported Jira json format. Option `send_email` specifies whether an email should be sent upon export result. Option `include_attachments` specifies whether attachments are included in the export.

	*/
	Body *models.ExportOptions
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

// WithTimeout adds the timeout to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithBody(body *models.ExportOptions) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetBody(body *models.ExportOptions) {
	o.Body = body
}

// WithRepoSlug adds the repoSlug to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithRepoSlug(repoSlug string) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WithWorkspace(workspace string) *PostRepositoriesWorkspaceRepoSlugIssuesExportParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the post repositories workspace repo slug issues export params
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesWorkspaceRepoSlugIssuesExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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